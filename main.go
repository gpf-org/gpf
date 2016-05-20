package main

import (
	"os"

	"github.com/gpf-org/gpf/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
