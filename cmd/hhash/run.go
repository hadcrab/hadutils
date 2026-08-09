package main

import (
	"github.com/hadcrab/hadutils/internal/hash"
)

func run() error {
	cfg, err := collectArgs()
	if err != nil {
		return err
	}
	sum, err := hash.Compute(cfg.Path, cfg.Algorithm)
	if err != nil {
		return err
	}
	PrintResult(cfg, sum)
	return nil
}