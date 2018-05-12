package main

import (
	"github.com/abiosoft/ishell"
)

const (
	DicesVersion = "1.2.2"
)

func setupCLI(shell *ishell.Shell) {
	shell.AddCmd(&ishell.Cmd{
		Name: "dice",
		Help: "Roll dices",
		Func: cmdDice,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "mouv",
		Help: "Move dices",
		Func: func(c *ishell.Context) {
			c.Args = append(c.Args, "3D6", "-9")
			roll(c)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "doom",
		Help: "Dices of Doom",
		Func: func(c *ishell.Context) {
			c.Printf("Thou art Doomed\n")
			c.Args = append(c.Args, "2D6")
			roll(c)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "version",
		Help: "Display version",
		Func: func(c *ishell.Context) {
			c.Printf("Version: %s\n", DicesVersion)
		},
	})
}

func main() {

	shell := ishell.New()
	shell.Printf("Dices version %s\n", DicesVersion)
	shell.SetPrompt("Dices> ")
	shell.SetHistoryPath(".history")

	setupCLI(shell)

	shell.Run()
}
