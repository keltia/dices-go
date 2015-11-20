package dice

import (
	"testing"
)

func TestParseRoll (t *testing.T) {

	str := "D20"
	res, err := ParseRoll(str)
	if len(res.Result) != 1 || (res.Sum < 1 || res.Sum > 20) {
		t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.Result), res.Sum)
	}

	str = "d20"
	res, err = ParseRoll(str)
	if len(res.Result) != 1 || (res.Sum < 1 || res.Sum > 20) {
		t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.Result), res.Sum)
	}

	str = "3d4"
	res, err = ParseRoll(str)
	if len(res.Result) != 3 || (res.Sum < 3 || res.Sum > 12) {
		t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.Result), res.Sum)
	}

	str = "D20 +1"
	res, err = ParseRoll(str)
	if err != nil {
		t.Errorf("Error parsing %s: %v", str, err)
	}
	if res.Bonus != 1 {
		t.Errorf("Error applying bonus(%s) %v: %d", str, err, res.Bonus)
	}

	if len(res.Result) != 1 || (res.Sum < 2 || res.Sum > 21) {
		t.Errorf("%v: %d (%d)", err, len(res.Result), res.Sum)
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