package main

import (
	"fmt"

	readline "gopkg.in/readline.v1"

	dice "github.com/keltia/dices-go/dice"
	"strings"
)

func main() {
	rl, err := readline.New("Dices> ")
	if err != nil {
		fmt.Printf("Error reading line: %v\n", err)
	}
	defer rl.Close()

	// Save (and reload) our history
	rl.Config.HistoryFile = ".history"

	for {
		str, err := rl.Readline()
		if err != nil { // EOF
			break
		}

		// Prepare
		str = strings.TrimSpace(str)

		// For the fun
		if str == "doom" || str == "DOOM" {
			str = "2D6"
		}

		// Parse the thing
		res, err := dice.ParseRoll(str)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("%s = %v = %d", str, res.List, res.Sum)
			if res.Bonus != 0 {
				fmt.Printf(" (%d %+d)\n", res.Sum - res.Bonus, res.Bonus)
			} else {
				fmt.Println()
			}
		}
	}
}
