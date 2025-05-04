package main

import (
	"log"

	"github.com/cassiofariasmachado/remove-build-assets/core"
)

func main() {
	config, err := core.ReadConfig("defaultConfig.json")

	if err != nil {
		log.Fatalf("Error reading default config: %v", err)
	}

	service := core.NewRemoveBuildAssetsService(*config)

	service.ListFolders("path/to")

	service.RemoveFolders()
}
