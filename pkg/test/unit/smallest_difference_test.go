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

func TestSmallestDifference(t *testing.T) {
	Convey("Test SmallestDifference", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/smallest_difference.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []ArraysTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := arrays.SmallestDifference(testCase.ArrayOne, testCase.ArrayTwo)
				c.So(output, ShouldResemble, testCase.ExpectedArray)
			})
		}
	})
}
