package api

import (
	"github.com/go-monsters/sullivan/cmd/commands"
	"github.com/spf13/cobra"
)

var (
	apiCommand = &cobra.Command{
		Use:   "api [appName]",
		Short: "Creates a monster API application",
		Long: `The command 'api' creates a Beego API application.
  now default support generate a go modules project.

  {{"Example:"|bold}}
      $ mike api [appname] [-tables=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]  [-gopath=false] [-sullivan=v1.12.3]

  If 'conn' argument is empty, the command will generate an example API application. Otherwise the command
  will connect to your database and generate models based on the existing tables.

  The command 'api' creates a folder named [appname] with the following structure:

	    ├── main.go
	    ├── go.mod
	    ├── {{"conf"|foldername}}
	    │     └── app.conf
	    ├── {{"controllers"|foldername}}
	    │     └── object.go
	    │     └── user.go
	    ├── {{"routers"|foldername}}
	    │     └── router.go
	    ├── {{"tests"|foldername}}
	    │     └── default_test.go
	    └── {{"models"|foldername}}
	          └── object.go
	          └── user.go
`,
		Run: createApi,
	}
)

func init() {
	commands.RootCmd.AddCommand(apiCommand)
}

func createApi(cmd *cobra.Command, args []string) {

}
