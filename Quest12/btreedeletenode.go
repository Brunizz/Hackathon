package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == node {
		if root.Left != nil {
			min := BTReeMin(root.Left)
			root.Data = min.Data
			node = min
		} else if root.Right != nil {
			max := BTreeMax(root.Right)
			root.Data = max.Data
			node = max
		} else {
			return nil
		}
	}
	root.Left = BTreeDeleteNode(root.Left, node)
	root.Right = BTreeDeleteNode(root.Right, node)
	return root
}
