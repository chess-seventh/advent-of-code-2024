package days

import (
	"fmt"
	"github.com/chess-seventh/aoc_2024/cmd"
	"strconv"
	"strings"
)

func Day01(fileContent string) {
	var column1 []int
	var column2 []int
	var part01Result = Day01part01(fileContent, column1, column2)
	var part02Result = Day01part02(fileContent, column1, column2)
	fmt.Println(part01Result)
	fmt.Println(part02Result)
}

func Day01part01(fileContent string, column1 []int, column2 []int) int {
	var splitString = strings.Fields(fileContent)

	column1, column2 = ExtractColumns(splitString, column1, column2)

	column1 = cmd.SortColumn(column1)
	column2 = cmd.SortColumn(column2)
	return FindDistance(column1, column2)

}

func Day01part02(fileContent string, column1 []int, column2 []int) int {
	var splitString = strings.Fields(fileContent)
	column1, column2 = ExtractColumns(splitString, column1, column2)
	column1 = cmd.SortColumn(column1)
	column2 = cmd.SortColumn(column2)
	numCount := make(map[int]int)

	for i := 0; i < len(column1); i++ {
		var counter = 0
		for j := 0; j < len(column2); j++ {
			if column1[i] == column2[j] {
				counter++
			}
		}
		numCount[column1[i]] = counter
	}

	var totalSum int = 0

	//fmt.Println(numCount)

	for key, value := range numCount {
		totalSum += key * value
	}
	return totalSum
}

func FindDistance(column1 []int, column2 []int) int {
	var totalSum int
	for i := 0; i < len(column1); i++ {
		var diffCol1Col2 int
		if column1[i] > column2[i] {
			diffCol1Col2 = column1[i] - column2[i]
		} else {
			diffCol1Col2 = column2[i] - column1[i]
		}
		totalSum += diffCol1Col2
	}
	return totalSum
}

func ExtractColumns(splitString []string, column1 []int, column2 []int) ([]int, []int) {
	for i, word := range splitString {
		intWord, _ := strconv.Atoi(word)
		if i%2 == 0 {
			column1 = append(column1, intWord)
		} else {
			column2 = append(column2, intWord)
		}
	}
	return column1, column2
}
