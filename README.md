# hadutils

A small collection of Go utilities written for learning purposes.

## Tools

### hhash

Compute file hashes, print them, optionally copy to clipboard (X11/Wayland).

	go build ./cmd/hhash

Usage:

	hhash file.txt                      # SHA-256 by default
	hhash -a md5 file1.txt file2.txt    # pick algorithm, multiple files
	hhash -q -c file.txt                # print only the hash, copy it to clipboard

Supported algorithms: sha256, sha512, sha1, md5.

## Packages

- [cli](./cli) — small generics-based CLI parser, an alternative to `flag` and `cobra`. Work in progress.

## TODO

- [x] hhash
- [ ] cli parser
