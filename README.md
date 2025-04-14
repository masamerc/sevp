# SEVP: **Simple Environment Variable Picker**  

![SEVP](./assets/sevp.png)

A CLI/TUI for seamlessly switching between values for environment variables.

![SEVP_DEMO](./assets/sevp-demo.gif)

> [!Note]
> This program uses a shellhook to set the environment variable for the current shell and currently supports
> - `zsh`
> - `bash`

- [Usage](#usage)
- [Configuration](#configuration)
- [Installation](#installation)

## Usage

> [!Important]
> `sevp` will create a default config file in `$HOME/.config/sevp.toml` if it doesn't exist when you run it for the first time.


### Main command: `sevp`

Running `sevp` without any arguments will read your config and use the `default` target.
- `sevp` command will bring up a TUI for you to select a value. 
- for `aws` target as an example, you will be choosing a profile by setting the `AWS_PROFILE` environment variable.


```bash
$ sevp
```

**Select target**

You can select a target from your config ([configuration](#configuration)) by passing its name as an argument.

```bash
$ sevp google_cloud
$ sevp my_custom_var
```

### List available targets: `sevp list` (config required)

```bash
$ sevp list
```

`list` will list all available targets defined in your config. ([configuration](#configuration))

![SEVP_LIST_DEMO](./assets/list-demo.gif)


### View target configuration: `sevp view` (config required)

```bash
$ sevp view aws
$ sevp view google_cloud
$ sevp view my_custom_var
```

Displays the configuration for a specific target—shows the `target_var` and its `possible_values`.

![SEVP_VIEW_DEMO](./assets/view-demo.gif)

## Configuration

### Custom configuration
You can use a configuration file to define your own environment variables and the values you want to switch between.

Here’s a sample config: 

```toml
# Here we specify which target to use when using SEVP without any argument
default = "aws"

# Currently the following targets (called external config providers) support reading external settings:
# - aws: source settings from ~/.aws/config for AWS_PROFILE
# - docker-context: source settings from ~/.docker/contexts/meta dir

[aws]
# external config provider has a special option called `read_config` 
# if true, it will read profiles from ~/.aws/config
read_config = false 
# if read_config = true, the following options are ignored
target_var = "AWS_PROFILE"
possible_values = ["prod1", "prod2"]

[docker-context]
read_config = true
target_var = "DOCKER_CONTEXT"
possible_values = ["default", "sample1", "sample2"]

# The following are user-defined targets with manual configuration
[google_cloud]
target_var = "GOOGLE_CLOUD_PROJECT"
possible_values = ["proj1", "proj2", "proj3"]

[some_var]
target_var = "MY_CUSTOM_ENV_VAR"
possible_values = ["val1", "val2"]

```

 `sevp` will look for the config file in the following locations (in order of precedence):
- `$HOME/.config/sevp.toml`
- `$HOME/sevp.toml`

Anatomy of the config file:
- `default`: Specifies the default target to use when no argument is provided to `sevp`.
- Each section (e.g., `[aws]`, `[google_cloud]`) defines a **target**.
  - The `[aws]` target supports `read_config = true` to auto-load profiles from `~/.aws/config`.
  - You can also provide custom values for `[aws]` by setting `read_config = false` and defining `target_var` and `possible_values`.
- `target_var`: The name of the environment variable you want to manage.
- `possible_values`: A list of values that can be selected for that variable.

### External Config Provider

External Config Providers allow `sevp` to dynamically read and manage environment variable values from external configuration files or directories. This feature is particularly useful for tools like AWS CLI or Docker, where profiles or contexts are defined externally.

#### Supported External Config Providers:
1. **AWS**  
   - Reads profiles from `~/.aws/config`.
   - Automatically sets the `AWS_PROFILE` environment variable based on the selected profile.
   - To enable this, set `read_config = true` in the `[aws]` section of your configuration.

2. **Docker Context**  
   - Reads Docker contexts from `~/.docker/contexts/meta`.
   - Automatically sets the `DOCKER_CONTEXT` environment variable based on the selected context.
   - To enable this, set `read_config = true` in the `[docker-context]` section of your configuration.

#### How It Works:
- When `read_config = true` is set for a target, `sevp` will ignore the `possible_values` field and instead fetch the values dynamically from the respective external configuration.
- This allows `sevp` to stay in sync with changes made outside of the tool, such as adding or removing AWS profiles or Docker contexts.

#### Example Configuration:
```toml
[aws]
read_config = true
target_var = "AWS_PROFILE"

[docker-context]
read_config = true
target_var = "DOCKER_CONTEXT"
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

3. Add the shellhook to your shell config file:
```bash
eval "$(sevp init <shell>)"
```

### Shell One-liner (requires `bash`)
1. Run the one-liner:
```bash
$ curl -sSL https://raw.githubusercontent.com/masamerc/sevp/pre-release/scripts/install.sh | bash
```

2. Add the shellhook to your shell config file:
```bash
eval "$(sevp init <shell>)"
```

### Go Install
1. Run `go install`:
```
$ go install github.com/masamerc/sevp@latest
```

2. Add the shellhook to your shell config file:
```bash
eval "$(sevp init <shell>)"
```

### Build from Source
1. Clone this repository:
```bash
$ git clone https://github.com/masamerc/sevp.git
$ cd sevp
```

2. Run thne install command (requires Taskfile https://taskfile.dev/)
```bash
$ task install
```

3. Add the shellhook to your shell config file:
```bash
eval "$(sevp init <shell>)"
```

## Notes on `direnv` Compatibility

Environment variables set by `sevp` may conflict with tools like [`direnv`](https://direnv.net/) since both rely on shell hooks (e.g., in `.zshrc`). The order in which these hooks are evaluated determines which tool takes precedence.

- **`sevp` takes precedence**  
  `sevp` is evaluated after `direnv`:
  ```sh
  eval "$(direnv hook zsh)"
  eval "$(sevp init zsh)"
  ```

- **`direnv` takes precedence**  
`direnv` is evaluated after `sevp`
  ```sh
  eval "$(sevp init zsh)"
  eval "$(direnv hook zsh)"
  ```

## Contribution
Any contribution is welcome and appreciated!
Please see [CONTRIBUTING](CONTRIBUTING.md).
