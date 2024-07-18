package y2024

import "leetcode/data"

type TreeNode = data.TreeNode

func smallestFromLeaf(root *TreeNode) string {
	var nodes, next []*TreeNode
	var traverse func(*TreeNode, *TreeNode)
	traverse = func(node, parent *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			nodes = append(nodes, node)
		}
		if node.Left != nil {
			traverse(node.Left, node)
		}
		if node.Right != nil {
			traverse(node.Right, node)
		}
		node.Left = parent
	}
	traverse(root, nil)
	var res []byte
	for {
		cur := 26
		for _, node := range nodes {
			if node == nil {
				return string(res)
			}
			if node.Val < cur {
				next = append(next[:0], node.Left)
				cur = node.Val
			} else if node.Val == cur {
				next = append(next, node.Left)
			}
		}
		nodes, next, res = next, nodes, append(res, byte(cur)+'a')
	}
}
