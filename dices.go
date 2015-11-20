package main

import (
	readline "gopkg.in/readline.v1"

	dice "github.com/keltia/dices-go/dice"
	"fmt"
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
		res, err := dice.ParseRoll(str)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Printf("%s = %v (%d)\n", str, res.Result, res.Sum)
	}
}
