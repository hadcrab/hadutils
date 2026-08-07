package main

import (
	"fmt"
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
	fmt.Printf(sum)
	return nil
}