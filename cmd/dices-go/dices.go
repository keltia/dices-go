package main

import (
	"fmt"

	"os"
	"github.com/abiosoft/ishell"
	"dices-go/dice"
	"github.com/chzyer/readline"
)

var (
)

func finish() {
	os.Exit(0)
}

func doom() {
	fmt.Printf("you are doomed")
}

func roll(c *ishell.Context) {
	if len(c.Args) == 0 {
		fmt.Printf("error: you must specify something (nn)Ddd( +nn)")
		return
	}
	res, err := dice.ParseRoll(c.Args[0])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%s = %v (%d)\n", c.Args[0], res.List, res.Sum)
		if res.Bonus != 0 {
			fmt.Printf(" Bonus was %d\n", res.Bonus)
		}
	}
}

func init() {
}

func main() {
	c := &readline.Config{Prompt: "Dices> "}
	shell := ishell.NewWithConfig(c)

	shell.AddCmd(&ishell.Cmd{
		Name: "dice",
		Help: "Roll dices",
		Func: roll,
	})

	shell.Run()
}
