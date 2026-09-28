package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/hadcrab/hadutils/internal/clipboard"
	"github.com/hadcrab/hadutils/internal/env"
	"github.com/hadcrab/hadutils/internal/hash"
)

func run() error {
	env := env.Collect()
	cfg, err := collectArgs()
	var sums []string
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
			sums = append(sums, sum)
			continue
		}
		if cfg.Quiet {
			fmt.Println(sum)
		} else {
			PrintResult(path, string(cfg.Algorithm), sum)
		}
	}
	if cfg.CopyToClipboard {
		stringToCopy := strings.Join(sums, "\n")
		if err := clipboard.CopyIn(stringToCopy, env.Clipboard); err != nil {
			return err
		}
		return nil
	}
	return nil
}
