package longestturbulentsubarray

func maxTurbulenceSize(arr []int) int {
	if len(arr) <= 1 {
		return len(arr)
	}

	curr := 1
	max := 1
	prev := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] && prev != 1 {
			curr += 1
		} else if arr[i] < arr[i-1] && prev != -1 {
			curr += 1
		} else if arr[i] != arr[i-1] {
			curr = 2
		} else {
			curr = 1
		}

		if arr[i] > arr[i-1] {
			prev = 1
		} else if arr[i] < arr[i-1] {
			prev = -1
		} else {
			prev = 0
		}

		if curr > max {
			max = curr
		}

	}

	return max
}
