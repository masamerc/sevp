# SEVP: **Simple Environment Variable Picker**  

![SEVP](https://github.com/user-attachments/assets/00402040-99ff-4bdd-81ce-757b09fc62cc)

A lightweight TUI for seamlessly switching environment variable values.

> [!Important]
> SEVP is a work in progress and **currently supports only `AWS_PROFILE`**.  
> This program uses a shellhook to set the environment variable for the current shell and currently supports
> - `zsh`
> - `bash`
> - `fish`
> - `nu`

---

## Features
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


3. Install the shellhook for your shell:
```bash
$ ./scripts/install_shellhook.sh
```

### Go Install
1. Run `go install`:
```
$ go install github.com/masamerc/sevp@latest
```

2. Install the shellhook for shell:
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
- [ ] support installation via homebrew
- [ ] support for other types of environment variables (currently only supports AWS_PROFILE)
