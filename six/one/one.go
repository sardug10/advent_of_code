package one

import "fmt"

func One() {
	fmt.Println("Part 1 is running")

	timeInput := []int{55999793}
	distanceInput := []int{401148522741405}
	output := []int{}

	for index, time := range timeInput {
		distancetoBeat := distanceInput[index]
		beatCount := 0

		for htt := 0; htt <= time; htt++ {
			speed := htt
			distanceCovered := speed * (time - htt)

			if distanceCovered > distancetoBeat {
				beatCount += 1
			}

		}

		output = append(output, beatCount)
	}

	result := 1
	for _, count := range output {
		result = result * count
	}
	fmt.Println("result: ", result)

}
