package dice

import (
	"testing"
)

func TestParseRoll (t *testing.T) {

	// valid
	str := "D20"
	t.Logf("==> %s", str)
	res, err := ParseRoll(str)
	if len(res.Result) != 1 || (res.Sum < 1 || res.Sum > 20) {
		t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.Result), res.Sum)
	}

	str = "d20"
	t.Logf("==> %s", str)
	res, err = ParseRoll(str)
	if len(res.Result) != 1 || (res.Sum < 1 || res.Sum > 20) {
		t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.Result), res.Sum)
	}

	str = "3d4"
	t.Logf("==> %s", str)
	res, err = ParseRoll(str)
	if len(res.Result) != 3 || (res.Sum < 3 || res.Sum > 12) {
		t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.Result), res.Sum)
	}

	str = "D20 +1"
	t.Logf("==> %s", str)
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

	str = "d100"
	t.Logf("==> %s", str)
	res, err = ParseRoll(str)
	if err != nil {
		t.Errorf("Error, should be valid: %v %v", res, err)
	} else {
		if len(res.Result) != 1 || (res.Sum < 1 || res.Sum > 100) {
			t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.Result), res.Sum)
		}
	}

	// invalid
	str = "D2"
	t.Logf("==> %s", str)
	res, err = ParseRoll(str)
	if err == nil {
		t.Errorf("Bad size(%s): %v: %d (%d)", str, err, len(res.Result), res.Sum)
	}

	str = "D101"
	t.Logf("==> %s", str)
	res, err = ParseRoll(str)
	if err == nil {
		t.Errorf("Bad size(%s): %v", str, err)
	}
}

