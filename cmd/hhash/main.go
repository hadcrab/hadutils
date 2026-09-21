package main

import (
	"fmt"
	"github.com/hadcrab/hadutils/internal/env"
	"github.com/hadcrab/hadutils/internal/hash"
	"os"
)

type Config struct {
	Path            []string
	Algorithm       hash.Algorithm
	Quiet           bool
	CopyToClipboard bool
}

func main() {
	env.Collect()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
