package core_test

import (
	"testing"

	"github.com/cassiofariasmachado/remove-build-assets/core"
	"github.com/stretchr/testify/require"
)

func TestReadConfigValidFile(t *testing.T) {
	expectedConfig := &core.Config{
		Stacks: []core.StackConfig{
			{
				FoldersToRemove:    []string{"node_modules"},
				ProjectFilePattern: "package.json",
			},
			{
				FoldersToRemove:    []string{"bin", "obj"},
				ProjectFilePattern: "*.csproj",
			},
		},
	}

	config, err := core.ReadConfig("../defaultConfig.json")

	require.NoError(t, err)
	require.Equal(t, expectedConfig, config)
}

func TestReadConfigNotExistentFile(t *testing.T) {
	_, err := core.ReadConfig("../invalidConfig.json")

	require.Error(t, err)
}
