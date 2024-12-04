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

func parseInput(input string) [][]int {
	var lines = strings.Split(input, "\n")

	var list [][]int

	for _, line := range lines {
		var numbers = strings.Split(line, " ")
		var secondaryList []int
		for _, v := range numbers {
			var num, _ = strconv.Atoi(v)
			secondaryList = append(secondaryList, num)
		}

		list = append(list, secondaryList)
	}

	return list
}

func firstTask(data [][]int) int {
	var correctLines int = 0

	for _, numbers := range data {
		var isCorrectList bool = true
		var length = len(numbers)
		var action string

		for i := 0; i < length; i++ {
			if i == 0 {
				continue
			}

			var currentNumber = numbers[i]
			var prevIndex = i - 1
			var prevNumber = numbers[prevIndex]
			var currentAction string

			var totalDifference = prevNumber - currentNumber
			if totalDifference < 0 {
				currentAction = "decrement"
				totalDifference = totalDifference * -1
			} else {
				currentAction = "increment"
			}

			if action == "" {
				action = currentAction
			} else if action != currentAction {
				isCorrectList = false
				break
			}

			if totalDifference < 1 || totalDifference > 3 {
				isCorrectList = false
				break
			}
		}

		if isCorrectList {
			correctLines++
		}
	}

	return correctLines
}

func isCurrentListSafe(levels []int) bool {
	var isCorrectList bool = true
	var length = len(levels)
	var action string

	for i := 0; i < length; i++ {
		if i == 0 {
			continue
		}

		var currentNumber = levels[i]
		var prevIndex = i - 1
		var prevNumber = levels[prevIndex]
		var currentAction string

		var totalDifference = prevNumber - currentNumber
		if totalDifference < 0 {
			currentAction = "decrement"
			totalDifference = totalDifference * -1
		} else {
			currentAction = "increment"
		}

		if action == "" {
			action = currentAction
		} else if action != currentAction {
			isCorrectList = false
			break
		}

		if totalDifference < 1 || totalDifference > 3 {
			isCorrectList = false
			break
		}
	}

	return isCorrectList
}

func changeArray(numbers []int, index int) []int {
	var list []int
	for i := 0; i < len(numbers); i++ {
		if(i == index) {
			continue
		}

		list = append(list, numbers[i])
	}

	return list
}

func secondTask(data [][]int) int {
	var correctLines int = 0

	for _, levels := range data {
		var isCorrectList = isCurrentListSafe(levels)
		
		if(isCorrectList) {
			correctLines++
			continue
		}

		for i := 0; i < len(levels); i++ {
			var newLevels = changeArray(levels, i)
			isCorrectList = isCurrentListSafe(newLevels)

			if(isCorrectList) {
				correctLines++
				break
			}
		}
	}

	return correctLines
}

func main() {
	var data string = getData("data")
	var list [][]int = parseInput(data)

	var firstTaskResult = firstTask(list)
	var secondTaskResult = secondTask(list)

	// if(secondTaskResult != 4) {
	// 	fmt.Println(`Task 2`, secondTaskResult, `Expected 4`)
	// 	panic("Second task failed")
	// }

	fmt.Println(`Task 1`, firstTaskResult)
	fmt.Println(`Task 2`, secondTaskResult)
}
