package reverse_string

func ReverseString(input string) (output string) {
	n := []rune(input)
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		j := len(n) - 1 - i
		n[i], n[j] = n[j], n[i]
		output = string(n)
	}
	return output
}
