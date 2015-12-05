package dice

import "testing"

func TestNewDices (t *testing.T) {
	var d Dices
	var r Result

	validSizes := []int{ 4, 6, 8, 10, 12, 20, 100 }

	// check valid ones
	for _, size := range validSizes {
		r = NewDices().Append(regularDice(size)).Roll(r)
	}

	invalidSizes := []int{ 0, 40, 23 }

	// check valid ones
	for _, size := range invalidSizes {
		d = NewDices().Append(regularDice(size))
		if d != nil {
			t.Errorf("Invalid size %d, should be nil %v", size, d)
		}
	}
}

func TestRoll (t *testing.T) {
	var r Result

	res := NewDices().Append(regularDice(20)).Roll(r)

	if len(res.List) != 2 {
		t.Errorf("Bad roll: %v", res.Sum)
	}

	res = NewDices().Append(regularDice(10)).Roll(r)
	if len(res.List) != 10 {
		t.Errorf("Bad roll: %v", res.Sum)
	}

	res = NewDices().Append(regularDice(0)).Roll(r)
	if len(res.List) != 0 || res.Sum != -1 {
		t.Errorf("Invalid roll for %d: %v", 0, res.Sum)
	}
}

func TestApplyBonus (t *testing.T) {
	var r Result

	r = NewDices().Append(regularDice(20)).Roll(r)

	sum := r.Sum
	//r.ApplyBonus(1)
	if (sum + 1) != r.Sum {
		t.Errorf("Error with bonus")
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

	f = isValid(23)
	if f {
		t.Errorf("Bad size: d%d should be false", 23)
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
