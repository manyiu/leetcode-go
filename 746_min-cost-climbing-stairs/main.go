package main

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	p := make([]int, l+1)

	for i := 2; i <= l; i++ {
		p[i] = min(p[i-1]+cost[i-1], p[i-2]+cost[i-2])
	}

	return p[l]
}

func main() {}
