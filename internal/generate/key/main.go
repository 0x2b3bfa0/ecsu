package main

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	var rand io.Reader = nil

	if seed, err := hex.DecodeString(os.Getenv("SEED")); err == nil {
		rand = bytes.NewReader(seed)
	}

	public, private, err := ed25519.GenerateKey(rand)
	if err != nil {
		panic(err)
	}

	file := os.Getenv("GOFILE")
	name := strings.TrimSuffix(file, filepath.Ext(file))
	newFile := fmt.Sprintf("%s_key.go", name)

	var keyName = os.Args[1]
	var keyValue []byte

	switch keyName {
	case "PublicKey":
		keyValue = public
	case "PrivateKey":
		keyValue = private
	default:
		panic("unknown key type")
	}

	var keyNumbers []string
	for _, number := range keyValue {
		keyNumbers = append(keyNumbers, strconv.Itoa(int(number)))
	}

	template := `
// Code generated by internal/generate/keys.go DO NOT EDIT.
package %s

import "crypto/ed25519"

var %s = ed25519.%[2]s{%s}
`[1:]

	code := fmt.Sprintf(template, os.Getenv("GOPACKAGE"), keyName, strings.Join(keyNumbers, ", "))

	if err = os.WriteFile(newFile, []byte(code), 0600); err != nil {
		panic(err)
	}
}
