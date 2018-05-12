package main

import (
	"github.com/abiosoft/ishell"
	"github.com/keltia/dices-go"
	"strings"
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
	if res.Natural() {
		c.Printf("Natural %d!\n", res.Size)
	}
}

func cmdDice(c *ishell.Context) {
	if len(c.Args) == 0 {
		c.Printf("error: you must specify something (nn)Ddd( +nn)\n")
		return
	}
	roll(c)
}
