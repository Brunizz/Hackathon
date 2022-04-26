package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {

	height := BTreeLevelCount(root)
	for i := 0; i < height; i++ {
		applylevel(root, i, f)
	}
}

func applylevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 0 {
		f(root.Data)
	} else if level > 0 {
		applylevel(root.Left, level-1, f)
		applylevel(root.Right, level-1, f)
	}
}
