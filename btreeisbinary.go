package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isBinaryUtil(root, nil, nil)
}

func isBinaryUtil(node *TreeNode, min *string, max *string) bool {
	if node == nil {
		return true
	}

	if (min != nil && node.Data <= *min) || (max != nil && node.Data >= *max) {
		return false
	}

	return isBinaryUtil(node.Left, min, &node.Data) && isBinaryUtil(node.Right, &node.Data, max)
}
