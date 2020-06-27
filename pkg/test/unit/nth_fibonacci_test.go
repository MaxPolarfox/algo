package unit

import (
"encoding/json"
"fmt"
"io/ioutil"
"path/filepath"
"testing"

"github.com/algos/pkg/recursion"
. "github.com/smartystreets/goconvey/convey"
)

type nthFibonacciTestCase struct {
	Input int `json:"input"`
	Expected int `json:"expected"`
}

func TestGetNthFib(t *testing.T) {
	Convey("TestGetNthFib", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/get_nth_fib.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []nthFibonacciTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := recursion.GetNthFib(testCase.Input)
				c.So(output, ShouldEqual, testCase.Expected)
			})
		}
	})
}