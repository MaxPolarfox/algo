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

func TestThreeNumberSum(t *testing.T) {
	Convey("Test ThreeNumberSum", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/three_number_sum.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []ArraysTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := arrays.ThreeNumberSum(testCase.Array, testCase.TargetSum)
				c.So(output, ShouldResemble, testCase.ExpectedNestedArrayInt)
			})
		}
	})
}
