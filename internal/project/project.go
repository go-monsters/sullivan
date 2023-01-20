package project

import "github.com/spf13/cobra"

var CmdNew = &cobra.Command{
	Use:   "new",
	Short: "Create a service template",
	Long:  "Create a service project using the repository template. Example: mike new helloworld",
}
