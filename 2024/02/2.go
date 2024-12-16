package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInputFile(fileName string) [][]int {
	reports := [][]int{}

	file, err := os.Open(fileName)
	if err != nil {
		return reports
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		report := []int{}

		for _, value := range parts {
			level, _ := strconv.Atoi(value)

			report = append(report, level)
		}

		reports = append(reports, report)
	}

	return reports
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}

	return val
}

func CheckSafetyOfTwoLevels(level1 int, level2 int, reportIsIncreasing bool) bool {
	isIncreasing := level1 < level2

	diff := Abs(level1 - level2)
	if diff < 1 || diff > 3 || isIncreasing != reportIsIncreasing {
		return false
	}

	return true
}

func IsSafe(report []int) bool {
	isIncreasing := false

	if len(report) >= 2 {
		isIncreasing = report[0] < report[1]
	}

	for index := 0; index < len(report)-1; index++ {
		level1 := report[index]
		level2 := report[index+1]

		isSafe := CheckSafetyOfTwoLevels(level1, level2, isIncreasing)

		if !isSafe {
			return false
		}
	}

	return true
}

func IsSafeWithDampener(report []int) bool {
	// Loop through existing slice, and create new slices where one entry is removed and check if they are safe...

	for index := range report {
		dampenedReport := append([]int(nil), report[:index]...)
		dampenedReport = append(dampenedReport, report[index+1:]...)
		if IsSafe(dampenedReport) {
			return true
		}
	}

	return false
}

func IdentifySafeReports(reports [][]int, problemDampener bool) int {
	safeReportCounter := 0

	for _, report := range reports {
		isSafe := IsSafe(report)

		if !isSafe && problemDampener {
			isSafe = IsSafeWithDampener(report)
		}

		if isSafe {
			safeReportCounter++
		}
	}

	return safeReportCounter
}

func main() {
	fmt.Println("Parsing Input File...")
	reports := ParseInputFile("input.txt")

	fmt.Printf("Found %d reports...identifying safe reports", len(reports))
	fmt.Println()

	safeReports := IdentifySafeReports(reports, false)
	fmt.Printf("Found %d safe reports!", safeReports)
	fmt.Println()

	fmt.Println("Identifying safe reports with Problem Dampener enabled")
	safeReports = IdentifySafeReports(reports, true)
	fmt.Printf("Found %d safe reports!", safeReports)
}
