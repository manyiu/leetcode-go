package main

func recursive(team int, sum int) int {
	if team == 2 {
		return sum + 1
	} else if team <= 1 {
		return sum
	}

	if team % 2 == 0 {
		return recursive(team/2, sum + team/2)
	}

	return recursive(team/2 + 1, sum + team/2)
}

func numberOfMatches(n int) int {
    return recursive(n, 0)
}
