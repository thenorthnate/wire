package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "wirectl",
		Short: "wirectl is a tool for interacting with the wire serialization format",
		Long:  `TBD`,
	}
	verbose = false
)

func init() {
	// cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Set the verbosity of logs (default is not verbose)")

	rootCmd.AddCommand(
		createMakeCommand(),
	)
}

func initConfig() {}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
