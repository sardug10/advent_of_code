package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func contaiins(arr []string, ele int) bool {
	for _, char := range arr {
		currentNum, _ := strconv.Atoi(char)
		if currentNum == ele {
			return true
		}
	}
	return false
}

func main() {
	var output []int

	file, _ := os.Open("input.txt")

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var inputArr []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		inputArr = append(inputArr, line)
	}

	// input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	for _, input := range inputArr {

		game := strings.Split(input, ":")

		gameNumbers := strings.Split(game[1], "|")

		winningNumbers := strings.Split(gameNumbers[0], " ")
		winningNumbersPointer := &winningNumbers
		myNumbers := strings.Split(gameNumbers[1], " ")

		currentInputResult := 0

		for _, myChar := range myNumbers {
			trimmedMychar := strings.TrimSpace(myChar)
			trimmedMycharNum, error := strconv.Atoi(trimmedMychar)

			if error == nil {
				isInWinningNumbers := contaiins(*winningNumbersPointer, trimmedMycharNum)
				if isInWinningNumbers {
					if currentInputResult == 0 {
						currentInputResult += 1
					} else {
						currentInputResult = currentInputResult * 2
					}
				}
			}
		}

		output = append(output, currentInputResult)

	}

	result := 0
	for _, currentResult := range output {
		result += currentResult
	}

	fmt.Println("Current Input Result: ", result)

}
