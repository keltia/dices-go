package main

import (
	dices "github.com/keltia/dices/dice"
	"fmt"
)

func main() {
	res, err := dices.ParseRoll("3D6")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("3D6 = %v (%d)\n", res.Result, res.Sum)
}
