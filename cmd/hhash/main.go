package main

import (
	"fmt"
	"os"
	"github.com/hadcrab/hadutils/internal/hash"
	"github.com/hadcrab/hadutils/internal/env"
)

type Config struct {
    Path      string
    Algorithm hash.Algorithm
    Quiet      bool
    CopyToClipboard bool
}

func main() {
	env.Collect()
	if err := run(); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
} 