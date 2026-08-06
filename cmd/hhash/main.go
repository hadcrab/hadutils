package main

import (
	"io"
	"os"
	"fmt"	
	// "github.com/hadcrab/hadutils/internal/hash"
)

func readFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	
	defer file.Close()
	
	return string(data), nil
} 

func main() {
	file, err := readFile("testfile.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fileTest, err := os.ReadFile("testfile.txt")
	
	fmt.Printf(file)
	fmt.Printf(string(fileTest))
} 