package problem0015

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		if nums[i] > 0 {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s < 0 {
				l++
				continue
			}
			if s > 0 {
				r--
				continue
			}
			res = append(res, []int{nums[i], nums[l], nums[r]})
			for nums[l] == nums[l+1] && l+1 < r {
				l++
			}
			l++
			r--
		}
	}

	return res
}
