package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
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
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var h hash.Hash
	switch algorithm {
		case SHA256:
			h = sha256.New()
		case SHA512:
			h = sha512.New()
		case SHA1:
			h = sha1.New()
		case MD5:
			h = md5.New()			
		default:
    		return "", fmt.Errorf("unsupported algorithm: %s", algorithm)
	}
	_, err = io.Copy(h, file)
	if err != nil {
		return "", err
	}
    return hex.EncodeToString(h.Sum(nil)), nil	
}
