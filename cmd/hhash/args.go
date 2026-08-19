package main

import (
	"errors"
	"flag"
	"strings"
	"github.com/hadcrab/hadutils/internal/hash"
)

const defaultAlgorithm = "SHA256" 

func collectArgs() (Config, error) {
	var (
		algo string
		quiet bool
		copyToClipboard bool
	)
	flag.StringVar(&algo, "a", defaultAlgorithm, "Algorithm")
	flag.StringVar(&algo, "algorithm", defaultAlgorithm, "Algorithm")
	flag.BoolVar(&quiet, "q", false, "Print only hash")
	flag.BoolVar(&quiet, "quiet", false, "Print only hash")
	flag.BoolVar(&copyToClipboard, "c", false, "Quietly copy hash to clipboard")
	flag.BoolVar(&copyToClipboard, "copy", false, "Quietly copy hash to clipboard")
	flag.Parse()
	algorithm := hash.Algorithm(strings.ToLower(algo))
	args := flag.Args()
	
	if len(args) != 1 {
    	return Config{}, errors.New("exactly one file path is required")
	}
	cfg := Config{
		Path: args[0],
		Algorithm: algorithm,
		Quiet: quiet,
		CopyToClipboard: copyToClipboard,
	}
	
	return cfg, nil
}