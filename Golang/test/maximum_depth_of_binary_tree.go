package test

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaximumDepthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return  0
	}

	depth := 0
	leftDepth := MaximumDepthOfBinaryTree(root.Left)
	rightDepth := MaximumDepthOfBinaryTree(root.Right)

	if leftDepth > rightDepth {
		depth = leftDepth + 1
	}

	depth = rightDepth + 1

	return depth
}
