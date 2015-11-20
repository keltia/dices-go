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
}

func TestParseRoll (t *testing.T) {
	str := "D20"
	res, err := ParseRoll(str)
	if len(res.Result) != 1 || (res.Sum < 1 || res.Sum > 20) {
		t.Errorf("%v: %d (%d)", err, len(res.Result), res.Sum)
	}

	str = "d20"
	res, err = ParseRoll(str)
	if len(res.Result) != 1 || (res.Sum < 1 || res.Sum > 20) {
		t.Errorf("%v: %d (%d)", err, len(res.Result), res.Sum)
	}

	str = "3d4"
	res, err = ParseRoll(str)
	if len(res.Result) != 3 || (res.Sum < 3 || res.Sum > 12) {
		t.Errorf("Bad string: %v: %d (%d)", err, len(res.Result), res.Sum)
	}

}