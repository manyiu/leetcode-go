package rangesumqueryimmutable

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	total := 0

	for i := 0; i < len(nums); i++ {
		total += nums[i]
		nums[i] = total
	}

	return NumArray{
		prefixSum: nums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	preLeft := 0
	preRight := this.prefixSum[right]

	if left > 0 {
		preLeft = this.prefixSum[left-1]
	}

	return preRight - preLeft
}
