package main

import (
	"flag"
	"fmt"
	"log"
)

const (
	indexFilename = "index.ts"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", defaultConfigFilename, "config file path")
	flag.Parse()
	if len(configPath) == 0 {
		configPath = defaultConfigFilename
	}
	cfg, parseCfgErr := parseConfig(configPath)
	if parseCfgErr != nil {
		log.Fatalln(fmt.Errorf("parse config: %w", parseCfgErr))
	}
	result, generateErr := generate(cfg)
	if generateErr != nil {
		log.Fatalln(fmt.Errorf("generate: %w", generateErr))
	}
	for _, index := range result {
		log.Printf("index generated: %s\n", index)
	}
}
