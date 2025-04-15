# Contributing to SEVP

Thank you for considering contributing to SEVP! This project follows a standard workflow using feature branches and pull requests.

## Prerequisites
- Taskfile (https://taskfile.dev/)
- golanci-lint (https://github.com/golangci/golangci-lint)
- gosec (https://github.com/securego/gosec)
- gitleaks (https://github.com/gitleaks/gitleaks)
- go 1.22+

## Contribution Guide

### General

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

### Add Custom `extconfig` Support

To add support for a new external configuration provider:
- Create a new file in the `internal/extconfig` directory, e.g., `myprovider.go`.
- Implement the `Selector` interface by defining the `Read()` method to fetch the target variable and possible values.
- Write unit tests for your implementation in a corresponding `_test.go` file.
- Update `GetExternalConfigSelector` in `internal/extconfig.go` to include your new provider.

Example:
```go
// filepath: /Users/masafukui/personal/sevp/internal/extconfig/myprovider.go
package extconfig

type MyProviderSelector struct{}

func (s *MyProviderSelector) Read() (string, []string, error) {
   // Implement logic to fetch target variable and possible values
   return "MY_PROVIDER_VAR", []string{"value1", "value2"}, nil
}

func NewMyProviderSelector() *MyProviderSelector {
   return &MyProviderSelector{}
}
```

```go
// filepath: /Users/masafukui/personal/sevp/internal/extconfig.go
func GetExternalConfigSelector(selectorName string) (Selector, error) {
   switch selectorName {
   case "myprovider":
       return NewMyProviderSelector(), nil
   // ...existing code...
   default:
       return nil, fmt.Errorf("the external config provider is not supported for selector %s", selectorName)
   }
}
```
