package internal

import (
	"fmt"

	"github.com/masamerc/sevp/internal/extconfig"
)

// GetExternalConfigSelector returns the appropriate Selector implementation based on the selectorName.
func GetExternalConfigSelector(selectorName string) (Selector, error) {
	switch selectorName {
	case "aws":
		return extconfig.NewAWSProfileSelector(), nil
	case "docker-context":
		return extconfig.NewDockerContextSelector(), nil
	case "tfenv":
		return extconfig.NewTfEnvSelector(), nil
	case "goenv":
		return extconfig.NewGoEnvSelector(), nil
	default:
		return nil, fmt.Errorf("the external config provider is not supported for selector %s", selectorName)
	}
}
