package bst

func (tree *BST) FindClosestValue(target int) int {
	return tree.FindClosestValueHelper(target, tree.Value)
}

func (tree *BST) FindClosestValueHelper(target, closest int) int {
	curr := tree
	for curr != nil {
		if absDiff(target, closest) > absDiff(target, curr.Value) {
			closest = curr.Value
		}
		if target < curr.Value {
			curr = curr.Left
		} else if target > curr.Value {
			curr = curr.Right
		} else {
			break
		}

	}
	return closest
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
