package main

import (
	"bufio"
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/0x2b3bfa0/ecsu/internal/timestep"
)

//go:generate go run ../../internal/generate/key/main.go PublicKey

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s command [argument...]\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Fprint(os.Stderr, "time-based code: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	code := scanner.Text()

	if err := verify(code); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := run(os.Args[1], os.Args[2:]...); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func verify(code string) error {
	signature, err := base64.StdEncoding.DecodeString(code)
	if err != nil {
		return err
	}

	if !ed25519.Verify(PublicKey, timestep.Get(), signature) {
		return errors.New("invalid signature")
	}

	return nil
}

func run(command string, arguments ...string) error {
	path, err := exec.LookPath(command)
	if err != nil {
		return err
	}

	if err := syscall.Setuid(0); err != nil {
		return err
	}

	if err := syscall.Setgid(0); err != nil {
		return err
	}

	if err := syscall.Exec(path, arguments, os.Environ()); err != nil {
		return err
	}

	return nil // unreachable
}
