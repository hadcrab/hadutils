package main

import (
	"fmt"
	"os"

	"github.com/hadcrab/hadutils/internal/clipboard"
	"github.com/hadcrab/hadutils/internal/env"
	"github.com/hadcrab/hadutils/internal/hash"
)

func run() error {
	env := env.Collect()
	cfg, err := collectArgs()
	if err != nil {
		return err
	}
	
	for _, path := range cfg.Path {
		var sum string
		var err error
		if path == "-" {
			sum, err = hash.ComputeReader(os.Stdin, cfg.Algorithm)
		} else {
			sum, err = hash.Compute(path, cfg.Algorithm)
		}
		if err != nil {
			return err
		}
		if cfg.CopyToClipboard {
			clipboard.CopyIn(sum, env.Clipboard)
			// todo: make it work with multiple files
			return nil
    	}
		if cfg.Quiet {
			fmt.Println(sum)
		} else {
			PrintResult(path, string(cfg.Algorithm), sum)
		}
	}
	return nil
}