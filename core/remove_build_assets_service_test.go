package core_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/cassiofariasmachado/remove-build-assets/core"
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

	createFiles()
}

func createFiles() {
	os.MkdirAll("tmp/path/to/cs_project/bin", os.ModeDir)
	os.MkdirAll("tmp/path/to/cs_project/obj", os.ModeDir)
	os.Create("tmp/path/to/cs_project/project.csproj")

	os.MkdirAll("tmp/path/to/js_project/node_modules", os.ModeDir)
	os.Create("tmp/path/to/js_project/package.json")
}

func removeFiles() {
	os.RemoveAll("tmp")
}

func TestListFoldersShouldListFoldersToRemoveCorrectly(t *testing.T) {
	service := core.NewRemoveBuildAssetsService(*config)

	service.ListFolders("tmp/path/to")

	expecteFolderToRemove := []string{
		filepath.Clean("tmp/path/to/cs_project/bin"),
		filepath.Clean("tmp/path/to/cs_project/obj"),
		filepath.Clean("tmp/path/to/js_project/node_modules"),
	}

	require.Equal(t, expecteFolderToRemove, service.FoldersToRemove)

	t.Cleanup(func() {
		removeFiles()
	})
}

func TestListFoldersShouldListFoldersToRemoveCorrectlyWhenEmpty(t *testing.T) {
	service := core.NewRemoveBuildAssetsService(*config)

	service.ListFolders("non/existent/path/to")

	expecteFolderToRemove := []string{}

	require.Equal(t, expecteFolderToRemove, service.FoldersToRemove)

	t.Cleanup(func() {
		removeFiles()
	})
}
