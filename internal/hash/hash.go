package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"fmt"
)

type Algorithm string

const (
	UnknownAlgorithm Algorithm = ""
    SHA256 Algorithm = "sha256"
)

func Compute(path string, algorithm Algorithm) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	switch algorithm {
		case SHA256:
			sum := sha256.Sum256(data)
			hash := hex.EncodeToString(sum[:])
			return hash, nil
		default:
    		return "", fmt.Errorf("unsupported algorithm: %s", algorithm)
	}
}
