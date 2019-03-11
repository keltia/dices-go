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

func TestOpenDice_Roll(t *testing.T) {
	d := openDice{d: constantDice(0)}
	r := Result{Sum: 2}
	r1 := d.Roll(r)
	assert.EqualValues(t, r, r1)
}
