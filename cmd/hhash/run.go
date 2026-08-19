package main

import (
	"errors"
	"fmt"
	//"os"
	"github.com/hadcrab/hadutils/internal/hash"
	// "github.com/hadcrab/hadutils/internal/env"
	// "github.com/hadcrab/hadutils/internal/clipboard"
)

func run() error {
	// env := env.Collect()
	cfg, err := collectArgs()
	if err != nil {
		return err
	}

	for _, path := range cfg.Path {
		sum, err := hash.Compute(path, cfg.Algorithm)
		if err != nil {
			return err
		}
		if cfg.CopyToClipboard {
			//todo add support for this
			return errors.New("cannot use --copy with multiple files ")
    	}
		if cfg.Quiet {
			fmt.Println(sum)
		} else {
			PrintResult(path, string(cfg.Algorithm), sum)
		}
	}
	return nil
}