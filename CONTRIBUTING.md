# Contributing to SEVP

Thank you for considering contributing to SEVP! This project follows a standard workflow using feature branches and pull requests.

## Prerequisites
- Taskfile (https://taskfile.dev/)
- golanci-lint (https://github.com/golangci/golangci-lint)
- gosec (https://github.com/securego/gosec)
- gitleaks (https://github.com/gitleaks/gitleaks)
- go 1.22+

## Contribution Guide

1. **Fork the repository**  
Begin by forking this repository and creating a new branch for your changes.

2. **Test, lint and scan**  
 Before pushing your branch, please run the following tasks to ensure everything is in order:
   ```bash
   $ task test
   $ task scan
   $ task lint
   ```
3. **Open a Pull Request**  
   Open a pull request to the main repository and provide a description of your changes.
