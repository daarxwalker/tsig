package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func generate(cfg *config) ([]string, error) {
	result := make([]string, 0)
	for _, opt := range cfg.Options {
		dir := filepath.Join(cfg.Root, opt.Dir)
		if walkErr := filepath.WalkDir(
			dir,
			func(path string, entry os.DirEntry, err error) error {
				if !entry.IsDir() || path != dir {
					return nil
				}
				subresult, createIndexErr := writeIndex(path, opt.Export, opt.Recursive)
				if createIndexErr != nil {
					return fmt.Errorf("create index: %w", createIndexErr)
				}
				result = append(result, subresult...)
				return nil
			},
		); walkErr != nil {
			return nil, fmt.Errorf("walk dir %s: %w", dir, walkErr)
		}
		result = append(result, filepath.Join(dir, indexFilename))
	}
	return result, nil
}
