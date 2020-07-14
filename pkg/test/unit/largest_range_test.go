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

type largestRangeTestCase struct {
	Input    []int `json:"input"`
	Expected []int `json:"expected"`
}

func TestLargestRange(t *testing.T) {
	Convey("Test LargestRange", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/largest_range.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []largestRangeTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := arrays.LargestRange(testCase.Input)
				c.So(output, ShouldResemble, testCase.Expected)
			})
		}
	})
}
