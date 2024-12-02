package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Day02(fileContent []string) {
	part01(fileContent)
	part02(fileContent)
}

func part01(fileContent []string) {
	var safeLevels = 0
	for line := range fileContent {
		splitStringInt := convertLineToInt(fileContent[line])
		safeLevels += isLevelSafe(splitStringInt)
	}
	fmt.Println("Part 01:", safeLevels)
}

func isLevelSafe(level []int) int {
	valInc, _ := checkListIsIncreasing(level)
	valDec, _ := checkListIsDecreasing(level)
	if valInc || valDec {
		return 1
	}
	return 0
}

func checkListIsIncreasing(level []int) (bool, int) {
	var mistakesCount = 0
	var returnValue = true
	for i := 0; i < len(level)-1; i++ {
		if level[i] > level[i+1] {
			returnValue = false
			mistakesCount++
		}
		if level[i] == level[i+1] {
			returnValue = false
			mistakesCount++
		} else {
			var distance = calculateDistance(level[i], level[i+1])
			if !checkIfDistanceIsInBounds(distance) {
				mistakesCount++
				returnValue = false
			}
		}
	}
	return returnValue, mistakesCount
}

func checkListIsDecreasing(level []int) (bool, int) {
	var mistakesCount = 0
	var returnValue = true

	for i := 0; i < len(level)-1; i++ {
		if level[i] < level[i+1] {
			returnValue = false
			mistakesCount++
		}
		if level[i] == level[i+1] {
			returnValue = false
			mistakesCount++
		} else {
			var distance = calculateDistance(level[i], level[i+1])
			if !checkIfDistanceIsInBounds(distance) {
				returnValue = false
				mistakesCount++
			}
		}
	}
	return returnValue, mistakesCount
}

func checkIfDistanceIsInBounds(value int) bool {
	if value >= 1 && value <= 3 {
		return true
	}
	return false
}

func calculateDistance(i int, j int) int {
	var tmpVal = i - j
	if tmpVal < 0 {
		tmpVal = -tmpVal
	}
	return tmpVal
}

func convertLineToInt(level string) []int {
	var splitString = strings.Fields(level)
	var splitStringInt []int
	for _, i := range splitString {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		splitStringInt = append(splitStringInt, j)
	}
	return splitStringInt
}

func part02(lines []string) {
	var report [][]int
	score := 0

	for i := 0; i < len(lines); i++ {
		currentLine := splitLines(lines[i], " ")
		var currentLineAsInt []int
		for j := 0; j < len(currentLine); j++ {
			currentLineAsInt = append(currentLineAsInt, stringToNumber(currentLine[j]))
		}
		report = append(report, currentLineAsInt)
	}

	for i := 0; i < len(report); i++ {
		goodResult := false
		testReport := report[i]

		for j := 0; j < len(testReport); j++ {
			testSlice := make([]int, 0, len(testReport)-1)
			testSlice = append(testSlice, testReport[:j]...)
			testSlice = append(testSlice, testReport[j+1:]...)

			if checkReport(testSlice) {
				goodResult = true
			}
		}

		if goodResult {
			score += 1
		}

	}

	println(score)

}

func checkReport(report []int) bool {
	isValidResult := true

	if report[0] < report[1] {
		for j := 1; j < len(report); j++ {
			if report[j] < report[j-1] {
				isValidResult = false

			}
		}
	} else if report[0] > report[1] {
		for j := 1; j < len(report); j++ {
			if report[j] > report[j-1] {
				isValidResult = false

			}
		}
	} else {
		isValidResult = false

	}

	for j := 1; j < len(report); j++ {
		if math.Abs(float64(report[j]-report[j-1])) > 0 &&
			math.Abs(float64(report[j]-report[j-1])) < 4 {
			// pass
		} else {
			isValidResult = false
		}
	}

	if isValidResult == true {
		return true
	}
	return false
}

func splitLines(line string, deliminator string) []string {
	eachCharString := strings.Split(line, deliminator)
	var cleanCharString []string
	for i := 0; i < len(eachCharString); i++ {
		if eachCharString[i] != "" {
			cleanCharString = append(cleanCharString, eachCharString[i])
		}
	}
	return cleanCharString
}

func stringToNumber(str string) int {
	stringAsInt, _ := strconv.Atoi(str)
	return stringAsInt
}
