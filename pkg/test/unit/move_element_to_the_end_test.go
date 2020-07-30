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

func TestMoveElementToTheEnd(t *testing.T) {
	Convey("Test MoveElementToTheEnd", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/move_element_to_the_end.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []ArraysTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				fmt.Printf("\ntest %v,  %v", testCase.Array, testCase.ToMove)
				output := arrays.MoveElementToEnd(testCase.Array, testCase.ToMove)
				c.So(output, ShouldResemble, testCase.ExpectedArray)
			})
		}
	})
}
