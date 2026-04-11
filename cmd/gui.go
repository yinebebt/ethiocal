//go:build !noui

package cmd

import "github.com/yinebebt/ethiocal/gui"

func runGUI() {
	gui.Run()
}
