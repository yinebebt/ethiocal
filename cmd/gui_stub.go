//go:build noui

package cmd

import (
	"fmt"
	"os"
)

func runGUI() {
	fmt.Println("GUI is not available in this build. Use CLI subcommands or --server instead.")
	os.Exit(1)
}
