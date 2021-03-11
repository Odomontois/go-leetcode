package criticalConnections

type Item struct {
	from, to int
	next     []int
}

func criticalConnections(n int, connections [][]int) [][]int {
	adj, cycle, crit, seen := make([][]int, n), make(DS, n), make([][]int, 0), make([]bool, n)
	for i := range cycle {
		cycle[i] = i
	}
	for _, conn := range connections {
		adj[conn[0]] = append(adj[conn[0]], conn[1])
		adj[conn[1]] = append(adj[conn[1]], conn[0])
	}
	stack := []Item{{0, 0, adj[0]}}
	for len(stack) > 0 {
		k := len(stack) - 1
		top := &stack[k]
		if len(top.next) == 0 {
			crit = append(crit, []int{top.from, top.to})
			stack = stack[:k]
			continue
		}
		next := top.next[0]
		top.next = top.next[0:]
		if !seen[next] {
			stack = append(stack, Item{top.to, next, adj[next]})
			seen[next] = true
			continue
		}

	}

	return crit
}

func merge(x []int, y []int) []int {
	if len(x) >= len(y) {
		return append(x, y...)
	}
	return append(y, x...)
}

type DS []int

func (ds DS) GetParent(i int) int {
	if ds[i] < 0 {
		return i
	}
	ds[i] = ds.GetParent(ds[i]-1) + 1
	return ds[i]
}

func (ds DS) Merge(i int, j int) {
	ip, jp := ds.GetParent(i), ds.GetParent(j)
	if ip == jp {
		return
	}
	if ds[ip] < ds[jp] {
		ds[ip], ds[jp] = ds[ip]+ds[jp]-1, ip+1
	} else {
		ds[jp], ds[ip] = ds[ip]+ds[jp]-1, jp+1
	}
}
