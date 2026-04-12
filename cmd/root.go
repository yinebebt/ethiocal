package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yinebebt/ethiocal/handler"
)

var server bool

var rootCmd = &cobra.Command{
	Use:   "ethiocal",
	Short: "Ethiocal — Ethiopian Calendar (ባሕረ-ሐሳብ) and date converter",
	Long: `Ethiocal is used to get fasting and holiday dates within a year based on Ethiopian
Orthodox church calendar. It also converts dates between Ethiopian and Gregorian calendars.

Running without arguments launches the GUI. Use subcommands for CLI mode,
or --server to start the HTTP API.`,
	Run: func(cmd *cobra.Command, args []string) {
		if server {
			fmt.Println("Running in server mode")
			handler.Init()
			return
		}
		runGUI()
	},
}

// Execute runs the root command. When invoked with no arguments, it launches
// the GUI. Subcommands (bahir, convert) provide CLI access. The --server flag
// starts the HTTP API.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVar(&server, "server", false, "Start the HTTP API server")

	rootCmd.AddCommand(bahirCmd)
	rootCmd.AddCommand(convertCmd)
}
