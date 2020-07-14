package unit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/GO_algos/pkg/arrays"
)

type minRewardsTestCase struct {
	Input    []int `json:"input"`
	Expected int   `json:"expected"`
}

func TestMinRewards(t *testing.T) {
	Convey("Test MinRewards", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/min_rewards.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []minRewardsTestCase{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := arrays.MinRewards(testCase.Input)
				c.So(output, ShouldResemble, testCase.Expected)
			})
		}
	})
}
