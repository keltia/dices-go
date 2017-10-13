package dice

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	f := isValid(1)
	assert.False(t, f)

	f = isValid(0)
	assert.False(t, f)

	f = isValid(23)
	assert.False(t, f)

	f = isValid(100)
	assert.True(t, f)
}

func TestCheckBonus(t *testing.T) {
	d := "3D4"
	bonus, str := checkBonus(d)
	assert.EqualValues(t, 0, bonus)
	assert.EqualValues(t, d, str)

	d = "D6 +1"
	bonus, str = checkBonus(d)
	assert.EqualValues(t, 1, bonus)
	assert.EqualValues(t, "D6", str)
}
