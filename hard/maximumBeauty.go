package hard

import "sort"

func maximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {

	sort.Ints(flowers)

	sum, sums := int64(0), make([]int64, len(flowers))
	for i, x := range flowers {
		sum += int64(x)
		sums[i] = sum
	}

	if flowers[0] >= target {
		return int64(len(flowers)) * int64(full)
	}

	p := len(flowers)
	best := int64(0)

	getLevel := func() (level int64, ok bool) {
		if p >= len(flowers) || p < 0 {
			return
		}
		flood := int64(flowers[p])*int64(p+1) - sums[p]
		if flood <= newFlowers {
			return (newFlowers-flood)/int64(p+1) + int64(flowers[p]), true
		}
		return
	}

	for i := len(flowers); i >= 0; i-- {
		if i < len(flowers) {
			fill := target - flowers[i]
			if fill < 0 {
				fill = 0
			}
			newFlowers -= int64(fill)
			if newFlowers < 0 {
				break
			}
		}

		level, ok := getLevel()
		for ; p >= 0 && (p >= i || !ok); level, ok = getLevel() {
			p--
		}
		if i == 0 {
			level = 0
		} else if level >= int64(target) {
			level = int64(target - 1)
		}
		points := level*int64(partial) + int64(len(flowers)-i)*int64(full)
		if points > best {
			best = points
		}
	}

	return best
}
