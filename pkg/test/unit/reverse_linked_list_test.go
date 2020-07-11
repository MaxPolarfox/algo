package unit

import (
	"encoding/json"
	"fmt"
	"github.com/GO_algos/pkg/linked_lists"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type reverseLinkedListTestCase struct {
	Input    reverseLinkedTestData `json:"input"`
	Output reverseLinkedTestData `json:"output"`
}

type reverseLinkedTestData struct {
	Head string `json:"head"`
	Nodes []testNode `json:"nodes"`
}

type testNode struct {
	ID		string	`json:"id"`
	Next	*string `json:"next"`
	Value	int		`json:"value"`
}


func TestReverseLinkedList(t *testing.T) {
	Convey("TestReverseLinkedList", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/reverse_linked_list.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []reverseLinkedListTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				// Build input LL
				inputLL := BuildLinkedList(testCase.Input.Head, testCase.Input.Nodes)
				output := linked_lists.ReverseLinkedList(inputLL)
				// Build Expected output LL
				expectedOutPutLL := BuildLinkedList(testCase.Output.Head, testCase.Output.Nodes)
				c.So(output, ShouldResemble, expectedOutPutLL)
			})
		}

		BuildLinkedList(testCases[0].Input.Head, testCases[0].Input.Nodes)
	})
}

func BuildLinkedList (startID string, nodes []testNode) *linked_lists.LinkedList {
	var head *linked_lists.LinkedList
	var nextID *string

	for _, node := range nodes {
		if node.ID == startID {
			head = &linked_lists.LinkedList{Value: node.Value, Next: nil}
			nextID = node.Next
			break
		}
	}

	current := head
	for nextID != nil {
		for _, node := range nodes {
			if node.ID == *nextID {
				current.Next = &linked_lists.LinkedList{Value: node.Value, Next: nil}
				current = current.Next
				nextID = node.Next
				continue
			}
		}
	}

	return head
}