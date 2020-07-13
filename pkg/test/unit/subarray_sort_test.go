package unit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/GO_algos/pkg/sorting"
)

type subarraySumTestCase struct {
	Input	[]int	`json:"input"`
	Output	[]int   `json:"output"`
}

func TestSubarraySum(t *testing.T) {
	Convey("TestSubarraySum", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/subarray_sum.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []subarraySumTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := sorting.QuickSort(testCase.Input)
				c.So(output, ShouldEqual, testCase.Output)
			})
		}
	})
}