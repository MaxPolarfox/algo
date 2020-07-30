package dinamic_programing

// Write a function that takes in two strings and returns their longest common subsequence. A subsequence of a string is
// a set of characters that aren't necessarily adjacent in the string but that are in the same order as they appear in the string.
// For instance, the characters ["a", "c" , "d"] form a subsequence of the string "absd", and so do the characters ["b", "d"].
// Note that a single character in a string and the string itself are both valid subsequences of the string.
// You can assume that there will only be one longest common subsequence.

// Sample Input: strl = "ZXVVYZW" str2 = " XKYKZPW"
// Sample Output: ["X", "Y", "Z" "W"]

// Time: O(nm) Space: O(nm)
func LongestCommonSubsequence(s1 string, s2 string) string {
	lengths := make([][]int, len(s2)+1)
	for i := range lengths {
		lengths[i] = make([]int, len(s1)+1)
	}
	for i := 1; i < len(s2)+1; i++ {
		for j := 1; j < len(s1)+1; j++ {
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
		} else if lengths[i][j] == lengths[i][j-1] {
			j--
		} else {
			sequence = append([]byte{s1[j-1]}, sequence...)
			i--
			j--
		}
	}
	return string(sequence)
}
