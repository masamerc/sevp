# SEVP: **Simple Environment Variable Picker**  

![SEVP](./assets/sevp.png)

A lightweight TUI for seamlessly switching environment variable values.

> [!Important]
> SEVP is a work in progress and **currently supports only `AWS_PROFILE`**.  
> This program uses a shellhook to set the environment variable for the current shell and currently supports
> - `zsh` (tested)
> - `bash` (tested)
> - `fish` (not tested)
> - `nu` (not tested)

---

## Features

![SEVP_DEMO](./assets/sevp-demo.gif)

- A TUI for quickly switching environment variable values for `AWS_PROFILE` which persists in the current shell.
- Search through a list of AWS profiles declared in `~/.aws/config`.

---

## Requirements (for building from source)
- [Task](https://taskfile.dev/) runner (for building from source)
- Go 1.22+

---

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

2. Run thne install command:
```bash
$ task install
```

3. Add the shellhook to your shell config file:
```bash
eval "$(sevp init <shell>)"
```

## Usage
Simply run sevp which will bring up a TUI for selecting a value:
```bash
$ sevp
```

## Notes on `direnv` compatibility 

## Todo
- [x] CI
- [x] automatic tagging & releasing
- [x] one-liner installation 
- [x] support all commonly used shells
- [ ] support for other types of environment variables (currently only supports AWS_PROFILE)
- [ ] wiki docs

