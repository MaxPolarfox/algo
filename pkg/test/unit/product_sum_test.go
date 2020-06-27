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

type productSumTestCase struct {
	Input []interface{} `json:"input"` // Since input type []interface{} the unmarshalled int type is float64
	Expected int `json:"expected"`
}

func TestProductSum(t *testing.T) {
	Convey("TestProductSum", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/product_sum.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []productSumTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := recursion.ProductSum(testCase.Input)
				c.So(output, ShouldEqual, testCase.Expected)
			})
		}
	})
}
