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

type knapsackProblem struct {
	Capacity int     `json:"capacity"`
	Items    [][]int `json:"items"`
	Expected struct {
		Items      []int `json:"items"`
		TotalValue int   `json:"totalValue"`
	} `json:"expected"`
}


func TestKnapsackProblem(t *testing.T) {
	Convey("Test KnapsackProblem", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/knapsack_problem.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []knapsackProblem{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {
				output := dinamic_programing.KnapsackProblem(testCase.Items, testCase.Capacity)
				c.So(output, ShouldResemble, []interface{}{testCase.Expected.TotalValue, testCase.Expected.Items})
			})
		}
	})
}
