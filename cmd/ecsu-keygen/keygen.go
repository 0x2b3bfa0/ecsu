package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"

	"github.com/0x2b3bfa0/ecsu/internal/timestep"
)

//go:generate go run ../../internal/generate/key/main.go PrivateKey

func main() {
	signature := ed25519.Sign(PrivateKey, timestep.Get())
	fmt.Println(base64.StdEncoding.EncodeToString(signature))
}
