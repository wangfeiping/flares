package main

import (
	"os"

	"github.com/wangfeiping/flares/cmd/flaresd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
