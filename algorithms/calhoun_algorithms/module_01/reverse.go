package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	// new_word := ""
	// for i := len(word) - 1; i >= 0; i-- {
	// 	new_word += string(word[i])
	// }
	// return new_word

	// Builder makes it more efficient
	// var new_word strings.Builder
	// for i := len(word) - 1; i >= 0; i-- {
	// 	new_word.WriteByte(word[i])
	// }
	// return new_word.String()
	new_word := ""
	for _, letter := range word {
		new_word = string(letter) + new_word
	}
	return new_word
}
