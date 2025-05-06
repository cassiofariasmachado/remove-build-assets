package core_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/cassiofariasmachado/rm-build-assets/core"
	"github.com/stretchr/testify/require"
)

var (
	config *core.Config
)

func init() {
	var err error

	config, err = core.ReadConfig("../defaultConfig.json")

	if err != nil {
		log.Fatal("Error reading default config.")
	}
}

func createMockFiles() {
	os.MkdirAll("tmp/path/to/cs_project/bin", os.ModeDir)
	os.MkdirAll("tmp/path/to/cs_project/obj", os.ModeDir)
	file, _ := os.Create("tmp/path/to/cs_project/project.csproj")

	if file != nil {
		defer file.Close()
	}

	os.MkdirAll("tmp/path/to/js_project/node_modules", os.ModeDir)
	file, _ = os.Create("tmp/path/to/js_project/package.json")

	if file != nil {
		defer file.Close()
	}
}

func removeMockFiles() {
	os.RemoveAll("tmp")
}

func TestListFoldersShouldListFoldersToRemoveCorrectly(t *testing.T) {
	createMockFiles()

	service := core.NewRemoveBuildAssetsService(*config)

	service.ListFolders("tmp/path/to")

	expecteFolderToRemove := []string{
		filepath.Clean("tmp/path/to/cs_project/bin"),
		filepath.Clean("tmp/path/to/cs_project/obj"),
		filepath.Clean("tmp/path/to/js_project/node_modules"),
	}

	require.Equal(t, expecteFolderToRemove, service.FoldersToRemove)

	t.Cleanup(func() {
		removeMockFiles()
	})
}

func TestListFoldersShouldListFoldersToRemoveCorrectlyWhenEmpty(t *testing.T) {
	createMockFiles()

	service := core.NewRemoveBuildAssetsService(*config)

	service.ListFolders("non/existent/path/to")

	expecteFolderToRemove := []string{}

	require.Equal(t, expecteFolderToRemove, service.FoldersToRemove)

	t.Cleanup(func() {
		removeMockFiles()
	})
}

func TestRemoveShouldRemoveFilesCorrectly(t *testing.T) {
	createMockFiles()

	service := core.NewRemoveBuildAssetsService(*config)

	service.ListFolders("tmp/path/to")
	service.RemoveFolders()

	require.NoDirExists(t, "tmp/path/to/cs_project/bin")
	require.NoDirExists(t, "tmp/path/to/cs_project/obj")
	require.NoDirExists(t, "tmp/path/to/js_project/node_modules")

	t.Cleanup(func() {
		removeMockFiles()
	})
}
