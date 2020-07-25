package binary_trees

import "math"

// Max Path Sum In Binary Tree
// Write a function that takes in a Binary Tree and returns its max path sum. A path is a collection of connected nodes
// in a tree where no node is connected to more than two other nodes; a path sum is the sum of the values of the nodes
// in a particular path. Each Binary Tree node has an integer value left child node, and a right child node.
// Children nodes can either be Binary Tree nodes themselves or None / null
// Sample Input: tree = 1
//                    /   \
//				    2       3
//			      /   \   /   \
//				 4     5 6     7
// Sample Output: 18 (5 + 2 + 1 + 3 + 7)

// Time: O(n) Space: O(log(n))
func MaxPathSum(tree *BinaryTree) int {
	_, maxSum := findMaxSum(tree)
	return maxSum
}

func findMaxSum(tree *BinaryTree) (int, int) {
	if tree == nil {
		return 0, math.MinInt32
	}
	leftMaxSumAsBranch, leftMaxPathSum := findMaxSum(tree.Left)
	rightMaxSumAsBranch, rightMaxPathSum := findMaxSum(tree.Right)
	maxChildSumAsBranch := max(leftMaxSumAsBranch, rightMaxSumAsBranch)

	value := tree.Value
	maxSumAsBranch := max(maxChildSumAsBranch+value, value)
	maxSumAsRootNode := max(leftMaxSumAsBranch+value+rightMaxSumAsBranch, maxSumAsBranch)
	maxPathSum := max(leftMaxPathSum, rightMaxPathSum, maxSumAsRootNode)
	return maxSumAsBranch, maxPathSum
}

func max(first int, vals ...int) int {
	for _, val := range vals {
		if val > first {
			first = val
		}
	}
	return first
}
