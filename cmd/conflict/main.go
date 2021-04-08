package main

import (
	"github.com/go-conflict/conflict/internal/commands"
	"github.com/spf13/cobra"
	"log"
)

var (
	version = "v0.0.1"

	rootCmd = &cobra.Command{
		Use:     "conflict",
		Short:   "conflict",
		Long:    `conflict`,
		Version: version,
	}
)

func init() {
	rootCmd.AddCommand(commands.NewNewProjectCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
