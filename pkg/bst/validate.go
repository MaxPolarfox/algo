package bst

import "math"

func (tree *BST) Validate() bool {
	return tree.validateHelper(math.MinInt32, math.MaxInt32)
}

func (tree *BST) validateHelper(min, max int) bool {
	if tree.Value < min || tree.Value >= max {
		return false
	}

	if tree.Left != nil && !tree.Left.validateHelper(min, tree.Value) {
		return false
	}
	if tree.Right != nil && !tree.Right.validateHelper(tree.Value, max) {
		return false
	}
	return true
}
