package dinamic_programing

func LongestCommonSubsequence(s1 string, s2 string) string {
	lengths := make([][]int, len(s2) + 1)
	for i := range lengths {
		lengths[i] = make([]int, len(s1)+1)
	}

	for i := 1; i < len(s2) +1; i++ {
		for j := 1; j < len(s1) +1 ; j++ {
			if s2[i-1] == s1[j-1] {
				lengths[i][j] = lengths[i-1][j-1] + 1
			} else {
				lengths[i][j] = max(lengths[i-1][j], lengths[i][j-1])
			}
		}
	}

	return buildLongestCommonSubsequence(lengths, s1)
}

func buildLongestCommonSubsequence(lengths [][]int, s1 string) string {
	sequence := make([]byte, 0)
	i := len(lengths) - 1
	j := len(lengths[0]) - 1
	for i != 0 && j != 0 {
		if lengths[i][j] == lengths[i-1][j] {
			i--
		} else if lengths[i][j] == lengths[i][j - 1] {
			j--
		} else {
			sequence = append([]byte{s1[j-1]}, sequence...)
			i--
			j--
		}
	}
	return string(sequence)
}
