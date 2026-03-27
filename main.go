package main

import (
	_ "embed"

	"github.com/InkShaStudio/go-command"
	common "github.com/summink/sumi-common-command"
)

//go:embed manifest.json
var manifest []byte

func init() {
	if manifest, err := common.LoadManifestByByte(manifest); err != nil {
		panic(err.Error())
	} else {
		common.LoadManifest(manifest)
	}
}

func mainCommand() *command.SCommand {
	// Override the func for customization

	target := command.
		NewCommandArg[string]("target").
		ChangeDescription("say hello target").
		ChangeValue("world!")

	cmd := command.
		NewCommand("hello").
		ChangeDescription("say hello").
		AddArgs(target).
		RegisterHandler(func(cmd *command.SCommand) {
			println("hello " + target.Value)
		})

	return cmd
}

func main() {
	cmd := mainCommand()

	caller := command.RegisterCommand(common.WithCommand(cmd))

	caller.Execute()
}
