package main

import (
	"bufio"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	daily "github.com/uncleBlobby/aoc2024/internal"
)

func main() {
	fmt.Println("Hello, aoc2024 day 11!")

	input := daily.GetInputFromFile("testdata")

	scanner := bufio.NewScanner(input)

	startInput := []int{}

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		stones := strings.Split(scanner.Text(), " ")

		for _, stone := range stones {
			num, err := strconv.Atoi(stone)
			if err != nil {
				log.Printf("error converting int: %s", err)
				continue
			}
			startInput = append(startInput, num)
		}
	}

	fmt.Println(startInput)
	stones := Blink(startInput)
	stones = Blink(stones)
	stones = Blink(stones)

	// testSampleArray := []int{1, 2, 3, 4, 5}

	// fmt.Println(testSampleArray)
	// testSampleArray = slices.Insert(testSampleArray, 2, 999)

	// fmt.Println(testSampleArray)

}

func Blink(inputArray []int) []int {
	output := make([]int, len(inputArray))

	for i := 0; i < len(inputArray); i++ {
		resultElem := ProcessStone(inputArray[i])
		if len(resultElem) == 1 {
			output[i] = resultElem[0]
		}
		if len(resultElem) > 1 {
			output[i] = resultElem[0]
			output = slices.Insert(output, i+1, resultElem[1])
			i++
		}
	}
	fmt.Println(output)
	return output

	// for ind, elem := range inputArray {
	// 	resultElem := ProcessStone(elem)
	// 	if len(resultElem) == 1 {
	// 		output[ind] = resultElem[0]
	// 	}
	// 	if len(resultElem) > 1 {
	// 		log.Printf("applied rule two on elem[%d]: resultElems: [%d] [%d] to be inserted at index [%d] [%d]", elem, resultElem[0], resultElem[1], ind, ind+1)
	// 		output[ind] = resultElem[0]
	// 		output = slices.Insert(output, ind+1, resultElem[1])
	// 		// output = append(output, 0)
	// 		// copy(output[ind+2:], output[ind+1:])
	// 		// output[ind+1] = resultElem[1]
	// 		// output = append(output[:ind+2], output[ind+1:]...)
	// 		// output[ind+1] = resultElem[1]
	// 		// output = output[:len(output)-1]
	// 		// fmt.Println(newOutput)
	// 		// return newOutput
	// 	}
	// }
	// fmt.Println(output)
	// return output
}

func ProcessStone(inputValue int) []int {
	one, done := RuleOne(inputValue)
	if one {
		return []int{done}
	}
	two, done2 := RuleTwo(inputValue)
	if two {
		return []int{done2[0], done2[1]}
	}
	three := RuleThree(inputValue)
	return []int{three}
}

func RuleOne(input int) (bool, int) {
	if input == 0 {
		return true, 1
	} else {
		return false, input
	}
}

func RuleTwo(input int) (bool, []int) {
	output := []int{}

	inputAsStr := strconv.Itoa(input)
	digitCount := len(inputAsStr)
	if digitCount == 1 {
		return false, nil
	}

	if digitCount > 1 && digitCount%2 == 0 {
		// 7654 -> length 4, 2 digits each new term
		part1 := inputAsStr[:digitCount/2]
		// convert back to num

		part2 := inputAsStr[digitCount/2:]

		// drop leading zeros from second part, if any exist

		// count leading zeros
		zeroCount := 0

		for i := 0; i < len(part2); i++ {
			if part2[i] == '0' && i != len(part2)-1 {
				zeroCount += 1
			} else {
				break
			}
		}
		part2Trimmed := part2[zeroCount:]

		// convert back to num
		// log.Printf("new stones: %s | %s\n", part1, part2)
		// log.Printf("part2 trimmed: %s\n", part2Trimmed)

		output = append(output, StringToInt(part1))
		output = append(output, StringToInt(part2Trimmed))
		return true, output
	}

	return false, nil
}

func RuleThree(input int) int {
	return input * 2024
}

func StringToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}
