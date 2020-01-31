package problem0011

func maxArea(height []int) int {
	max, width, area := 0, 0, 0
	i, j := 0, len(height)-1
	for i < j {
		width = j - i

		if height[i] < height[j] {
			area = height[i] * width
			i++
		} else {
			area = height[j] * width
			j--
		}

		if area > max {
			max = area
		}
	}

	return max
}
