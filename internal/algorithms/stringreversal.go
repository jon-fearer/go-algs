package algorithms

// StringReversalSimple uses a for loop to reverse the string
func StringReversalSimple(input string) string {
	output := ""
	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])
	}
	return output
}

// StringReversalIndexReversal uses a for loop with character index reversal
// to reverse the string
func StringReversalIndexReversal(input string) string {
	output := []rune(input)
	start := 0
	end := len(input) - 1
	for start < end {
		output[start] = rune(input[end])
		output[end] = rune(input[start])
		start++
		end--
	}
	return string(output)
}
