# SEVP: **Simple Environment Variable Picker**  

![SEVP](./assets/sevp.png)

A CLI/TUI for seamlessly switching environment variable values. 

**Started as a AWS-profile switcher**

This project began as a simple terminal UI to quickly switch between AWS profiles, with built-in support for reading profiles from `~/.aws/config`.

**Now a flexible environment variable switcher**

Over time, it evolved to support more flexible use cases—now you can define any environment variable as the target and provide your own list of values to toggle between via a custom config file!


![SEVP_DEMO](./assets/sevp-demo.gif)

> [!Important]
> This program uses a shellhook to set the environment variable for the current shell and currently supports
> - `zsh`
> - `bash`

- [Usage](#usage)
- [Configuration](#configuration)
- [Installation](#installation)


## Usage

### Main command: `sevp`

Running `sevp` without any arguments will read your config and use the `default` target.

```bash
$ sevp
```

**Select target**

With a custom config, you can select a custom target ([configuration](#configuration)) by passing its name as an argument.

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

> [!Note]
> SEVP will create a default config file in `$HOME/.config/sevp.toml` if it doesn't exist when you run it for the first time.


Here’s a sample config: 

```toml
# which target to use when using SEVP without any argument
default = "aws"

# currently only aws supports read_config (~/.aws/config)
[aws]
read_config = true

# user-defined sets of target variable and list of values
[google_cloud]
target_var = "GOOGLE_CLOUD_PROJECT"
possible_values = ["proj1", "proj2", "proj3"]

[some_var]
target_var = "MY_CUSTOM_ENV_VAR"
possible_values = ["val1", "val2"]
```

 SEVP will look for the config file in the following locations (in order of precedence):
- `$HOME/.config/sevp.toml`
- `$HOME/sevp.toml`

Anatomy of the config file:
- `default`: Specifies the default target to use when no argument is provided to `sevp`.
- Each section (e.g., `[aws]`, `[google_cloud]`) defines a **target**.
  - The `[aws]` target supports `read_config = true` to auto-load profiles from `~/.aws/config`.
  - You can also provide custom values for `[aws]` by setting `read_config = false` and defining `target_var` and `possible_values`.
- `target_var`: The name of the environment variable you want to manage.
- `possible_values`: A list of values that can be selected for that variable.


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

Environment variables set by SEVP may conflict with tools like [`direnv`](https://direnv.net/) since both rely on shell hooks (e.g., in `.zshrc`). The order in which these hooks are evaluated determines which tool takes precedence.

- **SEVP takes precedence**  
  SEVP is evaluated after `direnv`:
  ```sh
  eval "$(direnv hook zsh)"
  eval "$(sevp init zsh)"
  ```

- **`direnv` takes precedence**  
`direnv` is evaluated after SEVP
  ```sh
  eval "$(sevp init zsh)"
  eval "$(direnv hook zsh)"
  ```

## Contribution
Any contribution is welcome and appreciated!
Please see [CONTRIBUTING](CONTRIBUTING.md).
