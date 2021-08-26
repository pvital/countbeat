package main

import (
	"os"

	"github.com/pvital/countbeat/cmd"

	_ "github.com/pvital/countbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
