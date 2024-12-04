package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getData(name string) string {
	var fileName string = name + ".in"
	var file, _ = os.ReadFile(fileName)
	return strings.TrimSpace(string(file))
}

func includes(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}

	return false
}

func parseInput(input string) []string {
	// Example of correct data for parsing mul(5,10)
	// 1. Should start at every m letter
	// 2. Should end if some of the next letters are not u, l, (, number, ",", number, )
	// 3. Should return a list of the numbers

	var len = len(input)
	var allowedLetters = []string{"m", "u", "l", "(", ")", ",", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	var currentWorkingIndex = 0
	var currentSequence = ""
	var list []string

	// Check if letter is in the list
	// If not break the current sequence and start over until you find the next m

	var sequenceStarted = false
	for i := 0; i < len; i++ {
		var char = string(input[i])

		var isValidChar = includes(allowedLetters, char)
		var isStartChar = char == "m"
		var isEndChar = char == ")"

		if isStartChar {
			currentSequence = char
			sequenceStarted = true
			currentWorkingIndex = 0
			continue
		}

		if isEndChar && sequenceStarted {
			currentSequence += char
			currentWorkingIndex = 0
			sequenceStarted = false
			list = append(list, currentSequence)
			currentSequence = ""
			continue
		}

		if !isValidChar {
			sequenceStarted = false
			currentWorkingIndex = 0
			currentSequence = ""
			continue
		}

		currentWorkingIndex++
		currentSequence += char
	}

	fmt.Println(list)
	return list
}

func firstTask(data []string) int {
	var total int = 0
	for _, item := range data {
		if !strings.Contains(item, "mul") || !strings.Contains(item, "(") || !strings.Contains(item, ")") || !strings.Contains(item, ",") {
			continue
		}
		var removedMul = strings.Replace(item, "mul", "", 1);
		var removedBrackets = strings.Replace(removedMul, "(", "", 1);
		var removeEndBrackets = strings.Replace(removedBrackets, ")", "", 1);
		var split = strings.Split(removeEndBrackets, ",")
		fmt.Println(item, split)
		var first, _ = strconv.Atoi(split[0])
		var second, _ = strconv.Atoi(split[1])
		total += first * second
	}

	return total;
}

func secondTask(data [][]int) int {
	return 0
}

func main() {
	var data string = getData("data")
	var list []string = parseInput(data)

	var firstTaskResult = firstTask(list)
	// var secondTaskResult = secondTask(list)

	// if(firstTaskResult != 161) {
	// 	fmt.Println(`Task 1`, firstTaskResult, `Expected 161`)
	// 	panic("First task failed")
	// }

	fmt.Println(`Task 1`, firstTaskResult)
	// fmt.Println(`Task 2`, secondTaskResult)
}
