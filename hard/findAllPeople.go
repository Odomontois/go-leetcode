package hard

import "sort"

type moment struct {
	all []int
	adj map[int][]int
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	meetings = append(meetings, []int{0, firstPerson, 0})

	knows := make([]bool, n)
	knows[0] = true
	momentSet := make(map[int]*moment)
	add := func(time, p1, p2 int) {
		mom := momentSet[time]
		if mom == nil {
			mom = new(moment)
			momentSet[time] = mom
		}
		if mom.adj == nil {
			mom.adj = make(map[int][]int)
		}
		adj := mom.adj[p1]
		if adj == nil {
			mom.all = append(mom.all, p1)
		}
		mom.adj[p1] = append(adj, p2)
	}
	for _, meet := range meetings {
		add(meet[2], meet[1], meet[0])
		add(meet[2], meet[0], meet[1])
	}
	var moments []int
	for m := range momentSet {
		moments = append(moments, m)
	}

	sort.Ints(moments)

	for _, m := range moments {
		var q []int
		mom := momentSet[m]
		for _, p := range mom.all {
			if knows[p] {
				q = append(q, p)
			}
		}
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, next := range mom.adj[cur] {
				if !knows[next] {
					knows[next] = true
					q = append(q, next)
				}
			}
		}
	}

	var res []int
	for p, know := range knows {
		if know {
			res = append(res, p)
		}
	}
	return res
}
