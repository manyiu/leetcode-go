package validpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	e1 := "A man, a plan, a canal: Panama"
	assert.True(t, isPalindrome(e1))

	e2 := "race a car"
	assert.False(t, isPalindrome(e2))

	e3 := " "
	assert.True(t, isPalindrome(e3))

	e4 := "0P"
	assert.False(t, isPalindrome(e4))
}
