package main

import (
	"os"
	"strings"
)

func main() {
	expletive := "y"

	if len(os.Args) > 1 {
		expletive = strings.Join(os.Args[1:], " ")
	}

	b := []byte(expletive + "\n")

	for {
		_, err := os.Stdout.Write(b)
		if err != nil {
			os.Stderr.Write([]byte(err.Error()))
		}
	}
}
