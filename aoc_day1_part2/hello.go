package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type KeyValue struct {
	value string
	index int
}

// func findRealIndex(mainStr string, subString string, index int) int {
// 	fmt.Println("mainStr \n", mainStr)
// 	if len(mainStr) == 0 {
// 		return 0
// 	}

// 	subStringLen := len(subString)

// 	result := strings.LastIndex(mainStr[index:], subString)

// 	if result == -1 {
// 		return -1
// 	}

// 	fmt.Println("result \n", result)

// 	result = result + findRealIndex(mainStr[result+subStringLen:], subString, result+subStringLen)

// 	return result
// }

func main() {

	var output []string

	file, _ := os.Open("input.txt")

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var inputArr []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		inputArr = append(inputArr, line)
	}
	// recursiveResult := findRealIndex("two1ninetwo", "two", 0)
	// fmt.Println("recursiveResult", recursiveResult)
	digitsMap := make(map[string]string)
	digitsMap["one"] = "1"
	digitsMap["two"] = "2"
	digitsMap["three"] = "3"
	digitsMap["four"] = "4"
	digitsMap["five"] = "5"
	digitsMap["six"] = "6"
	digitsMap["seven"] = "7"
	digitsMap["eight"] = "8"
	digitsMap["nine"] = "9"

	keys := make([]string, 0, len(digitsMap))
	for k := range digitsMap {
		keys = append(keys, k)
	}

	for _, input := range inputArr {
		var keyValPairs []KeyValue

		for _, k := range keys {
			lastkIdx := strings.LastIndex(input, k)
			if lastkIdx != -1 {
				keyVal := digitsMap[k]
				keyValPairs = append(keyValPairs, KeyValue{keyVal, lastkIdx})
			}
			firstkIdx := strings.Index(input, k)
			if firstkIdx != -1 {
				keyVal := digitsMap[k]
				keyValPairs = append(keyValPairs, KeyValue{keyVal, firstkIdx})
			}
		}

		for idx, char := range input {
			currentChar := string(char)
			_, err := strconv.Atoi(currentChar)
			if err == nil {
				keyValPairs = append(keyValPairs, KeyValue{currentChar, idx})
			}
		}

		//sort the keyValPairs using val
		sort.Slice(keyValPairs, func(i, j int) bool {
			return keyValPairs[i].index < keyValPairs[j].index
		})

		temp := keyValPairs[0].value + keyValPairs[len(keyValPairs)-1].value
		// fmt.Println("debug", temp, input)
		output = append(output, temp)
	}

	// fmt.Println("output", output)

	var sum int = 0

	for _, outputVal := range output {
		intVal, _ := strconv.Atoi(outputVal)
		sum += intVal
	}

	fmt.Println("output", sum)
}
