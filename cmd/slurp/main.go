package main

import (
	"context"
	"os"

	"github.com/charmbracelet/fang"

	"github.com/r0whammer/slurp/internal/cli/commands"
)

func main() {
	rootCmd := commands.NewRootCmd()
	if err := fang.Execute(context.Background(), rootCmd); err != nil {
		os.Exit(1)
	}
}
