package main

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
	Long: `Ethiocal provides fasting and holiday dates based on the Ethiopian Orthodox
church calendar, and converts dates between the Ethiopian and Gregorian calendars.

Use subcommands (bahir, convert) for CLI access, or --server to start the HTTP API.
The graphical app ships separately as the Ethiocal desktop and mobile build.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		if server {
			fmt.Println("Running in server mode")
			handler.Init()
			return nil
		}
		return cmd.Help()
	},
}

// Execute runs the root command: bahir/convert subcommands, --server for the
// HTTP API, no args prints usage.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVar(&server, "server", false, "Start the HTTP API server")

	rootCmd.AddCommand(bahirCmd)
	rootCmd.AddCommand(convertCmd)
}
