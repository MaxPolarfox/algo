package binary_trees

func BranchSums(root *BinaryTree) []int {
	sums := []int{}
	calculateBranchSums(root, 0, &sums)
	return sums
}

func calculateBranchSums(node *BinaryTree, runningSum int, sums *[]int) {
	if node == nil {
		return
	}
	runningSum += node.Value
	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, runningSum)
		return
	}
	calculateBranchSums(node.Left, runningSum, sums)
	calculateBranchSums(node.Right, runningSum, sums)
}
