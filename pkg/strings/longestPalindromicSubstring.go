package strings

func LongestPalindromicSubstring(str string) string {
	longest := ""

	for i := 0; i < len(str); i++ {
		current := string(str[i])
		l := i - 1
		r := i + 1

		for r < len(str) && str[i] == str[r] {
			current += string(str[r])
			r++
		}

		for l >= 0 && r < len(str) && str[l] == str[r] {
			current = string(str[l]) + current + string(str[r])
			l--
			r++
		}

		if len(current) > len(longest) {
			longest = current
		}
	}
	return longest
}
