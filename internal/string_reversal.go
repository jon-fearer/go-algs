package internal

import "fmt"

func StringReversal() {
	input := "olleh"
	// Simple for loop approach
	output := ""
	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])
	}
	fmt.Println(output)
	// For loop with character index reversal
	input2 := "elpoep"
	output2 := []rune(input2)
	start := 0
	end := len(input2) - 1
	for start < end {
		output2[start] = rune(input2[end])
		output2[end] = rune(input2[start])
		start += 1
		end -= 1
	}
	fmt.Println(string(output2))
}
