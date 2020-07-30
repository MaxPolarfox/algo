package unit

type ArraysTestCase struct {
	Array                  []int   `json:"array"`
	NestedArrayInt         [][]int `json:"nested_array_int"`
	Sequence               []int   `json:"sequence"`
	TargetSum              int     `json:"target_sum"`
	ToMove					int `json:"to_move"`
	ExpectedArray          []int   `json:"expected_array"`
	ExpectedNestedArrayInt [][]int `json:"expected_nested_array_int"`
	ExpectedBool           bool    `json:"expected_bool"`
	ExpectedInteger        int     `json:"expected_integer"`
}
