package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type export string

const (
	exportSingle  export = "single"
	exportDefault export = "default"
	exportAll     export = "all"
)

const (
	defaultConfigFilename = "tsig.json"
)

type config struct {
	Root    string   `json:"root"`
	Options []option `json:"options"`
}

type option struct {
	Dir       string `json:"dir"`
	Export    export `json:"export"`
	Recursive bool   `json:"recursive"`
}

func parseConfig(configPath string) (*config, error) {
	cfgBts, readCfgErr := os.ReadFile(configPath)
	if readCfgErr != nil {
		return nil, fmt.Errorf("read config file: %w", readCfgErr)
	}
	cfg := new(config)
	if unmarshalErr := json.Unmarshal(cfgBts, cfg); unmarshalErr != nil {
		return nil, fmt.Errorf("unmarshal config file: %w", unmarshalErr)
	}
	return cfg, nil
}
