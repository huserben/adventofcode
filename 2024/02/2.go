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

func IsSafe(report []int) bool {
	isIncreasing := false
	isDecreasing := false

	for index := 0; index < len(report)-1; index++ {
		level1 := report[index]
		level2 := report[index+1]

		isIncreasing = isIncreasing || level1 < level2
		isDecreasing = isDecreasing || level1 > level2

		diff := Abs(level1 - level2)
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return isIncreasing != isDecreasing
}

func IdentifySafeReports(reports [][]int) int {
	safeReportCounter := 0

	for _, report := range reports {
		isSafe := IsSafe(report)

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

	safeReports := IdentifySafeReports(reports)
	fmt.Printf("Found %d safe reports!", safeReports)
}
