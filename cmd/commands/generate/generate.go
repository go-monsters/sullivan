package generate

import (
	"github.com/go-monsters/sullivan/cmd/commands"
	"github.com/spf13/cobra"
)

var (
	generateCommand = &cobra.Command{
		Use:   "generate [command]",
		Short: "Source code generator",
		Long: `▶ {{"To scaffold out your entire application:"|bold}}

     $ bee generate scaffold [scaffoldname] [-fields="title:string,body:text"] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]

  ▶ {{"To generate a Model based on fields:"|bold}}

     $ bee generate model [modelname] [-fields="name:type"]

  ▶ {{"To generate a controller:"|bold}}

     $ bee generate controller [controllerfile]

  ▶ {{"To generate a CRUD view:"|bold}}

     $ bee generate view [viewpath]

  ▶ {{"To generate a migration file for making database schema updates:"|bold}}

     $ bee generate migration [migrationfile] [-fields="name:type"]

  ▶ {{"To generate swagger doc file:"|bold}}

     $ bee generate docs

    ▶ {{"To generate swagger doc file:"|bold}}

     $ bee generate routers [-ctrlDir=/path/to/controller/directory] [-routersFile=/path/to/routers/file.go] [-routersPkg=myPackage]

  ▶ {{"To generate a test case:"|bold}}

     $ bee generate test [routerfile]

  ▶ {{"To generate appcode based on an existing database:"|bold}}

     $ bee generate appcode [-tables=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"] [-level=3]
`,
		Run: generateCode,
	}
)

func init() {
	commands.RootCmd.AddCommand(generateCommand)
}

func generateCode(cmd *cobra.Command, args []string) {

}
