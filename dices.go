package main

import (
	"fmt"

	readline "gopkg.in/readline.v1"

	dice "github.com/keltia/dices-go/dice"
)

func main() {
	rl, err := readline.New("Dices> ")
	if err != nil {
		fmt.Printf("Error reading line: %v\n", err)
	}
	defer rl.Close()

	for {
		str, err := rl.Readline()
		if err != nil { // EOF
			break
		}

		// For the fun
		if str == "doom" || str == "DOOM" {
			str = "2D6"
		}

		res, err := dice.ParseRoll(str)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("%s = %v (%d)\n", str, res.Result, res.Sum)
			if res.Bonus != 0 {
				fmt.Printf(" Bonus was %d\n", res.Bonus)
			}
		}
	}
}
