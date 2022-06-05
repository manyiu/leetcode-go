package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	e1input := []int{2, 7, 11, 15}
	e1expected := []int{0, 1}
	assert.Equal(t, e1expected, twoSum(e1input, 9))

	e2input := []int{3, 2, 4}
	e2expected := []int{1, 2}
	assert.Equal(t, e2expected, twoSum(e2input, 6))

	e3input := []int{3, 3}
	e3expected := []int{0, 1}
	assert.Equal(t, e3expected, twoSum(e3input, 6))
}
