package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	expletive := "y"

	if len(os.Args) > 1 {
		expletive = strings.Join(os.Args[1:], " ")
	}

	b := []byte(expletive + "\n")

	stdout := bufio.NewWriterSize(os.Stdout, os.Getpagesize())
	stderr := os.Stderr

	for {
		_, err := stdout.Write(b)
		if err != nil {
			stderr.Write([]byte("yuh: " + err.Error()))
		}
	}
}
