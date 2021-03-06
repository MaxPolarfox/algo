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

func TestFourNumbersSum(t *testing.T) {
	Convey("Test FourNumbersSum", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/four_numbers_sum.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []ArraysTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := arrays.FourNumberSum(testCase.Array, testCase.TargetSum)
				c.So(output, ShouldResemble, testCase.ExpectedNestedArrayInt)
			})
		}
	})
}
