package lettercombinationsofaphonenumber

var m = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func backtrack(digits string, i int, set string, result *[]string) {
	if i == len(digits) {
		*result = append(*result, set)
		return
	}

	for j := 0; j < len(m[digits[i]]); j++ {
		backtrack(digits, i+1, set+m[digits[i]][j], result)
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	result := []string{}

	backtrack(digits, 0, "", &result)

	return result
}
