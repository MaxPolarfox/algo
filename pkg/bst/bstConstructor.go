package bst

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}

func (tree *BST) Contains(value int) bool {
	if tree.Value > value {
		if tree.Left == nil {
			return false
		} else {
			return tree.Left.Contains(value)
		}
	} else if tree.Value < value {
		if tree.Right == nil {
			return false
		} else {
			return tree.Right.Contains(value)
		}
	}
	return true
}

func (tree *BST) Remove(value int) *BST {
	tree.removeHelper(value, nil)
	return tree
}

func (tree *BST) removeHelper(value int, parent *BST) {
	// Find the node to delete
	if value < tree.Value {
		if tree.Left != nil {
			tree.Left.removeHelper(value, tree)
		}
	} else if value > tree.Value {
		if tree.Right != nil {
			tree.Right.removeHelper(value, tree)
		}
	} else { // Node found
		if tree.Left != nil && tree.Right != nil { // Node has Left and Right
			tree.Value = tree.Right.getMinValue()
			tree.Right.removeHelper(tree.Value, tree)
		} else if parent == nil { // Root node
			if tree.Left != nil {
				tree.Value = tree.Left.Value
				tree.Right = tree.Left.Right
				tree.Left = tree.Left.Left
			} else if tree.Right != nil {
				if tree.Right != nil {
					tree.Value = tree.Right.Value
					tree.Left = tree.Right.Left
					tree.Right = tree.Right.Right
				}
			} else {} // single-node tree
		} else if parent.Left == tree { // Left node, Not root
			if tree.Left != nil {
				parent.Left = tree.Left
			} else {
				parent.Left = tree.Right
			}
		} else if parent.Right == tree { // Right node, Not root
			if tree.Right != nil {
				parent.Right = tree.Right
			} else {
				parent.Right = tree.Left
			}
		}
	}
}

func (tree *BST) getMinValue() int {
	if tree.Left == nil {
		return tree.Value
	}
	return tree.Left.getMinValue()
}