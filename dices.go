package main

import (
	dice "github.com/keltia/dices-go/dice"
	"fmt"
)

func main() {
	res, err := dice.ParseRoll("3D6")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("3D6 = %v (%d)\n", res.Result, res.Sum)
}
