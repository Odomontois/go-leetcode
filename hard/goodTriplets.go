package hard

func goodTriplets(nums1 []int, nums2 []int) (res int64) {
	n := len(nums1)
	place, counts, pairs := make([]int, n), make(fenwick, 4*n), make(fenwick, 4*n)
	for i, x := range nums1 {
		place[x] = i
	}
	for _, x := range nums2 {
		px := place[x]
		cx := counts.calc(px)
		counts.update(px, 1)
		res += pairs.calc(px)
		pairs.update(px, cx)
	}
	return
}

type fenwick []int64

func (f fenwick) calc(to int) (res int64) {
	start, end, pos := 0, len(f)/4, 0
	for end-start > 1 {
		m := (start + end) / 2
		if to < m {
			end, pos = m, 2*pos+1
		} else {
			res += f[2*pos+1]
			start, pos = m, 2*pos+2
		}
	}

	return res + f[pos]
}

func (f fenwick) update(idx int, d int64) {
	start, end, pos := 0, len(f)/4, 0
	f[0] += d
	for end-start > 1 {
		m := (start + end) / 2
		if idx < m {
			end, pos = m, 2*pos+1
		} else {
			start, pos = m, 2*pos+2
		}
		f[pos] += d
	}
}
