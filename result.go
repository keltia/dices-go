// result.go
//
// Author: Ollivier Robert <roberto@keltia.net>
// Copyright © 2015-2018 by Ollivier Robert

package dice

// Result of a roll
type Result struct {
	List  []int
	Sum   int
	Bonus int
	Size  int
}

// Append adds a constantDice (int) to the roll (i.e. bonus
func (r Result) Append(v int) Result {
	return Result{append(r.List, v), r.Sum + v, r.Bonus, r.Size}
}

// Merge everything incl. bonus
func (r Result) Merge(r1 Result) Result {
	return Result{append(r.List, r1.List...), r.Sum + r1.Sum, r.Bonus + r1.Bonus, r.Size}
}

func (r Result) Natural() bool {
	return len(r.List) == 1 && r.Sum == r.Size
}

