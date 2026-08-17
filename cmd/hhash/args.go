package main

import (
	"errors"
	"flag"
	"strings"
	"github.com/hadcrab/hadutils/internal/hash"
)


func collectArgs() (Config, error) {
	algo := flag.String("a", "SHA256", "Algorithm")
	quiet := flag.Bool("q", false, "Print only hash")
	copyToClipboard := flag.Bool("c", false, "Quietly copy hash to clipboard")
	flag.Parse()
	algorithm := hash.Algorithm(strings.ToLower(*algo))
	args := flag.Args()
	if len(args) != 1 {
    	return Config{}, errors.New("exactly one file path is required")
	}
	cfg := Config{
		Path: args[0],
		Algorithm: algorithm,
		Quiet: *quiet,
		CopyToClipboard: *copyToClipboard,
	}
	
	return cfg, nil
}