package binary_trees

func NodeDepths(root *BinaryTree) int {
	return calculateNodeDepths(root, 0)
}

func calculateNodeDepths(node *BinaryTree, depth int) int {
	if node == nil {
		return 0
	}
	return depth + calculateNodeDepths(node.Left, depth+1) + calculateNodeDepths(node.Right, depth+1)
}
