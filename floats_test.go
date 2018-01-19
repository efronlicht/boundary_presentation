package boundary

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumOfInts(t *testing.T) {
	assert.Equal(t, 0, sumOfInts(1, 2.2, 3, -4))
	assert.Equal(t, 4, sumOfInts(1, 3, math.Inf(-1), math.NaN()), "-inf was treated as an integer")
}

func Test_sumOfIntsFixed(t *testing.T) {
	got, _ := sumOfIntsFixed(1, 2.2, 3, -4)
	assert.Equal(t, 0, got)
	_, err := sumOfIntsFixed(1, 3, math.Inf(-1), math.NaN())
	assert.Error(t, err, "should error on Inf or NaN()")
}
