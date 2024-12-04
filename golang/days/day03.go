package days

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day03(fileContent string) {
	day03part01(fileContent)
	day03part02(fileContent)
}

func day03part01(input string) {
	fmt.Println("Solution Part One:", solvePuzzlePartOne(input))
}

func day03part02(input string) {
	fmt.Println("Solution Part Two:", solvePuzzlePartTwo(input))
}

// solvePuzzle takes the input and returns the solution to the puzzle.
func solvePuzzlePartOne(input string) int {
	instructions := parseInstructions(input)
	return calculateSum(instructions)
}

// solvePuzzlePartTwo takes the input and returns the solution for Part Two.
func solvePuzzlePartTwo(input string) int {
	instructions := parseAllInstructions(input)
	return calculateSumWithState(instructions)
}

// parseInstructions extracts potential `mul` instructions from the input.
func parseInstructions(input string) []string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	return re.FindAllString(input, -1)
}

// executeInstruction parses and calculates the result of a single `mul` instruction.
func executeInstruction(instruction string) (int, error) {
	// Trim the "mul(" prefix and ")" suffix
	cleaned := strings.TrimSuffix(strings.TrimPrefix(instruction, "mul("), ")")
	parts := strings.Split(cleaned, ",")

	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid instruction format: %s", instruction)
	}

	// Convert parts to integers
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid number X: %s", parts[0])
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid number Y: %s", parts[1])
	}

	return x * y, nil
}

// calculateSum computes the sum of all valid `mul` results.
func calculateSum(instructions []string) int {
	total := 0
	for _, inst := range instructions {
		result, err := executeInstruction(inst)
		if err == nil {
			total += result
		}
	}
	return total
}

// parseAllInstructions extracts all valid instructions (mul, do, don't) from the input.
func parseAllInstructions(input string) []string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	return re.FindAllString(input, -1)
}

// calculateSumWithState computes the sum of results from enabled `mul` instructions.
func calculateSumWithState(instructions []string) int {
	total := 0
	enabled := true // Initial state: `mul` instructions are enabled.

	for _, inst := range instructions {
		if inst == "do()" {
			enabled = true
		} else if inst == "don't()" {
			enabled = false
		} else if strings.HasPrefix(inst, "mul(") && enabled {
			result, err := executeInstruction(inst)
			if err == nil {
				total += result
			}
		}
	}

	return total
}
