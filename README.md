# SEVP: **Simple Environment Variable Picker**  
A lightweight TUI for seamlessly switching environment variable values.

> [!Important]
> SEVP is a work in progress and **currently supports only `AWS_PROFILE`**.  
> Compatibility is limited to `zsh` at this stage.

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


3. Install the shellhook for `zsh`:
```bash
$ ./scripts/install.sh
```

### Go Install
```
$ go install github.com/masamerc/sevp@latest
```

## Usage
Simply run sevp which will bring up a TUI for selecting a value:
```bash
$ sevp
```

## Todo
- [ ] CI
- [ ] automatic tagging & releasing
- [ ] web installation script
- [ ] support installation via homebrew
- [ ] support other shells than zsh
- [ ] support for other types of environment variables (currently only supports AWS_PROFILE)
