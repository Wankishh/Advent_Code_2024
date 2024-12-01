package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getData() string {
	var file, _ = os.ReadFile("data.in");
	return strings.TrimSpace(string(file));
}

func parseInput(input string) ([]int, []int) {
	var split []string = strings.Split(input, "\n")
	var left []int
	var right []int

	for _, v := range split {
		var split1 []string = strings.Split(v, "   ");
		var leftInt, _ = strconv.Atoi(split1[0]);
		var rightInt, _ = strconv.Atoi(split1[1]);
		left = append(left, leftInt);
		right = append(right, rightInt);
	}

	sort.Ints(left);
	sort.Ints(right);

	return left, right;
}

func firstTask(data []int, data1 []int) int {
	var difference int = 0

	for index, v := range data {
		var other int = data1[index];
		result := other - v;
		difference += int(math.Abs(float64(result)));
	}

	return difference
}

func secondTask(left []int, right []int) int {
	var total int = 0

	for _, leftItem := range left { // O(n^2)
		var totalFound int = 0;
		for _, rightItem := range right {
			if leftItem == rightItem {
				totalFound += 1;
			}
		}

		total += leftItem * totalFound;
	}

	return total
}



func main() {
	var data = getData()
	var leftList, rightList = parseInput(data)

	var firstTaskResult = firstTask(leftList, rightList)
	var secondTaskResult = secondTask(leftList, rightList)

	fmt.Println(`Task 1`, firstTaskResult)
	fmt.Println(`Task 2`, secondTaskResult)
}
