package days

import (
	"testing"
)

func TestParseInstructions(t *testing.T) {
	input := "mul(2,4)randomtextmul(3,7)!mul(5,5)"
	expected := []string{"mul(2,4)", "mul(3,7)", "mul(5,5)"}
	result := parseInstructions(input)

	if len(result) != len(expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %s, got %s", expected[i], v)
		}
	}
}

func TestExecuteInstruction(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		hasError bool
	}{
		{"mul(2,4)", 8, false},
		{"mul(10,5)", 50, false},
		{"mul(123,4)", 492, false},
		{"mul(4*)", 0, true},  // Invalid format
		{"mul(x,4)", 0, true}, // Invalid number
	}

	for _, test := range tests {
		result, err := executeInstruction(test.input)
		if test.hasError && err == nil {
			t.Errorf("expected an error for input %s", test.input)
		}
		if !test.hasError && result != test.expected {
			t.Errorf("for input %s, expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestCalculateSum(t *testing.T) {
	instructions := []string{"mul(2,4)", "mul(3,7)", "mul(5,5)"}
	expected := 62
	result := calculateSum(instructions)

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestSolvePuzzle(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := 161
	result := solvePuzzle(input)

	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}
