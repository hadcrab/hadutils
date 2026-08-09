package hash

import (
	"crypto/sha256"
	"crypto/sha512"
	"crypto/sha1"
	"crypto/md5"
	"encoding/hex"
	"os"
	"fmt"
)

type Algorithm string

const (
	UnknownAlgorithm Algorithm = ""
    SHA256 Algorithm = "sha256"
    SHA512 Algorithm = "sha512"
    SHA1 Algorithm = "sha1"
    MD5 Algorithm = "md5"
)

func Compute(path string, algorithm Algorithm) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	switch algorithm {
		case SHA256:
			sum := sha256.Sum256(data)
			return hex.EncodeToString(sum[:]), nil
		case SHA512:
			sum := sha512.Sum512(data)
			return hex.EncodeToString(sum[:]), nil
		case SHA1:
			sum := sha1.Sum(data)
			return hex.EncodeToString(sum[:]), nil
		case MD5:
			sum := md5.Sum(data)
			return hex.EncodeToString(sum[:]), nil
		default:
    		return "", fmt.Errorf("unsupported algorithm: %s", algorithm)
	}
}
