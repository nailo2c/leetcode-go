package problem0033

func search(nums []int, target int) int {
	size := len(nums)
	l, r := 0, size-1
	rotate := indexOfMin(nums)

	for l <= r {
		mid := (l + r) / 2
		rotateMid := (mid + rotate) % size
		switch {
		case target < nums[rotateMid]:
			r = mid - 1
		case nums[rotateMid] < target:
			l = mid + 1
		default:
			return rotateMid
		}
	}
	return -1
}

func indexOfMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[r] < nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
