package lcadepest

import "leetcode/data"

type TreeNode = data.TreeNode

func traverse(node *TreeNode, depth int) (int, *TreeNode) {
	if node == nil {
		return 0, nil
	}
	l, ln := traverse(node.Left, depth+1)
	r, rn := traverse(node.Right, depth+1)
	if ln == nil && rn == nil {
		return depth, node
	}
	if l == r {
		return l, node
	} else if l > r {
		return l, ln
	}
	return r, rn
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, lca := traverse(root, 1)
	return lca
}
