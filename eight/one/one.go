package one

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func One() {
	fmt.Println("Part 1 is running...")

	file, _ := os.Open("./input.txt")

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var inputArr []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		inputArr = append(inputArr, line)
	}

	directionsInput := "LLRRLRRRLLRLRRLLLLRLRRLRRRLRLRRRLLRRLRRRLLRRLRRLRRLLRRRLRRLRRLRRRLRRLRLRLRRLRRLRRRLLRRLLLRRLRRRLRRRLRRRLRRLRRRLRLLRLRRRLRLRRLLRLRRRLRRRLRLRRRLRRRLRLRLRRLRRLRLRRLLRRRLRRRLRRRLLRRRLRLRLRLRLLRRRLRRRLRRLRRRLLRLRRLRRLRRRLRRRLRRLRLRLRRRLRRLRRLRRRLLRRLRLRLRRRLRLRLRRLRRLLRRLRRRLLRLLRLRLRRRR"
	directionsQueue := []int{}

	for _, direction := range directionsInput {
		directionStr := string(direction)
		if directionStr == "R" {
			directionsQueue = append(directionsQueue, 1)
		} else {
			directionsQueue = append(directionsQueue, 0)
		}
	}

	inputValueMap := make(map[string][]string)

	for _, input := range inputArr {
		inputLineArr := strings.Split(input, "=")
		inputKey := strings.TrimSpace(inputLineArr[0])
		inputValueStr := inputLineArr[len(inputLineArr)-1]
		inputValueStrArrBracketSeparated := strings.Trim(inputValueStr, " ()")
		inputValueStrArrCommaSeparated := strings.Split(inputValueStrArrBracketSeparated, ", ")

		inputValueMap[inputKey] = inputValueStrArrCommaSeparated
	}

	steps := 0
	current := "AAA"

	for current != "ZZZ" {
		steps += 1

		// Mimic the behavior of Queue
		currentStep := directionsQueue[0]
		directionsQueue = directionsQueue[1:]
		directionsQueue = append(directionsQueue, currentStep)

		current = inputValueMap[current][currentStep]
	}

	fmt.Println("Steps", steps)
}
