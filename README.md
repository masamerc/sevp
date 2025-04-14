# SEVP: **Simple Environment Variable Picker**

![SEVP](./assets/sevp.png)

SEVP is a lightweight CLI/TUI tool designed to simplify switching between environment variable values. 

![SEVP_DEMO](./assets/sevp-demo.gif)

> [!Note]
> SEVP uses a shellhook to set environment variables for the current shell. It currently supports:
> - `zsh`
> - `bash`

- [What is SEVP?](#what-is-sevp)
- [Usage](#usage)
- [Configuration](#configuration)
- [Installation](#installation)
- [Compatibility with `direnv`](#compatibility-with-direnv)
- [Contributing](#contributing)

## What is SEVP?

SEVP helps you quickly switch between predefined values for environment variables with a simple and interactive interface. For example:
- Select an AWS profile by setting the `AWS_PROFILE` variable.
- Switch Docker contexts by setting the `DOCKER_CONTEXT` variable.
- Manage custom environment variables for your projects.

## Usage

> [!Important]
> SEVP will create a default config file at `$HOME/.config/sevp.toml` if it doesn't exist when you run it for the first time.

### Main Command: `sevp`

Running `sevp` without arguments uses the `default` target from your configuration. It launches a TUI where you can select a value for the environment variable.

```bash
$ sevp
```

For example, selecting an AWS profile will set the `AWS_PROFILE` variable.

### Specify a Target

You can specify a target from your configuration by passing its name as an argument.

```bash
$ sevp google_cloud
$ sevp my_custom_var
```

### List Available Targets: `sevp list`

```bash
$ sevp list
```

This command lists all available targets defined in your configuration.

![SEVP_LIST_DEMO](./assets/list-demo.gif)

### View Target Configuration: `sevp view`

```bash
$ sevp view aws
$ sevp view google_cloud
$ sevp view my_custom_var
```

This command displays the configuration for a specific target, including the `target_var` and its `possible_values`.

![SEVP_VIEW_DEMO](./assets/view-demo.gif)

## Configuration

### Custom Configuration: `~/.config/sevp.toml`

SEVP uses a configuration file to define environment variables and their possible values. It looks for the config file in the following locations (in order of precedence):
1. `$HOME/.config/sevp.toml` (default location)
2. `$HOME/sevp.toml`

Here’s an example configuration:

```toml
# Here we specify which target to use when using SEVP without any argument
default = "aws"

# ======================================================================
# External Config Selectors
#
# Currently the following targets support reading external settings:
# - aws: source settings from ~/.aws/config for AWS_PROFILE
# - docker-context: source settings from ~/.docker/contexts/meta dir
# - tfenv: source settings from ~/.tfenv/versions dir
# ======================================================================

[aws]
external_config = false # true -> read profiles from ~/.aws/config
target_var = "AWS_PROFILE"
possible_values = ["prod1", "prod2"]

[docker-context]
external_config = true # true -> read contexts from ~/.docker/contexts/meta
target_var = "DOCKER_CONTEXT"
possible_values = ["default", "sample1", "sample2"]

[tfenv]
external_config = true # true -> read versions from ~/.tfenv/versions
target_var = "TFENV_TERRAFORM_VERSION"
possible_values = ["1.0.0", "0.1.1"]


# ======================================================================
# User-defined Config Selectors
# 
# The following are user-defined targets with manual configuration.
# ======================================================================

[google_cloud]
target_var = "GOOGLE_CLOUD_PROJECT"
possible_values = ["proj1", "proj2", "proj3"]

[some_var]
target_var = "MY_CUSTOM_ENV_VAR"
possible_values = ["val1", "val2"]
```

### External Config Providers

External Config Providers allow SEVP to dynamically fetch values from external configuration files or directories. This is useful for tools like AWS CLI or Docker.

#### Supported Providers:
- **AWS**  
   - Reads profiles from `~/.aws/config`.
   - Automatically sets the `AWS_PROFILE` environment variable.
   - Enable by setting `external_config = true` in the `[aws]` section.

- **Docker Context**  
   - Reads contexts from `~/.docker/contexts/meta`.
   - Automatically sets the `DOCKER_CONTEXT` environment variable.
   - Enable by setting `external_config = true` in the `[docker-context]` section.

- **tfenv**
   - Support for https://github.com/tfutils/tfenv. 
   - Reads versions from `~/.tfenv/versions`.
   - Automatically sets the `TFENV_TERRAFORM_VERSION` environment variable.
   - Enable by setting `external_config = true` in the `[tfenv]` section.

#### How It Works:
- When `external_config = true`, SEVP ignores the `possible_values` field and dynamically fetches values from the external configuration.
- This ensures SEVP stays in sync with changes made outside the tool.

#### Example Configuration:
```toml
[aws]
external_config = true
target_var = "AWS_PROFILE"

[docker-context]
external_config = true
target_var = "DOCKER_CONTEXT"

[tfenv]
external_config = true
target_var = "TFENV_TERRAFORM_VERSION"
```

## Installation

### Homebrew (macOS & Linux)

1. Tap the repository:
   ```bash
   $ brew tap masamerc/sevp https://github.com/masamerc/sevp.git
   ```

2. Install:
   ```bash
   $ brew install sevp
   ```

3. Add the shellhook to your shell configuration:
   ```bash
   eval "$(sevp init <shell>)"
   ```

### Shell One-liner (requires `bash`)

1. Run the one-liner:
   ```bash
   $ curl -sSL https://raw.githubusercontent.com/masamerc/sevp/pre-release/scripts/install.sh | bash
   ```

2. Add the shellhook to your shell configuration:
   ```bash
   eval "$(sevp init <shell>)"
   ```

### Go Install

1. Install using `go`:
   ```bash
   $ go install github.com/masamerc/sevp@latest
   ```

2. Add the shellhook to your shell configuration:
   ```bash
   eval "$(sevp init <shell>)"
   ```

### Build from Source

1. Clone the repository:
   ```bash
   $ git clone https://github.com/masamerc/sevp.git
   $ cd sevp
   ```

2. Build and install (requires [Taskfile](https://taskfile.dev/)):
   ```bash
   $ task install
   ```

3. Add the shellhook to your shell configuration:
   ```bash
   eval "$(sevp init <shell>)"
   ```

## Compatibility with `direnv`

SEVP may conflict with tools like [`direnv`](https://direnv.net/) since both rely on shell hooks. The order of evaluation determines which tool takes precedence.

- **SEVP takes precedence**  
  Add SEVP after `direnv` in your shell configuration:
  ```sh
  eval "$(direnv hook zsh)"
  eval "$(sevp init zsh)"
  ```

- **`direnv` takes precedence**  
  Add `direnv` after SEVP:
  ```sh
  eval "$(sevp init zsh)"
  eval "$(direnv hook zsh)"
  ```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING](CONTRIBUTING.md) for guidelines.
