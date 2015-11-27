// dice.go
//
// Author: Ollivier Robert <roberto@keltia.net>
// Copyright © 2015 by Ollivier Robert

/*
  package describes two new types representing dices and rolls of
  said dices.
 */
package dice

import (
	"crypto/rand"
	"strings"
	"strconv"
	"fmt"
)

const (
	RE_DICE = `(?P<num>\d*)[dD](?P<dice>\d+)`
	RE_BONUS = ``
)

// Valid dice sizes
var (
	VALID_DICES = []int{ 4, 6, 8, 10, 12, 20, 30, 100 }
)

type Dice struct {
	Size int
}

// private func

// check size
func isValid(size int) bool {
    for _, s := range VALID_DICES {
        if size == s {
	 	    return true
		}
    }
    return false
}

// Check for possible bonus
func checkBonus(sRoll string) (int, string) {
	var (
		bonus   int
		diceStr string
	)

	// Look for possible bonus
	parts := strings.Split(sRoll, " ")
	if len(parts) == 2 {
		if parts[1] != "" {
			var err error

			bonus64, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				bonus = 0
			} else {
				bonus = int(bonus64)
			}
		}
	} else {
		bonus = 0
	}
	diceStr = parts[0]
	return bonus, diceStr
}

// Public interface

// Create a new dice
func NewDice (size int) (*Dice) {
	dice := new(Dice)
	if !isValid(size) {
		return nil
	}
	dice.Size = size
	return dice
}

// Throw a dice N times
func (d *Dice) Roll (num int) *Roll {

	b := make([]byte, num)
	_, err := rand.Read(b)
	if err != nil {
		res := new(Roll)
		return res
	}

	res := new(Roll)
	res.Result = make([]int, num)

	for i := 0; i < num; i++  {
		res.Result[i] = (int(b[i]) % d.Size) + 1
		res.Sum += res.Result[i]
	}
	fmt.Printf("%v:%d\n", res.Result, res.Sum)
	return res
}
