package unit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/GO_algos/pkg/arrays"
)

type subarraySortTestCase struct {
	Input	[]int	`json:"input"`
	Output	[]int   `json:"output"`
}

func TestSubarraySort(t *testing.T) {
	Convey("TestSubarraySort", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/subarray_sort.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []subarraySortTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := arrays.SubarraySort(testCase.Input)
				c.So(output, ShouldResemble, testCase.Output)
			})
		}
	})
}