package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	const testFileContent = "ok"
	root := t.TempDir()
	fakeFiles := []string{
		"src/components/Button.tsx",
		"src/components/form/Field.tsx",
		"src/components/style/color.ts",
		"src/components/style/variant.ts",
		"src/pages/Home.tsx",
	}
	expectedIndexes := []string{
		"src/components/index.ts",
		"src/components/form/index.ts",
		"src/components/style/index.ts",
		"src/pages/index.ts",
	}
	expectedIndexContent := map[string]string{
		"src/components/index.ts":       "export { Button } from './Button.tsx'\n",
		"src/components/form/index.ts":  "export { Field } from './Field.tsx'\n",
		"src/components/style/index.ts": "export * from './color.ts'\nexport * from './variant.ts'\n",
		"src/pages/index.ts":            "export { Home } from './Home.tsx'\n",
	}
	// Create the config file
	{
		path := filepath.Join(root, defaultConfigFilename)
		cfg := config{
			Root: root,
			Options: []option{
				{
					Dir:       "src/components",
					Export:    exportSingle,
					Recursive: true,
				},
				{
					Dir:       "src/components/style",
					Export:    exportAll,
					Recursive: false,
				},
				{
					Dir:       "src/pages",
					Export:    exportSingle,
					Recursive: true,
				},
			},
		}
		cfgBts, marshalErr := json.Marshal(cfg)
		require.NoError(t, marshalErr)
		require.NoError(t, os.WriteFile(path, cfgBts, 0777))
		for _, opt := range cfg.Options {
			require.NoError(t, os.MkdirAll(filepath.Join(root, opt.Dir), 0777))
		}
	}
	// Create fake components
	{
		for _, fakeFilePath := range fakeFiles {
			path := filepath.Join(root, fakeFilePath)
			dir := filepath.Join(root, fakeFilePath[:strings.LastIndexByte(fakeFilePath, os.PathSeparator)])
			require.NoError(t, os.MkdirAll(dir, 0777))
			require.NoError(t, os.WriteFile(path, []byte(testFileContent), 0777))
			bts, readErr := os.ReadFile(path)
			require.NoError(t, readErr)
			require.Equal(t, testFileContent, string(bts))
		}
	}
	cfg, parseCfgErr := parseConfig(filepath.Join(root, defaultConfigFilename))
	require.NoError(t, parseCfgErr)
	result, generateErr := generate(cfg)
	require.NoError(t, generateErr)
	require.Equal(t, len(expectedIndexes), len(result))
	for _, expectedIndexPath := range expectedIndexes {
		path := filepath.Join(root, expectedIndexPath)
		require.Contains(t, result, path)
		bts, readErr := os.ReadFile(path)
		require.NoError(t, readErr)
		content := string(bts)
		expectedContext, exists := expectedIndexContent[expectedIndexPath]
		require.True(t, exists)
		require.Equal(t, expectedContext, content)
	}
}
