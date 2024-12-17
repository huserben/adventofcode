package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Operation struct {
	Match    string
	Location [2]int
}

func ParseFile(inputFileName string) string {
	fileAsBytes, _ := os.ReadFile(inputFileName)
	return string(fileAsBytes)
}

func ExtractUncorruptedOperations(memory string) []Operation {
	// Regular expression to match mul(x,y) where x and y are 1-3 digits
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	matches := re.FindAllString(memory, -1)
	locations := re.FindAllStringIndex(memory, -1)

	// Combine matches and locations into a slice of Operation
	var operations []Operation
	for i := range matches {
		operations = append(operations, Operation{
			Match:    matches[i],
			Location: [2]int{locations[i][0], locations[i][1]},
		})
	}

	return operations
}

func ExtractEnableDisableOperations(memory string) []Operation {
	// Regular expression to match "do()" and "don't()"
	re := regexp.MustCompile(`do\(\)|don't\(\)`)

	matches := re.FindAllString(memory, -1)
	locations := re.FindAllStringIndex(memory, -1)

	// Combine matches and locations into a slice of Operation
	var operations []Operation
	for i := range matches {
		operations = append(operations, Operation{
			Match:    matches[i],
			Location: [2]int{locations[i][0], locations[i][1]},
		})
	}

	return operations
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

func FilterOperationsByState(mulOps []Operation, enableDisableOps []Operation) []Operation {
	var filtered []Operation
	enable := true

	for _, op := range mulOps {
		for _, edOp := range enableDisableOps {
			if edOp.Location[1] <= op.Location[0] {
				if edOp.Match == "do()" {
					enable = true
				} else if edOp.Match == "don't()" {
					enable = false
				}
			} else {
				break
			}
		}

		if enable {
			filtered = append(filtered, op)
		}
	}

	return filtered
}

func GetMatches(operations []Operation) []string {
	var matches []string
	for _, op := range operations {
		matches = append(matches, op.Match)
	}
	return matches
}

func main() {
	fmt.Println("Parsing Input File...")
	memory := ParseFile("input.txt")

	fmt.Println("Extracting Uncorrupted Operations")
	uncorruptedOperations := ExtractUncorruptedOperations(memory)

	fmt.Printf("Found %d uncorrupted operations - multiplying and summing them up", len(uncorruptedOperations))
	fmt.Println()

	uncorruptedOperationMatches := GetMatches(uncorruptedOperations)

	sumOfOperations := MultiplyAndSumOperations(uncorruptedOperationMatches)
	fmt.Printf("Sum of all valid operations: %d", sumOfOperations)
	fmt.Println()

	fmt.Println("Extracting Enabled and Disabled Operations")
	enabledDisabledOperations := ExtractEnableDisableOperations(memory)

	fmt.Println("Filtering Operations Based on Enable/Disable State")
	filteredOperations := FilterOperationsByState(uncorruptedOperations, enabledDisabledOperations)

	fmt.Printf("Found %d filtered operations - multiplying and summing them up", len(filteredOperations))
	fmt.Println()

	filteredMatches := GetMatches(filteredOperations)
	sumOfOperations = MultiplyAndSumOperations(filteredMatches)
	fmt.Printf("Sum of all valid operations that are enabled: %d", sumOfOperations)
}
