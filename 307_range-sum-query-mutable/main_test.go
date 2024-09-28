package rangesumquerymutable

import (
	"testing"
)

func TestNumArray(t *testing.T) {
	nums := []int{1, 3, 5}
	obj := Constructor(nums)

	if got := obj.SumRange(0, 2); got != 9 {
		t.Errorf("SumRange() = %v, want %v", got, 9)
	}

	obj.Update(1, 2)

	if got := obj.SumRange(0, 2); got != 8 {
		t.Errorf("SumRange() = %v, want %v", got, 8)
	}
}
