package main

import (
	"fmt"
	"os"

	"github.com/nathanielwheeler/go-cyoa/server"
)

func main() {
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "server error:\n\t%s\n", err)
		os.Exit(1)
	}
}
