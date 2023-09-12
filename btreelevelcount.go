package piscine

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := BTreeLevelCount(root.Left)
	rightHeight := BTreeLevelCount(root.Right)

	return 1 + max(leftHeight, rightHeight)
}
