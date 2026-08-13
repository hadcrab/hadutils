package main

import (
	"errors"
	"flag"
	"strings"
	"github.com/hadcrab/hadutils/internal/hash"
)


func collectArgs() (Config, error) {
	algo := flag.String("a", "", "algorithm")
	quiet := flag.Bool("q", false, "quiet")
	flag.Parse()
	algorithm := hash.Algorithm(strings.ToLower(*algo))
	isQuiet := *quiet
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
		Quiet: isQuiet,
	}
	
	return cfg, nil
}