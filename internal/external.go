package internal

import (
	"fmt"

	"github.com/masamerc/sevp/internal/external"
)

// GetExternalProviderSelector returns the appropriate Selector implementation based on the selectorName.
func GetExternalProviderSelector(selectorName string) (Selector, error) {
	switch selectorName {
	case "aws":
		return external.NewAWSProfileSelector(), nil
	default:
		return nil, fmt.Errorf("the external config provider is not supported for selector %s", selectorName)
	}
}
