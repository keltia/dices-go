package main

import (
	"fmt"

	"github.com/keltia/dices-go"
	"github.com/abiosoft/ishell"
    "strings"
)

func doom(c *ishell.Context) {
	fmt.Printf("you are doomed\n")
	res, err := dice.ParseRoll("3D6")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%s = %v (%d)\n", c.Args[0], res.List, res.Sum)
		if res.Bonus != 0 {
			fmt.Printf(" Bonus was %d\n", res.Bonus)
		}
	}
}

func roll(c *ishell.Context) {
	if len(c.Args) == 0 {
		fmt.Printf("error: you must specify something (nn)Ddd( +nn)")
		return
	}
	res, err := dice.ParseRoll(strings.TrimSpace(c.Args[0]))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%s = %v (%d)\n", c.Args[0], res.List, res.Sum)
		if res.Bonus != 0 {
			fmt.Printf(" Bonus was %d\n", res.Bonus)
		}
	}
}

func main() {

	shell := ishell.New()
    shell.SetPrompt("Dices> ")

	shell.AddCmd(&ishell.Cmd{
		Name: "dice",
		Help: "Roll dices",
		Func: roll,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "doom",
		Help: "Dices of Doom",
		Func: doom,
	})
	shell.Run()
}
