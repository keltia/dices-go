// roll.go

/*
  This package implements the Roll interface
 */
package dice

import (
	"errors"
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

// Parse a string representing a series of rolls incl. bonus
func ParseRoll (rollStr string) (Result, error) {
	var r Result

	allDices := map[int64]Dices{
		4: NewDices().Append(regularDice(4)),
		6: NewDices().Append(regularDice(6)),
		8: NewDices().Append(regularDice(8)),
		12: NewDices().Append(regularDice(12)),
		20: NewDices().Append(regularDice(20)),
		100: NewDices().Append(regularDice(100)),
	}

	// Normalize
	sRoll := strings.ToUpper(rollStr)

	bonus, diceStr := checkBonus(sRoll)

	// Look at possible dices
	numSize := strings.Split(diceStr, "D")
	if numSize == nil {
		return r, errors.New(fmt.Sprintf("Bad format: %v", rollStr))
	}

	var (
		diceSize int64
		numRoll int64
	)

	if len(numSize) == 2 {
		diceSize, _ = strconv.ParseInt(numSize[1], 10, 32)
		if !isValid(int(diceSize)) {
			return r, errors.New(fmt.Sprintf("Unknown dice: %v", rollStr))
		}
		numRoll, _ = strconv.ParseInt(numSize[0], 10, 32)
		if numRoll == 0 {
			numRoll = 1
		}
	} else {
			return r, errors.New(fmt.Sprintf("Bad format: %v", rollStr))
	}

	var dN Dices

	for i := 0; i < int(numRoll) - 1; i++ {
		r = dN.Append(allDices[diceSize]).Roll(r)
	}
	r = dN.Append(constantDice(bonus)).Roll(r)

	return r, nil
}

// Apply a bonus
func (r *Roll) ApplyBonus(bonus int) {
	r.Bonus = bonus
	r.Sum += bonus
}
