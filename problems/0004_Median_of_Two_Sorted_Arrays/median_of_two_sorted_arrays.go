package problem0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j, prev, cur int
	l1 := len(nums1)
	l2 := len(nums2)
	lt := l1 + l2
	mid := lt/2 + 1
	for (i + j) != mid {
		prev = cur
		if i < l1 && j < l2 && nums1[i] <= nums2[j] {
			cur = nums1[i]
			i++
		} else if j < l2 {
			cur = nums2[j]
			j++
		} else if i < l1 {
			cur = nums1[i]
			i++
		}
	}
	if lt%2 == 1 {
		return float64(cur)
	}
	return float64(prev+cur) / 2
}
