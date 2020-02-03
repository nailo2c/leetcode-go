package problem0034

func searchRange(nums []int, target int) []int {
	mid := binarySearch(nums, target)
	if mid == -1 {
		return []int{-1, -1}
	}

	start := mid
	for {
		s := binarySearch(nums[:start], target)
		if s == -1 {
			break
		}
		start = s
	}

	end := mid
	for {
		e := binarySearch(nums[end+1:], target)
		if e == -1 {
			break
		}
		end += e + 1
	}

	return []int{start, end}
}

func binarySearch(nums []int, target int) int {
	lo, hi, mid := 0, len(nums)-1, 0
	for lo <= hi {
		mid = (lo + hi) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid - 1
		} else if nums[mid] < target {
			lo = mid + 1
		}
	}
	return -1
}
