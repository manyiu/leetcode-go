package main

func groupThePeople(groupSizes []int) [][]int {
    m := make(map[int][][]int)

	for i, groupSize := range groupSizes {
		if groups, ok := m[groupSize]; ok {
			if len(groups[len(groups) - 1]) < groupSize {
				groups[len(groups) - 1] = append(groups[len(groups) - 1], i)
			} else {
				m[groupSize] = append(groups, []int{i})
			}
		} else {
			m[groupSize] = [][]int{{i}}
		}
	}

	r := [][]int{}

	for _, element := range m {
		r = append(r, element...)
	}

	return r
}

func main() {}
