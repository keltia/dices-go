package dice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRoll_1(t *testing.T) {

	// valid
	str := "D20"
	res, err := ParseRoll(str)
	if err != nil {
		t.Errorf("Error parsing %s: %v", str, err)
	}
	assert.EqualValues(t, 0, res.Bonus, "no bonus")
	assert.EqualValues(t, 2, len(res.List), "should be 2")
	if len(res.List) != 2 || (res.Sum < 1 || res.Sum > 20) {
		t.Errorf("Bad string(%s): %v: %d (%v)", str, err, len(res.List), res)
	}
}

func TestParseRoll_2(t *testing.T) {
	str := "d20"
	res, err := ParseRoll(str)
	assert.NoError(t, err, "no error")
	assert.EqualValues(t, 2, len(res.List))
	if res.Sum < 1 || res.Sum > 20 {
		t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.List), res.Sum)
	}
}

func TestParseRoll_3(t *testing.T) {
	str := "3d4"
	res, err := ParseRoll(str)
	assert.NoError(t, err, "no error")
	assert.EqualValues(t, 4, len(res.List))
	if res.Sum < 3 || res.Sum > 12 {
		t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.List), res.Sum)
	}
}

func TestParseRoll_4(t *testing.T) {

	str := "D20 +1"
	res, err := ParseRoll(str)
	assert.EqualValues(t, 1, res.Bonus, "should be one")
	assert.EqualValues(t, 2, len(res.List), "should be 2")
	assert.NoError(t, err, "no error")

	if len(res.List) != 2 || (res.Sum < 2 || res.Sum > 21) {
		t.Errorf("%v: %d (%d)", err, len(res.List), res.Sum)
	}
}

func TestParseRoll_5(t *testing.T) {
	// valid
	str := "2d10"
	res, err := ParseRoll(str)
	assert.NoError(t, err, "no error")
	assert.EqualValues(t, 10, res.Size)
	if res.Sum < 2 {
		t.Errorf("Should be >=2: %d", res)
	}
	assert.EqualValues(t, 3, len(res.List))
	assert.EqualValues(t, 0, res.Bonus, "no bonus")
}

func TestParseRoll_6(t *testing.T) {
	str := "d100"
	res, err := ParseRoll(str)
	assert.NoError(t, err, "no error")
	assert.EqualValues(t, 2, len(res.List))
	if res.Sum < 1 || res.Sum > 100 {
		t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.List), res.Sum)
	}
}

func TestParseRoll_7(t *testing.T) {
	// invalid
	str := "D2"
	_, err := ParseRoll(str)
	assert.Error(t, err, "error")

	// invalid
	str = "D101"
	_, err = ParseRoll(str)
	assert.Error(t, err, "should be in error")
}
