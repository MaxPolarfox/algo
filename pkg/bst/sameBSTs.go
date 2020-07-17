package bst


// An array of integers is said to represent the Binary Search Tree (BST) obtained by inserting each integer in the array,
// from left to right, into the BST. Write a function that takes in two arrays of integers and determines whether these arrays represent the same BST.
// Note that you're not allowed to construct any BSTs in your code. A BST is a Binary Tree that consists only of BST nodes.
// A node is said to be a valid BST node if and only if it satisfies the BST property:
// 		1) its value is strictly greater than the values of every node to its left;
// 		2) its value is less than or equal to the values of every node to its right;
//		3) its children nodes are either valid BST nodes themselves or None

//	Sample Input: arrayOne = [10, 15, 8, 12, 94, 81, 5, 2, 11]  arrayTwo = [10, 8, 5, 15, 2, 12, 11, 94, 81]
//  Sample Output: true

// Time O(n^2) Space O(n^2)

func SameBSTs(arrOne, arrTwo []int) bool {
	if  len(arrOne) != len(arrTwo) {
		return false
	}

	if len(arrOne) == 0 && len(arrTwo) == 0 {
		return true
	}

	if arrOne[0] != arrTwo[0] {
		return false
	}

	leftOne := getSmaller(arrOne)
	leftTwo := getSmaller(arrTwo)
	rightOne := getBiggerOrEqual(arrOne)
	rightTwo := getBiggerOrEqual(arrTwo)

	return SameBSTs(leftOne, leftTwo) && SameBSTs(rightOne, rightTwo)
}

func getSmaller(arr []int) []int {
	smaller := []int{}
	for i :=1; i < len(arr); i++ {
		if arr[i] < arr[0] {
			smaller = append(smaller, arr[i])
		}
	}
	return smaller
}

func getBiggerOrEqual(arr []int) []int {
	biggetOrEqual := []int{}
	for i :=1; i < len(arr); i++ {
		if arr[i] >= arr[0] {
			biggetOrEqual = append(biggetOrEqual, arr[i])
		}
	}
	return biggetOrEqual
}
