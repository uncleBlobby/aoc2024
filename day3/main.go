package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	daily "github.com/uncleBlobby/aoc2024/internal"
)

// REGEX samples
//
// [m][u][l][(]\d{1,3}[,]\d{1,3}[)]
//
// (mul\()\d{1,3}[,]\d{1,3}[)]
//
// DO
// (?:do\(\))
//
// DON'T
// (?:don't\(\))

func main() {
	fmt.Println("hello, aoc-2024 day3!")

	r, err := regexp.Compile("(mul\\()\\d{1,3}[,]\\d{1,3}[)]")
	if err != nil {
		log.Printf("regex compile: %s", err)
		os.Exit(1)
	}

	r2, err := regexp.Compile("\\d{1,3}")
	if err != nil {
		log.Printf("regex compile: %s", err)
		os.Exit(1)
	}

	doInst, err := regexp.Compile("(?:do\\(\\))")
	if err != nil {
		log.Printf("regex compile: %s", err)
		os.Exit(1)
	}

	donotInst, err := regexp.Compile("(?:don't\\(\\))")
	if err != nil {
		log.Printf("regex compile: %s", err)
		os.Exit(1)
	}

	input := daily.GetInputFromFile("input")

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var s strings.Builder
	for scanner.Scan() {
		_, err := s.WriteString(scanner.Text())
		if err != nil {
			log.Printf("error reading file to string: %s", err)
			os.Exit(1)
		}
		//fmt.Println(scanner.Text())
	}

	fmt.Println(s.String())

	// find all index of matches of do and don't in set to track on/off state

	completeInstructionList := []Instruction{}

	doInstructions := doInst.FindAllString(s.String(), -1)
	doIndexes := doInst.FindAllStringIndex(s.String(), -1)

	for i := 0; i < len(doInstructions); i++ {
		i := Instruction{
			Name:       DO,
			StartIndex: doIndexes[i][0],
			EndIndex:   doIndexes[i][1],
		}
		completeInstructionList = append(completeInstructionList, i)
	}

	for _, do := range doIndexes {
		log.Printf("do matches: %v", do)
	}

	dontInstructions := donotInst.FindAllString(s.String(), -1)
	dontIndexes := donotInst.FindAllStringIndex(s.String(), -1)

	for i := 0; i < len(dontInstructions); i++ {
		i := Instruction{
			Name:       DONT,
			StartIndex: dontIndexes[i][0],
			EndIndex:   dontIndexes[i][1],
		}
		completeInstructionList = append(completeInstructionList, i)
	}

	for _, dont := range dontIndexes {
		log.Printf("dont matches: %v", dont)
	}

	//dos := doInst.FindStringIndex()

	multiplyInstructions := r.FindAllString(s.String(), -1)
	multiplyInstructionIndices := r.FindAllStringIndex(s.String(), -1)

	for i := 0; i < len(multiplyInstructions); i++ {
		nums := r2.FindAllString(multiplyInstructions[i], -1)
		num0, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Printf("strconv: %s", err)
			os.Exit(1)
		}
		num1, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Printf("strconv: %s", err)
			os.Exit(1)
		}
		i := Instruction{
			Name:       MULTIPLY,
			StartIndex: multiplyInstructionIndices[i][0],
			EndIndex:   multiplyInstructionIndices[i][1],
			Param1:     num0,
			Param2:     num1,
		}
		completeInstructionList = append(completeInstructionList, i)
	}
	//var totalSum = 0

	//var machineOn = true

	//_ = SumAllMatches(r2, matches)

	for _, instruction := range completeInstructionList {
		fmt.Printf("%#v\n", instruction)
	}

	sort.Slice(completeInstructionList, func(a, b int) bool {
		return completeInstructionList[a].StartIndex < completeInstructionList[b].StartIndex
	})

	for _, instruction := range completeInstructionList {
		fmt.Printf("%s\n", instruction)
	}

	m := NewMachine()
	m.InitializeProgram(completeInstructionList)
	m.RunProgram()
	m.OutputResult()
}

func SumAllMatches(r2 *regexp.Regexp, matches []string) int {
	var totalSum = 0
	for ind, match := range matches {
		fmt.Printf("%d: %s", ind, match)
		fmt.Print("\t\t")
		terms := r2.FindAllString(match, -1)
		int1, err := strconv.Atoi(terms[0])
		if err != nil {
			log.Printf("error parsing int from regexp match: %s", err)
			os.Exit(1)
		}
		int2, err := strconv.Atoi(terms[1])
		if err != nil {
			log.Printf("error parsing int from regexp match: %s", err)
			os.Exit(1)
		}

		product := int1 * int2
		fmt.Printf("%d * %d = %d\n", int1, int2, product)
		totalSum += product
	}

	log.Printf("Total Matches: %d", len(matches))
	log.Printf("Total Sum: %d", totalSum)

	return totalSum
}
