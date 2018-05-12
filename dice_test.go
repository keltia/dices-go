package dice

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
