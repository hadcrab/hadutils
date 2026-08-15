package main

import (
	"fmt"
	//"os"
	"github.com/hadcrab/hadutils/internal/hash"
	"github.com/hadcrab/hadutils/internal/env"
	"github.com/hadcrab/hadutils/internal/clipboard"
)

func run() error {
	env := env.Collect()
	cfg, err := collectArgs()
	if err != nil {
		return err
	}
	sum, err := hash.Compute(cfg.Path, cfg.Algorithm)
	if err != nil {
		return err
	}
	if cfg.Quiet {
		fmt.Println(sum)
		return nil
	}
	if cfg.CopyToClipboard {
		if err := clipboard.CopyIn(sum, env.Clipboard); err != nil {
			return err
    	}
		return nil
	}
	PrintResult(cfg, sum)
	return nil
}