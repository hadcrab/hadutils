package main

import (
	"fmt"
	"os"
	"github.com/hadcrab/hadutils/internal/hash"
)

type Config struct {
    Path      string
    Algorithm hash.Algorithm
}

func main() {
 if err := run(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
} 