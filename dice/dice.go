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
	mrand "math/rand"
)

const (
	RE_DICE = `(?P<num>\d*)[dD](?P<dice>\d+)`
	RE_BONUS = ``
	SEED_SIZE = 8
)

// Valid dice sizes
var (
	VALID_DICES = []int{ 4, 6, 8, 10, 12, 20, 30, 100 }
)

// Private func

// Generate bounded values in a fair way
// Inspired by @ns_m code on IRC
func biasedCoin(p float64) bool {
	return mrand.Float64() < p
}

// We will use the values generated by crypto/rand to derive our "rolls"
func internalRoll(sides int) int {
	i := 0
	for {
		if biasedCoin(1 / float64(sides - i)) {
			return i + 1
		}
		i++
	}
}

// check size
func isValid(size int) bool {
    for _, s := range VALID_DICES {
        if size == s {
	 	    return true
		}
    }
    return false
}

// "key schedule" to seed the random generator
func keySchedule(seed int) int64 {

	// Seed mrand.Seed() with these
	b := make([]byte, seed)
	if _, err := rand.Read(b); err != nil {
		return -1
	}

	// Now, b is 8 bytes long, generate a 64 bit value
	acc := int64(0)
	for _, i := range b {
		acc = int64(acc * 16) + int64(i)
	}
	return acc
}

// Roll interface (instead of a separate type
type Dice interface {
	Roll(r Result) Result
}

// A set of dices
type Dices []Dice

// Result of a roll
type Result struct {
	List []int
	Sum  int
	Bonus int
}

// variable dice
type regularDice int

// This is a Dice()
func (nd regularDice) Roll(r Result) Result {
    return r.Append(internalRoll(int(nd)))
}

// Used to represent modifiers like bonus
type constantDice int

// This is a Dice() too
func (cd constantDice) Roll(r Result) Result {
	return r.Append(int(cd))
}

// Open-ended dices
type openDice struct {
    threshold int
    d Dice
}

// An openDice is a Dice()
func (td *openDice) Roll(r Result) Result {
    if r.Sum != td.threshold {
        return r
    }
    return r.Merge(td.d.Roll(Result{}))
}

// API

// Make a new set of rolls
func NewDices() Dices {
	return make([]Dice, 0, 10)
}

// That's how we do multiple rolls
func (set Dices) Append(d ...Dice) Dices {
	return append(set, d...)
}

// AleaJactaEst — the actual rolling
func (d Dices) Roll(r Result) Result {
    r1 := r

	// Seed the thing
	mrand.Seed(keySchedule(SEED_SIZE))

    for _,dice := range d {
        r1 = dice.Roll(r1)
    }
    return r1
}

// Add a constantDice (int) to the roll (i.e. bonus
func (r Result) Append(v int) Result {
	return Result{append(r.List, v), r.Sum + v, r.Bonus}
}

// Merge everything incl. bonus
func (r Result) Merge(r1 Result) Result {
    return Result { append(r.List, r1.List...) , r.Sum+r1.Sum, r.Bonus+r1.Bonus }
}
