package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		for level := 1; level <= BTreeLevelCount(root); level++ {
			BTreeApplyAtLevel(root, level, f)
		}
	}
}

func BTreeApplyAtLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root != nil {
		if level == 1 {
			f(root.Data)
			return
		}

		if level > 1 {
			BTreeApplyAtLevel(root.Left, level-1, f)
			BTreeApplyAtLevel(root.Right, level-1, f)
		}
	}
}
