package main

import (
	"os"

	"github.com/sukanthm/site24x7beat/cmd"

	_ "github.com/sukanthm/site24x7beat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
