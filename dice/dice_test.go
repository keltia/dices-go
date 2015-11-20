package dice

import "testing"

func TestNewDice (t *testing.T) {

	d := NewDice(0)
	if d != nil {
		t.Errorf("Can't have a dice with less than 4: %v", d)
	}

	d = NewDice(40)
	if d != nil {
		t.Errorf("Can't have dice bigger than 20: %v", d)
	}
}

func TestApplyBonus (t *testing.T) {
	d := NewDice(20)

	res := d.Roll(1)
	sum := res.Sum
	res.ApplyBonus(1)
	if (sum + 1) != res.Sum {
		t.Errorf("Error with bonus")
	}
}

