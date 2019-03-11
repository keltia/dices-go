package dice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResult_Append(t *testing.T) {
	r := Result{}
	s := r.Append(1)

	u := Result{List: []int{1}, Sum: 1}
	assert.NotNil(t, s)
	assert.EqualValues(t, u, s)
}

func TestResult_Merge(t *testing.T) {
	r1 := Result{List: []int{1}, Sum: 1}
	r2 := Result{List: []int{2}, Sum: 2}
	r3 := Result{List: []int{1, 2}, Sum: 3}
	s := r1.Merge(r2)

	assert.NotNil(t, s)
	assert.NotEmpty(t, s)
	assert.EqualValues(t, r3, s)
}

func TestResult_Natural(t *testing.T) {
	r := Result{List: []int{20}, Sum: 20, Size: 20}
	yes := r.Natural()
	assert.True(t, yes)

	r = Result{List: []int{13}, Sum: 13, Size: 20}
	no := r.Natural()
	assert.False(t, no)
}
