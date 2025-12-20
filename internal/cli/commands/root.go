package commands

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "example",
		Short: "A simple example program!",
	}
	return rootCmd
}
