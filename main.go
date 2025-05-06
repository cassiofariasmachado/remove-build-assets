package main

import (
	"flag"
	"log"

	"github.com/cassiofariasmachado/rm-build-assets/core"
	"github.com/cassiofariasmachado/rm-build-assets/utils"
)

var (
	configPath *string
	path       *string
)

func init() {
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.CommandLine.Init("rm-build-assets", flag.ContinueOnError)

	configPath = flag.String("config", "defaultConfig.json", "Name of the gitignore file to download")
	path = flag.String("path", "", "Name of the gitignore file to download")

	flag.Parse()

	if *path == "" {
		log.Fatal("Argument path must be informed")
	}
}

func main() {
	config, err := core.ReadConfig(*configPath)

	if err != nil {
		log.Fatalf("Error reading default config: %v", err)
	}

	service := core.NewRemoveBuildAssetsService(*config)

	service.ListFolders(*path)
	service.Summary()

	confirm := utils.Confirm("Confirm remove operation?", 3)
	if !confirm {
		log.Fatal("Operation cancelled")
	}

	service.RemoveFolders()
}
