package internal

import "testing"

func TestStringReversalSimple(t *testing.T) {
	input := "olleh"
	expected := "hello"
	result := StringReversalSimple(input)
	if result != expected {
		t.Fatalf(
			"StringReversalSimple failed to reverse string. Expected: %s, Actual: %s",
			expected,
			result,
		)
	}
}

func TestStringReversalIndexReversal(t *testing.T) {
	input := "elpoep"
	expected := "people"
	result := StringReversalIndexReversal(input)
	if result != expected {
		t.Fatalf(
			"StringReversalSimple failed to reverse string. Expected: %s, Actual: %s",
			expected,
			result,
		)
	}
}
