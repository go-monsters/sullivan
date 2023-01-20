package main

import (
	"log"

	"github.com/go-monsters/sullivan/internal/project"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mike",
	Short: "mike: An elegant toolkit for using with Monster framework",
	Long:  `mike: An elegant toolkit for using with Monster framework`,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
