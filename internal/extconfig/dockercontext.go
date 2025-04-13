package extconfig

import (
	"os/exec"
	"strings"
)

// DockerContextSelector is a struct that implements the Selector interface for selecting Docker contexts.
type DockerContextSelector struct{}

// Read is the main function of the DockerContextSelector struct, which reads the Docker context names.
func (s *DockerContextSelector) Read() (string, []string, error) {
	targetVar := "DOCKER_CONTEXT"
	contexts, err := getDockerContexts()
	return targetVar, contexts, err
}

// NewDockerContextSelector creates a new empty instance of DockerContextSelector.
func NewDockerContextSelector() *DockerContextSelector {
	return &DockerContextSelector{}
}

// getDockerContexts retrieves a list of Docker context names
func getDockerContexts() ([]string, error) {
	cmd := exec.Command("docker", "context", "ls", "--format", "{{.Name}}")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	return lines, nil
}

