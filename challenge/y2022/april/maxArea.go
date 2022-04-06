package april

func maxArea(height []int) (best int) {
	i, j, h := 0, len(height)-1, 0
	for i < j {
		hi, hj := height[i], height[j]
		if hi < hj {
			h = hi
			i++
		} else {
			h = hj
			j--
		}
		a := h * (j - i + 1)
		if a > best {
			best = a
		}
	}
	return
}
