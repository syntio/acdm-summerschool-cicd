package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	x := 2
	y := 2
	c := Sum(x, y)

	assert.Equal(t, 4, c)
}

func TestSub(t *testing.T) {
	x := 2
	y := 2
	c := Sub(x, y)
	assert.Equal(t, 0, c)
}

func TestMul(t *testing.T) {
	x := 2
	y := 3
	c := Mul(x, y)
	assert.Equal(t, 6, c)
}
