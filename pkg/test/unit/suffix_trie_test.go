package unit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/GO_algos/pkg/tries"
)

type testCaseSuffixTrie struct{
	String string `json:"string"`
	ClassMethodsToCall []methodToCallSuffixTreeTestCase `json:"classMethodsToCall"`
}

type methodToCallSuffixTreeTestCase struct {
	Arguments []string `json:"arguments"`
	Method string `json:"method"`
	Output bool `json:"output"`
}

func TestSuffixTrie(t *testing.T) {
	Convey("TestSuffixTrie", t, func(c C) {
		absPath, err := filepath.Abs("../test_data/suffix_trie.json")
		c.So(err, ShouldBeNil)

		testCasesDataBytes, err := ioutil.ReadFile(absPath)
		c.So(err, ShouldBeNil)

		testCases := []testCaseSuffixTrie{}

		err = json.Unmarshal(testCasesDataBytes, &testCases)
		c.So(err, ShouldBeNil)

		for i, testCase := range testCases {
			Convey(fmt.Sprintf("Test %v:", i+1), func(c C) {

				trie := tries.NewSuffixTrie()
				trie.PopulateSuffixTrieFrom(testCase.String)


				for _, method := range testCase.ClassMethodsToCall {
					if method.Method == "contains" {
						output := trie.Contains(method.Arguments[0])
						c.So(output, ShouldEqual, method.Output)
					}
				}
			})
		}
	})
}

