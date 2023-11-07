package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func createMakeCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "make",
		Short:   "Create the given serialization structure for the target language",
		Long:    `TBD`,
		Example: "wirectl make path/to/definition.toml -l=go -o=.",
		Run:     make,
	}
}

func make(cmd *cobra.Command, args []string) {
	log.Println("creating code for ...")
}
