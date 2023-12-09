package part2

import (
	"aoc_day4/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() {
	var output = 0

	file, _ := os.Open("input.txt")

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var inputArr []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		inputArr = append(inputArr, line)
	}

	cardsCountMap := make(map[int]int)
	inputArrLength := len(inputArr)

	for i := 1; i <= inputArrLength; i++ {
		cardsCountMap[i] = 1
	}

	// input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	for index, input := range inputArr {

		game := strings.Split(input, ":")
		currentOperatingCard := index + 1

		gameNumbers := strings.Split(game[1], "|")

		winningNumbers := strings.Split(gameNumbers[0], " ")
		winningNumbersPointer := &winningNumbers
		myNumbers := strings.Split(gameNumbers[1], " ")

		winningNumberCount := 0

		for _, myChar := range myNumbers {
			trimmedMychar := strings.TrimSpace(myChar)
			trimmedMycharNum, error := strconv.Atoi(trimmedMychar)

			if error == nil {
				isInWinningNumbers := utils.Contains(*winningNumbersPointer, trimmedMycharNum)
				if isInWinningNumbers {
					winningNumberCount += 1
				}
			}
		}

		currentCardCount := cardsCountMap[currentOperatingCard]

		for i := currentOperatingCard + 1; i <= winningNumberCount+currentOperatingCard; i++ {
			cardsCountMap[i] = cardsCountMap[i] + currentCardCount
		}

	}

	for _, value := range cardsCountMap {
		output += value
	}

	fmt.Println("Part 2 ", output)

}
