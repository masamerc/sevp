package extconfig

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDockerContextsFromPath(t *testing.T) {
	tmp := t.TempDir()

	// Simulate meta dir context structure
	ctx1 := filepath.Join(tmp, "ctx-1")
	ctx2 := filepath.Join(tmp, "ctx-2")
	_ = os.MkdirAll(ctx1, 0755)
	_ = os.MkdirAll(ctx2, 0755)

	// Create meta.json files in the meta dir
	writeMeta := func(dir, name string) {
		meta := map[string]string{"Name": name}
		b, _ := json.Marshal(meta)
		_ = os.WriteFile(filepath.Join(dir, "meta.json"), b, 0644)
	}
	writeMeta(ctx1, "default")
	writeMeta(ctx2, "custom")

	contexts, err := parseDockerContexts(tmp)
	require.NoError(t, err)
	require.ElementsMatch(t, []string{"default", "custom"}, contexts)
}

func TestGetDockerContextsFromPath_Empty(t *testing.T) {
	tmp := t.TempDir()
	_, err := parseDockerContexts(tmp)
	require.Error(t, err)
	require.Contains(t, err.Error(), "no docker contexts")
}
