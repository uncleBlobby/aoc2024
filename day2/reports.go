package main

import "slices"

type Report struct {
	Safe   bool
	Levels []int
}

type ReportSet = []Report

func CountSafeReports(reportSet []Report) int {
	var count = 0
	for _, report := range reportSet {
		if report.Safe {
			count += 1
		}
	}

	return count
}

func GetUnsafeReports(reportSet []Report) []Report {
	unsafe := []Report{}

	for _, r := range reportSet {
		if !r.Safe {
			unsafe = append(unsafe, r)
		}
	}

	return unsafe
}

// A report is safe if and only if:
//	1. all levels in the report data are either all increasing or all decreasing
//  2. any two adjacent levels in the report differ by at least one and at most three

func (r *Report) SetSafety() {
	var isSorted = false

	// if slice is already sorted ascending
	if slices.IsSorted(r.Levels) {
		isSorted = true
	} else {
		// if slice is not sorted ascending, reverse and check if sorted ascending
		// (same as checking if sorted descending but less efficient)
		slices.Reverse(r.Levels)
		if slices.IsSorted(r.Levels) {
			isSorted = true
		} else {
			// if slice is neither sorted ascending nor descending, report cannot be safe
			r.Safe = false
			return
		}
	}

	// if slice is sorted, check the difference between levels
	if isSorted {
		r.Safe = reportHasSafeLevels(r.Levels)
	}
	return
}

func reportIsSorted(report []int) bool {
	if slices.IsSorted(report) {
		return true
	} else {
		slices.Reverse(report)
		if slices.IsSorted(report) {
			return true
		}
	}
	return false
}

func (r *Report) CheckProblemDampener() {
	// if checkIfReportIsSortedWithOneDeletion(r.Levels) {
	// 	if reportHasSafeLevels(r.Levels) {
	// 		r.Safe = true
	// 	}
	// }

	if checkIfReportIsSafeWithOneDeletion(r.Levels) {
		r.Safe = true
	}
}

func checkIfReportIsSafeWithOneDeletion(report []int) bool {

	dup := make([]int, len(report))
	withOneDeleted := []int{}

	_ = copy(dup, report)

	for i := 0; i < len(dup); i++ {
		withOneDeleted = deleteIndex(dup, i)
		if reportIsSorted(withOneDeleted) {
			if reportHasSafeLevels(withOneDeleted) {
				return true
			}
		}
	}

	return false
	// withOneDeleted := []int{}
	// for i := 0; i < len(report)-1; i++ {
	// 	if report[i+1] <= report[i] {
	// 		withOneDeleted = deleteIndex(report, i+1)
	// 		return reportIsSorted(withOneDeleted)
	// 	}
	// }
	// return false
}

func deleteIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func reportHasSafeLevels(report []int) bool {
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
