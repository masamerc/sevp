package extconfig

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type DockerContextSelector struct{}

func (s *DockerContextSelector) Read() (string, []string, error) {
	targetVar := "DOCKER_CONTEXT"
	contexts, err := getDockerContextsFromMeta()
	return targetVar, contexts, err
}

func NewDockerContextSelector() *DockerContextSelector {
	return &DockerContextSelector{}
}

type dockerContextMeta struct {
	Name string `json:"Name"`
}

// getDockerContextsFromMeta returns the names of all docker contexts in the meta dir
func getDockerContextsFromMeta() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	return parseDockerContexts(filepath.Join(home, ".docker", "contexts", "meta"))
}

// parseDockerContexts returns the names of all docker contexts in the meta dir
func parseDockerContexts(metaDir string) ([]string, error) {
	entries, err := os.ReadDir(metaDir)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		metaPath := filepath.Join(metaDir, entry.Name(), "meta.json")
		data, err := os.ReadFile(filepath.Clean(metaPath))
		if err != nil {
			continue
		}

		var meta dockerContextMeta
		if err := json.Unmarshal(data, &meta); err != nil {
			continue
		}

		if meta.Name != "" {
			names = append(names, meta.Name)
		}
	}

	if len(names) == 0 {
		return nil, errors.New("no docker contexts found")
	}

	return names, nil
}
