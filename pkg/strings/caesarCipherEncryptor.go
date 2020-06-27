package strings

import "strings"

func CaesarCipherEncryptor(str string, key int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	updated := ""

	for i:=0; i < len(str); i++ {
		position := strings.Index(alphabet, string(str[i]))
		updated = updated + string(alphabet[(position + key) % 26])
	}
	return updated
}