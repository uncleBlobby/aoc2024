package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	daily "github.com/uncleBlobby/aoc2024/internal"
)

type Collection struct {
	Sorted   [][]int
	Unsorted [][]int
}

func main() {
	fmt.Println("hello, aoc day 2!")

	reports := GetReports("testdata")

	//log.Printf("total # of sorted reports: %d", len(sortedReports))

	// Filter reports for only those which are already sorted
	collection := FilterSortedReports(reports)

	log.Printf("unsorted reports length: %d", len(collection.Unsorted))
	// Filter reports for only those in which adjacent levels differ by >= 1 and <= 3

	validReports := FilterValidReports(collection.Sorted)

	log.Printf("%d of the reports are valid", len(validReports))

	result := FilterUnsortedCollectionWithProblemDampener(collection.Unsorted)

	log.Printf("%d of the reports are valid (with problem dampener)", len(result.Sorted))

	var newValidReportsCount = 0

	for _, report := range result.Sorted {
		if ReportIsSafe(report) {
			newValidReportsCount += 1
		}
	}

	// for _, report := range collection.Sorted {
	// 	if ReportIsSafeWithOneDeletion(report) {
	// 		newValidReportsCount += 1
	// 	}
	// }

	log.Printf("%d of the reports are valid with problem dampener", len(validReports)+newValidReportsCount)
}

// Read all input data into 2D array of int

func GetReports(inputFileName string) [][]int {
	input := daily.GetInputFromFile(inputFileName)

	defer input.Close()

	var reports [][]int

	var currentLine = 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		len := len(parts)
		reports = append(reports, make([]int, len))
		for ind, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				log.Printf("error parsing int: %s", err)
				os.Exit(1)
			}
			reports[currentLine][ind] = num
		}
		currentLine += 1
	}

	log.Printf("total # of reports: %d", len(reports))
	return reports
}

func FilterSortedReports(startData [][]int) Collection {

	var sortedReports [][]int
	var sortedCount = 0

	var unsortedReports [][]int
	var unsortedCount = 0

	for _, report := range startData {

		if slices.IsSorted(report) {
			sortedReports = append(sortedReports, make([]int, len(report)))
			sortedReports[sortedCount] = report
			sortedCount += 1
		} else {
			slices.Reverse(report)

			if slices.IsSorted(report) {
				//log.Printf("report #%d [%v] is sorted", ind, report)
				sortedReports = append(sortedReports, make([]int, len(report)))
				sortedReports[sortedCount] = report
				sortedCount += 1
			} else {
				unsortedReports = append(unsortedReports, make([]int, len(report)))
				unsortedReports[unsortedCount] = report
				unsortedCount += 1
			}
		}
	}
	// log.Printf("total # of sorted reports: %d", len(sortedReports))
	// log.Printf("sortedCount: %d", sortedCount)
	output := make([][][]int, 2)
	output = append(output, sortedReports)
	output = append(output, unsortedReports)

	c := Collection{
		Sorted:   sortedReports,
		Unsorted: unsortedReports,
	}

	return c
}

func FilterValidReports(sortedReports [][]int) [][]int {
	var validReports [][]int
	//var invalidCount = 0
	var validCount = 0

	for i := 0; i < len(sortedReports); i++ {

		if ReportIsSafe(sortedReports[i]) {
			validReports = append(validReports, sortedReports[i])
			validCount += 1
			//log.Printf("%v : Safe", sortedReports[i])
		} else {
			if ReportIsSafeWithOneDeletion(sortedReports[i]) {
				validReports = append(validReports, sortedReports[i])
				validCount += 1
			}
		}
	}

	return validReports
}

func ReportIsSafeWithOneDeletion(report []int) bool {
	log.Printf("report: %v", report)
	var reportIsSafe = true
	var maxDiff = 0
	var maxDiffIndex = 0
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if diff < 0 {
			diff = diff * -1
		}
		if diff < maxDiff {
			maxDiff = diff
			maxDiffIndex = i + 1
		}
		if diff < 1 || diff > 3 {
			reportIsSafe = false
		}

		if !reportIsSafe && i+1 == len(report) {
			arrayWithOneDeleted := DeleteIndex(report, maxDiffIndex)
			log.Printf("arrayWithOneDeleted: %v", arrayWithOneDeleted)
			if ReportIsSafe(arrayWithOneDeleted) {
				return true
			}
		}
	}
	return false
}

func ReportIsSafe(report []int) bool {
	var reportIsSafe = true
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if diff < 0 {
			diff = diff * -1
		}
		if diff < 1 || diff > 3 {
			reportIsSafe = false
		}

		if reportIsSafe && i+1 == len(report) {
			return reportIsSafe
		}
	}

	return reportIsSafe
}

func CheckIfArrayIsSortableWithOneDeletion(a []int) bool {
	var arrayWithOneDeleted []int
	for i := 0; i < len(a)-1; i++ {
		if a[i+1] <= a[i] {
			//log.Printf("array %v is not sorted", a)
			arrayWithOneDeleted = DeleteIndex(a, i+1)
			if slices.IsSorted(arrayWithOneDeleted) {
				return true
			}
		}
	}

	return false
}

func DeleteIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func FilterUnsortedCollectionWithProblemDampener(u [][]int) Collection {
	output := Collection{}
	for _, report := range u {
		if CheckIfArrayIsSortableWithOneDeletion(report) {
			output.Sorted = append(output.Sorted, report)
		}
	}

	log.Printf("%d reports were restored with problem dampener", len(output.Sorted))

	//output.Sorted = append(output.Sorted, c.Sorted...)
	return output
}
