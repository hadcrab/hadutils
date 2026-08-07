package main

import (
	"errors"
	"flag"
	"strings"
	"github.com/hadcrab/hadutils/internal/hash"
)


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