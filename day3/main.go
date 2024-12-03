package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	doIndexes := doInst.FindAllStringIndex(s.String(), -1)

	for _, do := range doIndexes {
		log.Printf("do matches: %v", do)
	}

	dontIndexes := donotInst.FindAllStringIndex(s.String(), -1)

	for _, dont := range dontIndexes {
		log.Printf("dont matches: %v", dont)
	}

	//dos := doInst.FindStringIndex()

	matches := r.FindAllString(s.String(), -1)

	//var totalSum = 0

	//var machineOn = true

	_ = SumAllMatches(r2, matches)
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
