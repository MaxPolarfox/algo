package unit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/GO_algos/pkg/recursion"
)

type getPermutationsTestCase struct {
	Input    []int `json:"input"`
	Expected [][]int `json:"expected"`
}

func TestGetPermutations(t *testing.T) {
	Convey("TestGetNthFib", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/get_permutations.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []getPermutationsTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := recursion.GetPermutations(testCase.Input)
				c.So(output, ShouldResemble, testCase.Expected)
			})
		}
	})
}