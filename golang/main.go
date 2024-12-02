package main

import (
	"fmt"
	"github.com/chess-seventh/aoc_2024/cmd"
	"github.com/chess-seventh/aoc_2024/days"
	"log"
	"os"
)

func main() {
	var daySelector = cmd.Execute()
	var fileContent = readInputFile(daySelector)
	var column1 []int
	var column2 []int

	if daySelector == "01" {

		days.Day01(fileContent, column1, column2)
		//for _, word := range strings.Fields(fileContent) {
		//	fmt.Println(word)
		//}
	}

}

func readInputFile(daySelector string) string {
	filename := fmt.Sprintf("inputs/day_%s.txt", daySelector)
	content, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
