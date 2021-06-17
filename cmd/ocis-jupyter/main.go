package main

import (
	"os"

	"github.com/anaswaratrajan/ocis-jupyter/pkg/command"
)

func main() {
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
