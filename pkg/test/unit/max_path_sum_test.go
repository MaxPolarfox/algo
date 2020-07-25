package unit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/GO_algos/pkg/binary_trees"
)

type maxPathSumTestCase struct {
	Tree     binaryTreeTestData `json:"tree"`
	Expected int                `json:"expected"`
}

type binaryTreeTestData struct {
	Root  string               `json:"root"`
	Nodes []testBinaryTreeNode `json:"nodes"`
}

type testBinaryTreeNode struct {
	ID    string  `json:"id"`
	Left  *string `json:"left"`
	Right *string `json:"right"`
	Value int     `json:"value"`
}

func TestMaxPathSum(t *testing.T) {
	Convey("Test MaxPathSum", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/max_path_sum.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []maxPathSumTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {

				// Build input BinaryTree
				inputBT := BuildBinaryTree(testCase.Tree.Root, testCase.Tree.Nodes)

				output := binary_trees.MaxPathSum(inputBT)

				c.So(output, ShouldResemble, testCase.Expected)
			})
		}
	})
}

func BuildBinaryTree(rootID string, nodes []testBinaryTreeNode) *binary_trees.BinaryTree {
	var root *binary_trees.BinaryTree

	// Build rootNode
	for _, node := range nodes {
		if node.ID == rootID {
			root = &binary_trees.BinaryTree{Value: node.Value, Left: nil, Right: nil}
			for _, childNode := range nodes {
				if node.Left != nil && childNode.ID == *node.Left {
					root.Left = &binary_trees.BinaryTree{Value: childNode.Value, Left: nil, Right: nil}
				}
				if node.Right != nil && childNode.ID == *node.Right {
					root.Right = &binary_trees.BinaryTree{Value: childNode.Value, Left: nil, Right: nil}
				}
			}
			break
		}
	}

	// Insert Child Nodes
	for i := 1; i < len(nodes); i++ {
		node := findNodeInBinaryTree(root, nodes[i].Value)
		InsertNodeInBinaryTree(node, nodes[i].Left, nodes[i].Right, nodes)
	}

	return root
}

func InsertNodeInBinaryTree(node *binary_trees.BinaryTree, left, right *string, nodes []testBinaryTreeNode) {
	for _, childNode := range nodes {
		if left != nil && childNode.ID == *left {
			node.Left = &binary_trees.BinaryTree{Value: childNode.Value, Left: nil, Right: nil}
		}
		if right != nil && childNode.ID == *right {
			node.Right = &binary_trees.BinaryTree{Value: childNode.Value, Left: nil, Right: nil}
		}
	}
}

func findNodeInBinaryTree(root *binary_trees.BinaryTree, value int) *binary_trees.BinaryTree {
	if root == nil {
		return nil
	}

	if root.Value == value {
		return root
	}

	leftSideResult := findNodeInBinaryTree(root.Left, value)
	if leftSideResult != nil {
		return leftSideResult
	}

	rightSideResult := findNodeInBinaryTree(root.Right, value)
	return rightSideResult
}
