package dice

import "testing"

func TestNewDice (t *testing.T) {
	var d *Dice

	validSizes := []int{ 4, 6, 8, 10, 12, 20, 100 }

	// check valid ones
	for _, size := range validSizes {
		d = NewDice(size)
		if d == nil {
			t.Errorf("Valid size %d, should be non-nil", size)
		}
	}

	invalidSizes := []int{ 0, 40, 23 }

	// check valid ones
	for _, size := range invalidSizes {
		d = NewDice(size)
		if d != nil {
			t.Errorf("Invalid size %d, should be nil %v", size, d)
		}
	}
}

func TestRoll (t *testing.T) {
	d := NewDice(20)

	res := d.Roll(2)
	if len(res.Result) != 2 {
		t.Errorf("Bad roll: %v", res.Result)
	}

	res = d.Roll(10)
	if len(res.Result) != 10 {
		t.Errorf("Bad roll: %v", res.Result)
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
	if (res.Bonus != 1) {
		t.Errorf("Error with stored bonus")
	}
}

func TestIsValid(t *testing.T) {
	f := isValid(1)
	if f {
		t.Errorf("Bad size: d%d should be false", 1)
	}

	f = isValid(0)
	if f {
		t.Errorf("Bad size: d%d should be false", 0)
	}

	f = isValid(100)
	if !f {
		t.Errorf("Bad size: d%d should be true", 100)
	}
}

func TestCheckBonus(t *testing.T) {
	d := "3D4"
	bonus, str := checkBonus(d)
	if bonus != 0 {
		t.Errorf("Bad format, bonus should be 0: %s/%d/%s", d, bonus, str)
	}
	if str != d {
		t.Errorf("Bad format, %s should be eq to %s", d, str)
	}

	d = "D6 +1"
	bonus, str = checkBonus(d)
	if bonus != 1 {
		t.Errorf("Bad format, bonus should be 1: %s/%d/%s", d, bonus, str)
	}
	if str != "D6" {
		t.Errorf("Bad format, %s should be eq to %s w/ the bonus", str, d)
	}
}
