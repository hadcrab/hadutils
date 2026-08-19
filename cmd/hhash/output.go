package main

import (
	"fmt"
)

func PrintResult(path, algo, hash string) {
	fmt.Printf("File: %s\n", path)
    fmt.Printf("Algorithm: %s\n", algo)
    fmt.Printf("Hash: %s\n", hash)
}