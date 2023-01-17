package main

import (
	"github.com/go-monsters/sullivan/cmd/commands"
	_ "github.com/go-monsters/sullivan/cmd/commands/api"
)

func main() {
	commands.Execute()
}
