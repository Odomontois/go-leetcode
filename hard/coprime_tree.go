package hard

import "fmt"

func getCoprimes(nums []int, edges [][]int) []int {
	res := make([]int, len(nums))

	res[0] = -1

	graph := makeGraph(len(nums), edges)

	down := func(x int) interface{} {
		return nil
	}

	up := func(x int) {}

	graph.Dfs(0, down, up)
	return res
}

func gcd(x int, y int) int {
	if x < y {
		return gcd(y, x)
	}
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

type Graph struct {
	size  int
	adj   [][]int
	stack *[]StackItem
}

type StackItem struct {
	next  []int
	idx   int
	state interface{}
}

func makeGraph(size int, edges [][]int) Graph {
	adj := make([][]int, size)
	for _, v := range edges {
		adj[v[0]] = append(adj[v[0]], v[1])
		adj[v[1]] = append(adj[v[1]], v[0])
	}
	return Graph{size, adj, nil}
}

func (g Graph) Current() interface{} {
	if g.stack == nil || len(*g.stack) == 0 {
		return nil
	}
	return (*g.stack)[len(*g.stack)-1].state
}

func (g Graph) Dfs(start int, down func(int) interface{}, up func(int)) {
	init := down(start)
	stack := []StackItem{{g.adj[start], start, init}}
	seen := make([]bool, g.size)
	seen[start] = true
	g.stack = &stack

	for len(stack) > 0 {
		n := len(stack) - 1
		k := len(stack[n].next) - 1
		if k < 0 {
			up(stack[n].idx)
			stack = stack[0:n]
			continue
		}
		next := &stack[n].next
		i := (*next)[k]
		*next = (*next)[0:k]
		if seen[i] {
			continue
		}
		seen[i] = true
		state := down(i)

		stack = append(stack, StackItem{g.adj[i], i, state})
	}
	g.stack = nil
}

func main0() {
	fmt.Println(
		getCoprimes(
			[]int{2, 3, 3, 2},
			[][]int{{0, 1}, {1, 2}, {1, 3}},
		))
	fmt.Println(
		getCoprimes(
			[]int{5, 6, 10, 2, 3, 6, 15},
			[][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
		))

	fmt.Println(
		getCoprimes(
			[]int{5, 6, 10, 2, 3, 7, 15},
			[][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
		))
}

func capCheck() {
	xs := make([]int, 0)
	fmt.Println(len(xs), cap(xs))
	xs = append(xs, 1)
	fmt.Println(len(xs), cap(xs))
	xs = append(xs, 1)
	fmt.Println(len(xs), cap(xs))
	xs = append(xs, 1)
	fmt.Println(len(xs), cap(xs))
	xs = append(xs, 1)
	fmt.Println(len(xs), cap(xs))
	xs = append(xs, 1)
	fmt.Println(len(xs), cap(xs))
	xs = append(xs, 1)
	fmt.Println(len(xs), cap(xs))
	xs = append(xs, 1)
	fmt.Println(len(xs), cap(xs))

	xs = xs[0 : len(xs)-1]
	fmt.Println(len(xs), cap(xs))
	xs = xs[0 : len(xs)-1]
	fmt.Println(len(xs), cap(xs))
	xs = xs[0 : len(xs)-1]
	fmt.Println(len(xs), cap(xs))
	xs = xs[0 : len(xs)-1]
	fmt.Println(len(xs), cap(xs))
	xs = xs[0 : len(xs)-1]
	fmt.Println(len(xs), cap(xs))
	xs = xs[0 : len(xs)-1]
	fmt.Println(len(xs), cap(xs))
}
