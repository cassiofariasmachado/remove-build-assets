package core

import (
	"log"
	"os"
	"path/filepath"

	"github.com/cassiofariasmachado/rm-build-assets/utils"
)

type RemoveBuildAssetsService struct {
	FoldersToRemove []string
	TotalSize       int64
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

	err := filepath.Walk(rootPath, r.listWalkFunc(pathsToRemove))

	if err != nil {
		log.Printf("Error walking the file tree: %v", err)
	}
}

func (r *RemoveBuildAssetsService) Summary() {
	r.TotalSize = 0

	for _, folder := range r.FoldersToRemove {
		size, err := utils.DirSize(folder)

		if err != nil {
			log.Printf("Error getting size of the folder: %v, Error: %v", folder, err)
			continue
		}

		log.Printf("Folder: \"%v\", Size: %.2f MB", folder, utils.ToMB(size))
		r.TotalSize += size
	}

	log.Printf("Total size: %.2f MB", utils.ToMB(r.TotalSize))
}

func (r *RemoveBuildAssetsService) listWalkFunc(pathsToRemove map[string]*StackConfig) filepath.WalkFunc {
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
		if !r.matchesProjectFilePattern(parentPath, config.ProjectFilePattern) {
			continue
		}

		for _, folder := range config.FoldersToRemove {
			pathsToRemove[folder] = &config
		}
	}

	if _, exists := pathsToRemove[file]; exists {
		r.FoldersToRemove = append(r.FoldersToRemove, path)
	}
}

func (r *RemoveBuildAssetsService) matchesProjectFilePattern(parentPath string, projectFilePattern string) bool {
	matches, err := filepath.Glob(filepath.Join(parentPath, projectFilePattern))

	if err != nil {
		log.Printf("Error verifying project file pattern (%v): %v", projectFilePattern, err)
		return false
	}

	return len(matches) > 0
}

func (r *RemoveBuildAssetsService) RemoveFolders() {
	for _, folder := range r.FoldersToRemove {
		os.RemoveAll(folder)
	}
}
