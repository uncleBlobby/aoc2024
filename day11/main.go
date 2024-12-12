package main

import (
	"bufio"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"time"

	daily "github.com/uncleBlobby/aoc2024/internal"
)

func main() {
	fmt.Println("Hello, aoc2024 day 11!")

	input := daily.GetInputFromFile("input")

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

	blinkCount := 75
	dupeCount := 0
	Part2(startInput, blinkCount, dupeCount)

	// fmt.Printf("Result after blinking %d times:\n%v\n", blinkCount, result)
	// fmt.Printf("Final count of stones: %d\n", len(result))
	// testSampleArray := []int{1, 2, 3, 4, 5}

	// fmt.Println(testSampleArray)
	// testSampleArray = slices.Insert(testSampleArray, 2, 999)

	// fmt.Println(testSampleArray)

}

func Part2(inputArray []int, X int, dupeStoneCount int) {

	for i := 0; i < X; i++ {
		// fmt.Printf("%d / %d\n", i, X)
		// fmt.Println(inputArray)
		// startLen := len(inputArray)
		inputArray = Blink(inputArray)
		startLen := len(inputArray)
		slices.Sort(inputArray)
		inputArray := slices.Compact(inputArray)
		endLen := len(inputArray)
		dupeStoneCount += startLen - endLen
		fmt.Printf("current stone count: %d\n", len(inputArray)+dupeStoneCount)
	}

	fmt.Printf("Total stone count after %d blinks: %d\n", X, len(inputArray)+dupeStoneCount)
}

func BlinkXTimes(inputArray []int, X int) []int {
	fmt.Printf("Blinking %d times...\n", X)
	for i := 0; i < X; i++ {
		startTime := time.Now()
		inputArray = Blink(inputArray)
		duration := time.Since(startTime)
		fmt.Printf("%d / %d -- %d s\n", i, X, duration/1000000000)
	}

	return inputArray
}

func Blink(inputArray []int) []int {
	// fmt.Printf("Preparing output array of length: %d\n", len(inputArray))
	output := make([]int, len(inputArray))

	for i := 0; i < len(inputArray); i++ {
		output[i] = inputArray[i]
	}

	slicedIn := 0

	for i := 0; i < len(inputArray); i++ {
		// fmt.Printf("Processing element [%d]: %d\n", i, inputArray[i])
		resultElem := ProcessStone(inputArray[i])
		if len(resultElem) == 1 {
			// fmt.Printf("Inserting [%d] at output position [%d]\n", resultElem[0], i+slicedIn)
			output[i+slicedIn] = resultElem[0]
		}
		if len(resultElem) > 1 {
			// fmt.Printf("Inserting [%d] at output position [%d], and [%d] at output position [%d]\n", resultElem[0], i, resultElem[1], i+1)
			output[i+slicedIn] = resultElem[0]
			output = slices.Insert(output, i+slicedIn+1, resultElem[1])
			slicedIn += 1
			// i++
			// inputArray = output
		}
		// fmt.Println("Current output:")
		// fmt.Println(output)
	}
	// fmt.Println("Final output:")
	// fmt.Println(output)
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
		// fmt.Printf("Triggered RuleOne on inputValue [%d], result: [%d]\n", inputValue, done)
		return []int{done}
	}
	two, done2 := RuleTwo(inputValue)
	if two {

		// fmt.Printf("Triggered RuleTwo on inputValue [%d], result: [%d], [%d]\n", inputValue, done2[0], done2[1])
		return []int{done2[0], done2[1]}
	}
	three := RuleThree(inputValue)
	// fmt.Printf("Triggered RuleThree on inputValue [%d], result: [%d]\n", inputValue, three)
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
