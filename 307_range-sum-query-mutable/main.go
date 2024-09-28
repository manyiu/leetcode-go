package rangesumquerymutable

type NumArray struct {
	l, r        int
	sum         int
	left, right *NumArray
}

func NewNumArray(nums []int, l, r int) *NumArray {
	if l == r {
		return &NumArray{
			l:   l,
			r:   r,
			sum: nums[l],
		}
	}

	m := (l + r) / 2

	left := NewNumArray(nums, l, m)
	right := NewNumArray(nums, m+1, r)
	sum := left.sum + right.sum

	return &NumArray{
		l,
		r,
		sum,
		left,
		right,
	}
}

func Constructor(nums []int) NumArray {
	return *NewNumArray(nums, 0, len(nums)-1)
}

func (this *NumArray) Update(index int, val int) {
	if this.l == this.r && this.l == index {
		this.sum = val
		return
	}

	m := (this.l + this.r) / 2

	if index <= m {
		this.left.Update(index, val)
	} else {
		this.right.Update(index, val)
	}

	this.sum = this.left.sum + this.right.sum
}

func (this *NumArray) SumRange(left int, right int) int {

	if this.l == left && this.r == right {
		return this.sum
	}

	m := (this.l + this.r) / 2

	if left > m {
		return this.right.SumRange(left, right)
	}

	if right <= m {
		return this.left.SumRange(left, right)
	}

	return this.left.SumRange(left, m) + this.right.SumRange(m+1, right)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
