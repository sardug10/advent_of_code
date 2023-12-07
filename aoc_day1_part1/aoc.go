package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	file, _ := os.Open("input.txt")

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var input []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		input = append(input, line)
	}
	var output []string

	for _, value := range input {

		var temp string = ""
		var reverseValue string = ""

		for _, char := range value {
			currentChar := string(char)
			_, err := strconv.Atoi(currentChar)

			if err == nil {
				temp = currentChar + temp
				break
			}
		}

		for _, char := range value {
			currentChar := string(char)
			reverseValue = currentChar + reverseValue
		}

		for _, char := range reverseValue {
			currentChar := string(char)
			_, err := strconv.Atoi(currentChar)

			if err == nil {
				temp = temp + currentChar
				break
			}
		}

		output = append(output, temp)

	}

	var sum int = 0

	for _, outputVal := range output {
		intVal, _ := strconv.Atoi(outputVal)
		sum += intVal
	}

	fmt.Println("Output", sum)
}
