// roll.go

/*
  This package implements the Roll interface
*/
package dice

import (
	"fmt"
	"strconv"
	"strings"
)

// roll is more complex
type Roll struct {
	Result []int
	Sum    int
	Bonus  int
	Tag    string
}

// Check for possible bonus
func checkBonus(sRoll string) (bonus int, diceStr string) {

	// Look for possible bonus
	parts := strings.Split(sRoll, " ")
	if len(parts) == 2 {
		if parts[1] != "" {
			b, err := strconv.Atoi(parts[1])
			if err != nil {
				bonus = 0
			} else {
				bonus = b
			}
		}
	} else {
		bonus = 0
	}
	diceStr = parts[0]
	return
}

// Parse a string representing a series of rolls incl. bonus
func ParseRoll(rollStr string) (Result, error) {
	var r Result

	allDices := map[int]Dices{
		4:   NewDices().Append(regularDice(4)),
		6:   NewDices().Append(regularDice(6)),
		8:   NewDices().Append(regularDice(8)),
		10:  NewDices().Append(regularDice(10)),
		12:  NewDices().Append(regularDice(12)),
		20:  NewDices().Append(regularDice(20)),
		100: NewDices().Append(regularDice(100)),
	}

	// Normalize
	sRoll := strings.ToUpper(rollStr)

	bonus, diceStr := checkBonus(sRoll)

	// Look at possible dices
	numSize := strings.Split(diceStr, "D")
	if numSize == nil {
		return r, fmt.Errorf("Bad format: %v", rollStr)
	}

	var (
		diceSize int
		numRoll  int
	)

	if len(numSize) == 2 {
		diceSize, _ = strconv.Atoi(numSize[1])
		if !isValid(diceSize) {
			return r, fmt.Errorf("Unknown dice: %v", rollStr)
		}
		numRoll, _ = strconv.Atoi(numSize[0])
		if numRoll == 0 {
			numRoll = 1
		}
	} else {
		return r, fmt.Errorf("Bad format: %v", rollStr)
	}

	var dN Dices

	for i := 0; i <= int(numRoll)-1; i++ {
		r = dN.Append(allDices[diceSize]).Roll(r)
	}
	r = dN.Append(constantDice(bonus)).Roll(r)
	r.Bonus = bonus

	return r, nil
}
