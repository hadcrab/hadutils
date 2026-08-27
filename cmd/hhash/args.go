package main

import (
	"errors"
	"strings"
	"os"

	"github.com/hadcrab/hadutils/cli"
	"github.com/hadcrab/hadutils/internal/hash"
)

const defaultAlgorithm = "sha256" 

func collectArgs() (Config, error) {
	input := os.Args
	algo := cli.String("algorithm", "a").Default(defaultAlgorithm).Description("Algorithm")
	quiet := cli.Bool("quiet", "q").Description("Print only hash")
	copyToClipboard := cli.Bool("copy", "c").Description("Copy hash quietly to clipboard")
	args, err := cli.Parse(input[1:], algo, quiet, copyToClipboard)
	algorithm := hash.Algorithm(strings.ToLower(algo.Value()))
	if err != nil {
		return Config{}, err
	}
	if len(args) < 1 {
    	return Config{}, errors.New("At least one file path is required")
	}
	cfg := Config{
		Path: args,
		Algorithm: algorithm,
		Quiet: quiet.Value(),
		CopyToClipboard: copyToClipboard.Value(),
	}
	
	return cfg, nil
}