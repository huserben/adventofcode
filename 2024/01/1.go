package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

func Sort(arr []int) []int {
	sort.Ints(arr)
	return arr
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}

	return val
}

func Diff(input1 []int, input2 []int) []int {
	diffSlice := []int{}

	for index := 0; index < len(input1); index++ {
		num1 := input1[index]
		num2 := input2[index]

		diff := Abs((num1 - num2))
		diffSlice = append(diffSlice, diff)
	}

	return diffSlice
}

func Sum(input []int) int {
	sum := 0

	for index := 0; index < len(input); index++ {
		val := input[index]
		sum += val
	}

	return sum
}

func ParseInputFile(fileName string) ([]int, []int) {
	column1 := []int{}
	column2 := []int{}

	file, err := os.Open(fileName)
	if err != nil {
		return column1, column2
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line) // Split line by whitespace
		if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				column1 = append(column1, num1)
				column2 = append(column2, num2)
			}
		}
	}

	return column1, column2
}

func CountNumber(input []int, number int) int {
	count := 0

	for _, num := range input {
		if num == number {
			count++
		}
	}

	return count
}

func CreateSimilarityScore(input1 []int, input2 []int) int {
	similarityScore := 0

	for _, number := range input1 {
		muliplier := CountNumber(input2, number)

		similarityScore += number * muliplier
	}

	return similarityScore
}

func main() {
	list1, list2 := ParseInputFile("input.txt")

	sortedList1 := Sort(list1)
	sortedList2 := Sort(list2)

	diff := Diff(sortedList1, sortedList2)

	sum := Sum(diff)

	fmt.Fprintln(os.Stdout, []any{"", sum}...)

	fmt.Fprintln(os.Stdout, []any{"Creating Similarity Score"}...)
	similarityScore := CreateSimilarityScore(list1, list2)
	fmt.Fprintln(os.Stdout, []any{"Similarity Score is ", similarityScore}...)
}
