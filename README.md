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

## Requirements
- [Task](https://taskfile.dev/) runner (for building from source)
- Go 1.22+

---

## Installation

### One-liner
1. Run the one-liner:
```bash
$ curl -sSL https://raw.githubusercontent.com/masamerc/sevp/pre-release/scripts/install.sh | sh
```

### Build From Source
1. Clone this repository:
```bash
$ git clone https://github.com/masamerc/sevp.git
$ cd sevp
```

2. Run thne install command:
```bash
$ task install
```

### Go Install
1. Run `go install`:
```
$ go install github.com/masamerc/sevp@latest
```

2. Install the shellhook for your shell:
```bash
$ curl -sSL https://raw.githubusercontent.com/masamerc/sevp/pre-release/scripts/install_shellhook.sh | sh
```

## Usage
Simply run sevp which will bring up a TUI for selecting a value:
```bash
$ sevp
```

## Todo
- [x] CI
- [x] automatic tagging & releasing
- [x] one-liner installation 
- [x] support all commonly used shells
- [ ] support for other types of environment variables (currently only supports AWS_PROFILE)
- [ ] web docs
