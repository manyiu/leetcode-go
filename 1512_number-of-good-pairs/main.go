package main

func numIdenticalPairs(nums []int) int {
	c := 0
	l := [101]int{}

	for i := 0; i < len(nums); i++ {
		c += l[nums[i]]
		l[nums[i]]++
	}

	return c
}

func main() {

}
