package main

import (
	"github.com/abiosoft/ishell"
	"github.com/keltia/dices-go"
	"strings"
)

const (
	DicesVersion = "1.2.1"
)

func roll(c *ishell.Context) {
	str := strings.Join(c.Args, " ")
	res, err := dice.ParseRoll(str)
	if err != nil {
		c.Printf("Error: %v\n", err)
	} else {
		c.Printf("%s = %v (%d)\n", str, res.List, res.Sum)
		if res.Bonus != 0 {
			c.Printf(" Bonus was %d\n", res.Bonus)
		}
	}

}

func cmdDice(c *ishell.Context) {
	if len(c.Args) == 0 {
		c.Printf("error: you must specify something (nn)Ddd( +nn)\n")
		return
	}
	roll(c)
}

func main() {

	shell := ishell.New()
	shell.Printf("Dices version %s\n", DicesVersion)
	shell.SetPrompt("Dices> ")
	shell.SetHistoryPath(".history")

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
	shell.Run()
}
