package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || elem == "" {
		return nil
	}

	current := root

	for current != nil {
		if current.Data == elem {
			return current
		} else if elem < current.Data {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return nil
}
