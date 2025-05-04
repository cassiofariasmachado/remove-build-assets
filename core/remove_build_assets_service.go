package core

import (
	"log"
	"os"
	"path/filepath"
)

type RemoveBuildAssetsService struct {
	FoldersToRemove []string
	Config          Config
}

func NewRemoveBuildAssetsService(config Config) *RemoveBuildAssetsService {
	return &RemoveBuildAssetsService{
		FoldersToRemove: []string{},
		Config:          config,
	}
}

func (r *RemoveBuildAssetsService) ListFolders(rootPath string) {
	pathsToRemove := make(map[string]*StackConfig)

	err := filepath.Walk(rootPath, r.createWalkFunc(pathsToRemove))
	if err != nil {
		log.Printf("Error walking the file tree: %v", err)
	}

	log.Printf("Folders to remove: %v", r.FoldersToRemove)
}

func (r *RemoveBuildAssetsService) createWalkFunc(pathsToRemove map[string]*StackConfig) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v", path, err)
			return err
		}

		if !info.IsDir() {
			return nil
		}

		r.processDirectory(path, pathsToRemove)
		return nil
	}
}

func (r *RemoveBuildAssetsService) processDirectory(path string, pathsToRemove map[string]*StackConfig) {
	parentPath, file := filepath.Split(path)

	for _, config := range r.Config.Stacks {
		if r.matchesProjectFilePattern(parentPath, config) {
			for _, folder := range config.FoldersToRemove {
				pathsToRemove[folder] = &config
			}
		}
	}

	if _, exists := pathsToRemove[file]; exists {
		r.FoldersToRemove = append(r.FoldersToRemove, path)
	}
}

func (r *RemoveBuildAssetsService) matchesProjectFilePattern(parentPath string, config StackConfig) bool {
	matches, err := filepath.Glob(filepath.Join(parentPath, config.ProjectFilePattern))
	if err != nil {
		log.Printf("Error verifying project file pattern (%v): %v", config.ProjectFilePattern, err)
		return false
	}
	return len(matches) > 0
}

func (r *RemoveBuildAssetsService) RemoveFolders() {
	for _, folder := range r.FoldersToRemove {
		os.RemoveAll(folder)
	}
}
