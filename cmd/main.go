package main

import (
	"github.com/lmx-Hexagram/conflict/internal/project"
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
	rootCmd.AddCommand(project.CmdNew)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
