package graphs

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

func (tree *AncestralTree) addAsAncestor(descendants ...*AncestralTree) {
	for _, descendant := range descendants {
		descendant.Ancestor = tree
	}
}

func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	depthOne := getDescendantDepth(descendantOne, top)
	depthTwo := getDescendantDepth(descendantTwo, top)
	if depthOne > depthTwo {
		return helper(descendantOne, descendantTwo, depthOne-depthTwo)
	}
	return helper(descendantTwo, descendantOne, depthTwo-depthOne)
}

func getDescendantDepth(descendant, top *AncestralTree) int {
	depth := 0
	for descendant != top {
		depth++
		descendant = descendant.Ancestor
	}
	return depth
}

func helper(lower, highter *AncestralTree, diff int) *AncestralTree {
	for diff > 0 {
		lower = lower.Ancestor
		diff--
	}
	for lower != highter {
		lower = lower.Ancestor
		highter = highter.Ancestor
	}
	return lower
}
