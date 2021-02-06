package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func init() {}

func TestMain(t *testing.T) {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
