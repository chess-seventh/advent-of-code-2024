package days

import (
	"fmt"
	"strconv"
	"strings"
)

func Day02(fileContent []string) {
	part01(fileContent)
}

func part01(fileContent []string) {
	var safeLevels = 0
	for line := range fileContent {
		safeLevels += isLevelSafe(fileContent[line])
	}
	fmt.Println("Part 01:", safeLevels)
}

func part02(fileContent []string) {
	var safeLevels = 0
	for line := range fileContent {
		safeLevels += isLevelSafe(fileContent[line])
	}
}

func isLevelSafe(level string) int {
	splitStringInt := convertLineToInt(level)
	if !checkListOrdered(splitStringInt) {
		return 0
	}
	return 1

}

func checkListOrdered(splitStringInt []int) bool {
	var eitherIncOrDec bool
	if checkListIsIncreasing(splitStringInt) {
		return true
	} else {
		eitherIncOrDec = false
	}

	if checkListIsDecreasing(splitStringInt) {
		return true
	} else {
		eitherIncOrDec = false
	}
	return eitherIncOrDec
}

func checkListIsIncreasing(level []int) bool {
	for i := 0; i < len(level)-1; i++ {
		if level[i] > level[i+1] || level[i] == level[i+1] {
			return false
		} else {
			var distance = calculateDistance(level[i], level[i+1])
			if !checkIfDistanceIsInBounds(distance) {
				return false
			}
		}
	}
	return true
}

func checkListIsDecreasing(level []int) bool {
	for i := 0; i < len(level)-1; i++ {
		if level[i] < level[i+1] || level[i] == level[i+1] {
			return false
		} else {
			var distance = calculateDistance(level[i], level[i+1])
			if !checkIfDistanceIsInBounds(distance) {
				return false
			}
		}
	}
	return true
}

func checkIfDistanceIsInBounds(tmpVal int) bool {
	if tmpVal >= 1 {
		if tmpVal <= 3 {
			return true
		}
	} else {
		return false
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
