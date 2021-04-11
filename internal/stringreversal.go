package internal

import "fmt"

func StringReversal() {
	input := "olleh"
	output := ""
	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])
	}
	fmt.Println(output)
}
