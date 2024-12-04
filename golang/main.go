package main

import (
	"github.com/chess-seventh/aoc_2024/cmd"
	"github.com/chess-seventh/aoc_2024/days"
)

func main() {
	var daySelector = cmd.Execute()
	var fileContent = cmd.ReadInputFile(daySelector)
	var goodFileContent = cmd.ReadInputFileGood(daySelector)
	if daySelector == "01" {
		days.Day01(fileContent)
	} else if daySelector == "02" {
		days.Day02(goodFileContent)
	} else if daySelector == "03" {
		days.Day03(fileContent)
	}

}
