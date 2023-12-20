package two

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputValueMap = make(map[string][]string)

func findNextNode(currentNodes []string, currentStep int) []string {
	nextNodes := []string{}

	for _, node := range currentNodes {
		nextNode := inputValueMap[node][currentStep]
		nextNodes = append(nextNodes, nextNode)
	}

	return nextNodes
}

// TODO Need to revisit this solution after learning heap
func Two() {
	fmt.Println("Part 2 is running...")

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

	for _, input := range inputArr {
		inputLineArr := strings.Split(input, "=")
		inputKey := strings.TrimSpace(inputLineArr[0])
		inputValueStr := inputLineArr[len(inputLineArr)-1]
		inputValueStrArrBracketSeparated := strings.Trim(inputValueStr, " ()")
		inputValueStrArrCommaSeparated := strings.Split(inputValueStrArrBracketSeparated, ", ")

		inputValueMap[inputKey] = inputValueStrArrCommaSeparated
	}

	steps := 0
	areNodesEndingWithZ := false

	keys := make([]string, 0, len(inputValueMap))
	for k := range inputValueMap {
		if strings.HasSuffix(k, "A") {
			keys = append(keys, k)
		}
	}

	fmt.Println("keys", keys)

	currentNodes := keys
	for !areNodesEndingWithZ {
		// fmt.Println("Inside the loop")
		steps += 1
		// Mimic the behavior of Queue
		currentStep := directionsQueue[0]
		directionsQueue = directionsQueue[1:]
		directionsQueue = append(directionsQueue, currentStep)

		nextNodes := findNextNode(currentNodes, currentStep)

		isCurrentNodeEndsWithZ := true
		for _, node := range nextNodes {
			isCurrentNodeEndsWithZ = strings.HasSuffix(node, "Z")
			if !isCurrentNodeEndsWithZ {
				break
			}
		}

		if isCurrentNodeEndsWithZ {
			areNodesEndingWithZ = true
			break
		} else {
			currentNodes = nextNodes
		}
		fmt.Println("Steps", steps)
	}

	fmt.Println("Steps", steps)
}
