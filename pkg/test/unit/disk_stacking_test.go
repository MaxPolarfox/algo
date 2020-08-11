package unit

import (
	"encoding/json"
	"fmt"
	"github.com/GO_algos/pkg/dinamic_programing"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type diskStackingTestCase struct {
	Disks    [][]int   `json:"disks"`
	Expected [][]int `json:"expected"`
}

func TestDiskStacking(t *testing.T) {
	Convey("Test DiskStacking", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/disk_stacking.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []diskStackingTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := dinamic_programing.DiskStacking(testCase.Disks)
				c.So(output, ShouldResemble, testCase.Expected)
			})
		}
	})
}
