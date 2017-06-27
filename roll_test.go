package dice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseRoll(t *testing.T) {

	// valid
	str := "D20"
	t.Logf("==> %s", str)
	res, err := ParseRoll(str)
	assert.NoError(t, err, "no error")

	if len(res.List) != 2 || (res.Sum < 1 || res.Sum > 20) {
		t.Errorf("Bad string(%s): %v: %d (%v)", str, err, len(res.List), res)
	}

	str = "d20"
	t.Logf("==> %s", str)
	res, err = ParseRoll(str)
	assert.NoError(t, err, "no error")

	if len(res.List) != 2 || (res.Sum < 1 || res.Sum > 20) {
		t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.List), res.Sum)
	}

	str = "d10"
	t.Logf("==> %s", str)
	res, err = ParseRoll(str)
	assert.NoError(t, err, "no error")

	if len(res.List) != 2 || (res.Sum < 1 || res.Sum > 10) {
		t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.List), res.Sum)
	}

	str = "3d4"
	t.Logf("==> %s", str)
	res, err = ParseRoll(str)
	assert.NoError(t, err, "no error")

	if len(res.List) != 4 || (res.Sum < 3 || res.Sum > 12) {
		t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.List), res.Sum)
	}

	str = "D20 +1"
	t.Logf("==> %s", str)
	res, err = ParseRoll(str)
	assert.NoError(t, err, "no error")

	if len(res.List) != 2 || (res.Sum < 2 || res.Sum > 21) {
		t.Errorf("%v: %d (%d)", err, len(res.List), res.Sum)
	}

	str = "d100"
	t.Logf("==> %s", str)
	res, err = ParseRoll(str)
	assert.NoError(t, err, "no error")

	if len(res.List) != 2 || (res.Sum < 1 || res.Sum > 100) {
		t.Errorf("Bad string(%s): %v: %d (%d)", str, err, len(res.List), res.Sum)
	}

	// invalid
	str = "D2"
	t.Logf("==> %s", str)
	res, err = ParseRoll(str)
	assert.Error(t, err, "no error")

	str = "D101"
	t.Logf("==> %s", str)
	res, err = ParseRoll(str)
	assert.Error(t, err, "no error")
}
