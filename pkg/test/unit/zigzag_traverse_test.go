package unit

import (
	"encoding/json"
	"fmt"
	"github.com/GO_algos/pkg/arrays"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type zigzagTraverseTestCase struct {
	Input    [][]int	`json:"input"`
	Expected []int 		`json:"expected"`
}

func TestZigzagTraverse(t *testing.T) {
	Convey("Test ZigzagTraverse", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/zigzag_traverse.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []zigzagTraverseTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := arrays.ZigzagTraverse(testCase.Input)
				c.So(output, ShouldResemble, testCase.Expected)
			})
		}
	})
}


