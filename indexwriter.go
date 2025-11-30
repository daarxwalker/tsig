package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func writeIndex(dir string, export export, recursive bool, excludedDirs []string) ([]string, error) {
	result := make([]string, 0)
	b := new(strings.Builder)
	if walkErr := filepath.WalkDir(
		dir,
		func(path string, entry fs.DirEntry, err error) error {
			if entry.Name() == indexFilename || path == dir {
				return nil
			}
			if entry.IsDir() && recursive {
				if len(excludedDirs) > 0 && slices.Contains(excludedDirs, path) {
					return nil
				}
				subresult, writeErr := writeIndex(path, export, recursive, excludedDirs)
				if writeErr != nil {
					return writeErr
				}
				result = append(result, subresult...)
				return nil
			}
			isFileInSubdir := strings.Index(strings.TrimPrefix(path, dir+"/"), "/") > -1
			if isFileInSubdir {
				return nil
			}
			nameWithSuffix := entry.Name()
			name := nameWithSuffix[:strings.IndexByte(nameWithSuffix, '.')]
			b.WriteString("export ")
			switch export {
			case exportSingle:
				b.WriteString("{ ")
				b.WriteString(name)
				b.WriteString(" }")
			case exportDefault:
				b.WriteString("{ ")
				b.WriteString("default as ")
				b.WriteString(name)
				b.WriteString(" }")
			case exportAll:
				b.WriteString("*")
			}
			b.WriteString(" from './")
			b.WriteString(nameWithSuffix)
			b.WriteString("'")
			b.WriteString("\n")
			return nil
		},
	); walkErr != nil {
		return nil, fmt.Errorf("walk dir: %w", walkErr)
	}
	indexPath := filepath.Join(dir, indexFilename)
	if writeErr := os.WriteFile(indexPath, []byte(b.String()), 0777); writeErr != nil {
		return nil, fmt.Errorf("write index file %s: %w", indexPath, writeErr)
	}
	result = append(result, indexPath)
	return result, nil
}
