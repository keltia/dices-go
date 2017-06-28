package dice

import "testing"

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
