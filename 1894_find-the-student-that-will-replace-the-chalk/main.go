package findthestudentthatwillreplacethechalk

func chalkReplacer(chalk []int, k int) int {

	totalPreCycle := 0

	for _, use := range chalk {
		totalPreCycle += use
	}

	k %= totalPreCycle

	i := 0

	for {
		k -= chalk[i]

		if k < 0 {
			return i
		}

		i++
	}
}
