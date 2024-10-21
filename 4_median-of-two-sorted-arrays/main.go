package medianoftwosortedarrays

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	n := len(nums1) + len(nums2)

	half := (n + 1) / 2

	l, r := 0, len(nums1)-1

	for {
		i := (l + r) >> 1
		j := half - i - 2

		nums1Left, nums1Right := math.MinInt, math.MaxInt
		nums2Left, nums2Right := math.MinInt, math.MaxInt

		if i >= 0 {
			nums1Left = nums1[i]
		}
		if j >= 0 {
			nums2Left = nums2[j]
		}
		if i+1 < len(nums1) {
			nums1Right = nums1[i+1]
		}
		if j+1 < len(nums2) {
			nums2Right = nums2[j+1]
		}

		if nums1Left <= nums2Right && nums2Left <= nums1Right {
			if n%2 == 0 {
				return float64(max(nums1Left, nums2Left)+min(nums1Right, nums2Right)) / 2
			}
			return float64(max(nums1Left, nums2Left))
		} else if nums1Left > nums2Right {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}
