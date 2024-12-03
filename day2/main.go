package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	daily "github.com/uncleBlobby/aoc2024/internal"
)

func main() {
	fmt.Println("hello, aoc2024 day 2!")

	input := daily.GetInputFromFile("input")

	defer input.Close()

	reports := []Report{}

	var currentLine = 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		len := len(parts)
		reports = append(reports, Report{
			Safe:   false,
			Levels: make([]int, len),
		})
		for ind, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				log.Printf("error parsing int: %s", err)
				os.Exit(1)
			}
			reports[currentLine].Levels[ind] = num
		}
		currentLine += 1
	}

	for i := 0; i < len(reports); i++ {
		reports[i].SetSafety()
	}

	for ind, report := range reports {
		//	report.SetSafety()
		fmt.Printf("%d: %v\t%v\n", ind, report.Levels, report.Safe)
	}

	fmt.Printf("total safe reports: %d\n", CountSafeReports(reports))

	unsafeReports := GetUnsafeReports(reports)

	for i := 0; i < len(unsafeReports); i++ {
		unsafeReports[i].CheckProblemDampener()
	}

	fmt.Printf("total safe reports: %d\n", CountSafeReports(reports)+CountSafeReports(unsafeReports))
}
