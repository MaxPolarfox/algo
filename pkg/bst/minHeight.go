package bst

func MinHeightBST(arr []int) *BST {
	return constructMinHeightBST(arr, 0, len(arr)-1)
}

func constructMinHeightBST(arr []int, start, end int) *BST {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	bst := &BST{Value: arr[mid]}
	bst.Left = constructMinHeightBST(arr, start, mid-1)
	bst.Right = constructMinHeightBST(arr, mid+1, end)
	return bst
}
