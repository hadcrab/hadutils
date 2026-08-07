package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/hadcrab/hadutils/internal/hash"
)

type Config struct {
    Path      string
    Algorithm hash.Algorithm
}

func collectArgs() (Config, error) {
	algo := flag.String("a", "", "algorithm")
	flag.Parse()
	algorithm := hash.Algorithm(strings.ToLower(*algo))
	if algorithm == hash.UnknownAlgorithm {
		return Config{}, errors.New("algorithm is required")
	}
	args := flag.Args()
	if len(args) != 1 {
    	return Config{}, errors.New("exactly one file path is required")
	}
	cfg := Config{
		Path: args[0],
		Algorithm: algorithm,
	}
	
	return cfg, nil
}

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

func main() {
 if err := run(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
} 