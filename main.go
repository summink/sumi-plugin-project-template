package main

import "github.com/InkShaStudio/go-command"

func helloWorld() *command.SCommand {
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
	cmd := command.RegisterCommand(helloWorld())

	cmd.Execute()
}
