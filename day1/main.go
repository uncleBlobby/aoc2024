package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hello, aoc day 1!")

	input, err := os.Open("input")
	if err != nil {
		log.Printf("os.Open: %s", err)
		os.Exit(1)
	}

	var firstSet []int
	var secondSet []int

	var differences []int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")

		int1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Printf("error parsing int: %s", err)
			os.Exit(1)
		}
		firstSet = append(firstSet, int1)

		int2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Printf("error parsing int: %s", err)
			os.Exit(1)
		}
		secondSet = append(secondSet, int2)

		//fmt.Printf("p1: %s p2: %s\n", parts[0], parts[1])

	}

	slices.Sort(firstSet)
	slices.Sort(secondSet)

	log.Printf("[%d]: [%d]", len(firstSet), len(secondSet))

	for ind, num := range firstSet {
		//fmt.Println(num)
		diff := num - secondSet[ind]
		if diff < 0 {
			diff = diff * -1
		}

		differences = append(differences, diff)
	}

	totalDistance := 0

	for _, num := range differences {
		totalDistance += num
	}

	fmt.Printf("total distance between sets: %d\n", totalDistance)

	input.Close()

	firstMap := map[int]int{}

	for _, num := range firstSet {
		firstMap[num] = 0
		for _, numb := range secondSet {
			if numb == num {
				firstMap[num] += 1
			}
			if numb > num {
				continue
			}
		}
	}

	var score = 0

	for key, count := range firstMap {
		score += (key * count)
	}

	log.Printf("total similarity score: %d", score)
}
