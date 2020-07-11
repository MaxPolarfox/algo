package unit

import (
	"encoding/json"
	"fmt"
	"github.com/algos/GO_algos/pkg/dinamic_programing"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type waterAreaTestCase struct {
	Heights   []int 	`json:"heights"`
	Output 		int     `json:"output"`
}

func TestWaterArea(t *testing.T) {
	Convey("TestWaterArea", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/water_area.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []waterAreaTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := dinamic_programing.WaterArea(testCase.Heights)
				c.So(output, ShouldEqual, testCase.Output)
			})
		}
	})
}