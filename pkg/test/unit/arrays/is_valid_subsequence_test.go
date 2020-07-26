package arrays

import (
	"encoding/json"
	"fmt"
	"github.com/GO_algos/pkg/arrays"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type isValidSubsequenceTestCase struct {
	Array    []int `json:"array"`
	Sequence []int `json:"sequence"`
	Expected bool `json:"expected"`
}

func TestIsValidSubsequence(t *testing.T) {
	Convey("Test IsValidSubsequence", t, func(c C) {
		absPath, err := filepath.Abs("../../test_data/is_valid_subsequence.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []isValidSubsequenceTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := arrays.IsValidSubsequence(testCase.Array, testCase.Sequence)
				c.So(output, ShouldResemble, testCase.Expected)
			})
		}
	})
}