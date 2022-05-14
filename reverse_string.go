package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	l := len(runes)
	reverse := make([]rune, l)
	for inx, r := range runes {
		reverse[l-inx-1] = r
	}
	output = string(reverse)
	return output
}
