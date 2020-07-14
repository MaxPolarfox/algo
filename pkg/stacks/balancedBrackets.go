package stacks

var opening = map[rune]bool{
	'{': true,
	'[': true,
	'(': true,
}
var closing = map[rune]bool{
	'}': true,
	']': true,
	')': true}
var matching = map[rune]rune{
	'}': '{',
	']': '[',
	')': '('}

func BalancedBrackets(s string) bool {
	stack := []rune{}
	for _, char := range s {
		if opening[char] {
			stack = append(stack, char)
			continue
		}
		if closing[char] {
			if len(stack) == 0 {
				return false
			}
			if matching[char] == stack[len(stack)-1] {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
