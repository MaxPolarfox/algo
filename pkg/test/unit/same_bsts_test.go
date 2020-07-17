package unit

import (
	"encoding/json"
	"fmt"
	"github.com/GO_algos/pkg/bst"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type sameBSTsTestCase struct {
	ArrayOne    []int   `json:"arrayOne"`
	ArrayTwo    []int   `json:"arrayTwo"`
	Expected 	bool 	`json:"expected"`
}

func TestSameBSTs(t *testing.T) {
	Convey("Test SameBSTs", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/same_bsts.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []sameBSTsTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				fmt.Printf("%v", testCase)
				output := bst.SameBSTs(testCase.ArrayOne, testCase.ArrayTwo)
				c.So(output, ShouldResemble, testCase.Expected)
			})
		}
	})
}
