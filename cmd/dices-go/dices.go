package main

import (
	"github.com/abiosoft/ishell"
	"github.com/keltia/dices-go"
	"strings"
)

const (
	DicesVersion = "1.2.1"
)

func roll(c *ishell.Context, strs []string) {
	str := strings.Join(strs, " ")
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

func cmdMouv(c *ishell.Context) {
	d := "3D6 -9"

	roll(c, []string{d})
}

func cmdDoom(c *ishell.Context) {
	d := "2D6"

	c.Printf("Thou art Doomed\n")
	roll(c, []string{d})
}

func cmdDice(c *ishell.Context) {
	if len(c.Args) == 0 {
		c.Printf("error: you must specify something (nn)Ddd( +nn)\n")
		return
	}
	roll(c, c.Args)
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
		Func: cmdMouv,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "doom",
		Help: "Dices of Doom",
		Func: cmdDoom,
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
