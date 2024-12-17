package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseFile(inputFileName string) string {
	fileAsBytes, _ := os.ReadFile(inputFileName)
	return string(fileAsBytes)
}

func ExtractUncorruptedOperations(memory string) []string {
	// Regular expression to match mul(x,y) where x and y are 1-3 digits
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	matches := re.FindAllString(memory, -1)

	return matches
}

func MultiplyAndSumOperations(validOperations []string) int {
	sumOfMultiplications := 0

	for _, op := range validOperations {
		splitOperation := strings.Split(op, ",")
		factor1, _ := strconv.Atoi(strings.Split(splitOperation[0], "(")[1])
		factor2, _ := strconv.Atoi(strings.Split(splitOperation[1], ")")[0])

		sumOfMultiplications += (factor1 * factor2)
	}

	return sumOfMultiplications
}

func main() {
	fmt.Println("Parsing Input File...")
	memory := ParseFile("input.txt")

	fmt.Println("Extracting Uncorrupted Operations")
	uncorruptedOperations := ExtractUncorruptedOperations(memory)

	fmt.Printf("Found %d uncorrupted operations - multiplying and summing them up", len(uncorruptedOperations))
	fmt.Println()

	sumOfOperations := MultiplyAndSumOperations(uncorruptedOperations)
	fmt.Printf("Sum of all valid operations: %d", sumOfOperations)
}
