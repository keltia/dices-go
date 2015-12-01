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
func ParseRoll (rollStr string) (*Roll, error) {

	// Normalize
	sRoll := strings.ToUpper(rollStr)

	bonus, diceStr := checkBonus(sRoll)

	// Look at possible dices
	numSize := strings.Split(diceStr, "D")
	if numSize == nil {
		return nil, errors.New(fmt.Sprintf("Bad format: %v", rollStr))
	}

	var (
		diceSize int64
		numRoll int64
	)

	if len(numSize) == 2 {
		diceSize, _ = strconv.ParseInt(numSize[1], 10, 32)
		if !isValid(int(diceSize)) {
			return nil, errors.New(fmt.Sprintf("Unknown dice: %v", rollStr))
		}
		numRoll, _ = strconv.ParseInt(numSize[0], 10, 32)
		if numRoll == 0 {
			numRoll = 1
		}
	} else {
			return nil, errors.New(fmt.Sprintf("Bad format: %v", rollStr))
	}

	d := NewDice(int(diceSize))
	r := d.Roll(int(numRoll))
	r.ApplyBonus(bonus)

	return r, nil
}

// Roll dice(s)
func NewRoll (dice Dice, num int) *Roll {
	res := dice.Roll(num)
	return res
}

// Return the sum of all rolls
func (r *Roll) PrintResult() string {
	return fmt.Sprintf("%d", r.Sum)
}

// Return all rolls
func (r *Roll) PrintDices() string {
	return fmt.Sprintf("%v", r.Result)
}

// Apply a bonus
func (r *Roll) ApplyBonus(bonus int) {
	r.Bonus = bonus
	r.Sum += bonus
}
