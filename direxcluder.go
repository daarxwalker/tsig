package main

import (
	"path/filepath"
	"strings"
)

func createExcludedDirs(cfg *config) []string {
	optionDirs := make([]string, len(cfg.Options))
	excludedDirs := make([]string, 0, len(cfg.Options))
	for i, opt := range cfg.Options {
		optionDirs[i] = filepath.Join(cfg.Root, opt.Dir)
	}
	for _, dir1 := range optionDirs {
		var excludedDir string
		for _, dir2 := range optionDirs {
			if dir1 == dir2 || !strings.HasPrefix(dir1, dir2) {
				continue
			}
			excludedDir = dir1
			break
		}
		if len(excludedDir) == 0 {
			continue
		}
		excludedDirs = append(excludedDirs, excludedDir)
	}
	return excludedDirs
}
