package main

import (
	"fmt"
)

func PrintResult(cfg Config, hash string) {
	fmt.Printf("File: %s\n", cfg.Path)
    fmt.Printf("Algorithm: %s\n", cfg.Algorithm)
    fmt.Printf("Hash: %s\n", hash)
}