package minimumcostfortickets

import (
	"math"
)

func mincostTickets(days []int, costs []int) int {
	dp := make([]int, len(days))
	passDays := [3]int{1, 7, 30}

	for i := 0; i < len(dp); i++ {
		minCost := math.MaxInt

		for j := 0; j < len(costs); j++ {
			for k := i; k >= 0 && days[i]-days[k] < passDays[j]; k-- {
				cost := costs[j]

				if k > 0 {
					cost += dp[k-1]
				}

				if cost < minCost {
					minCost = cost
				}
			}
		}

		dp[i] = minCost
	}

	return dp[len(days)-1]
}
