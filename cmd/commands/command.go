package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "mike",
		Short: "A generator for Monster framework",
		Long: `Mike is a CLI tools for Monster framework
This application is a tool to generate the needed files
to quickly create a business application.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				err := cmd.Help()
				if err != nil {
					//TODO: log
				}
				os.Exit(0)
			}
		},
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
}
