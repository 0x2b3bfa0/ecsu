package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	bytes := make([]byte, 32)
	if count, err := rand.Read(bytes); count != len(bytes) {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(bytes))
}
