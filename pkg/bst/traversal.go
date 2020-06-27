package bst

func (tree *BST) InOrderTraverse(arr []int) []int {
	if tree.Left != nil {
		arr = tree.Left.InOrderTraverse(arr)
	}
	arr = append(arr, tree.Value)
	if tree.Right != nil {
		arr = tree.Right.InOrderTraverse(arr)
	}
	return arr
}

func (tree *BST) PreOrderTraverse(arr []int) []int {
	arr = append(arr, tree.Value)
	if tree.Left != nil {
		arr = tree.Left.PreOrderTraverse(arr)
	}
	if tree.Right != nil {
		arr = tree.Right.PreOrderTraverse(arr)
	}
	return arr
}

func (tree *BST) PostOrderTraverse(arr []int) []int {
	if tree.Left != nil {
		arr = tree.Left.PostOrderTraverse(arr)
	}
	if tree.Right != nil {
		arr = tree.Right.PostOrderTraverse(arr)
	}
	arr = append(arr, tree.Value)

	return arr
}
