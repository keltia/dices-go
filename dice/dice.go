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
	"math/rand"
	"strings"
	"strconv"
	"time"
)

const (
	RE_DICE = `(?P<num>\d*)[dD](?P<dice>\d+)`
	RE_BONUS = ``
)

type Dice struct {
	Size int
}

// Create a new dice
func NewDice (size int) (*Dice) {
	rand.Seed(time.Now().Unix())
	dice := new(Dice)
	if size < 4 || size > 20 {
		return nil
	}
	dice.Size = size
	return dice
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

// Throw a dice N times
func (d *Dice) Roll (num int) *Roll {

	res := new(Roll)
	res.Result = make([]int, num)

	for i := 0; i < num; i++  {
		res.Result[i] = rand.Intn(d.Size - 1) + 1
		res.Sum += res.Result[i]
	}
	return res
}
