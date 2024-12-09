package main

import (
	"reflect"
	"testing"
)

func TestArraySort(t *testing.T) {
	numbers := []int{2, 3, 1, 5, 4}

	got := Sort(numbers)
	want := []int{1, 2, 3, 4, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		expected int
	}{
		{"Positive Value", 12, 12},
		{"Negative Value", -42, 42},
		{"Zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Abs(tt.value)

			if got != tt.expected {
				t.Errorf("Got %d expected %d", got, tt.expected)
			}
		})
	}
}

func TestArrayDiff(t *testing.T) {
	tests := []struct {
		name   string
		input1 []int
		input2 []int
		want   []int
	}{
		{"Identical Arrays", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, []int{0, 0, 0, 0, 0}},
		{"Different Arrays", []int{0, 0, 0, 0, 0}, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Different Arrays Reverse", []int{1, 2, 3, 4, 5}, []int{0, 0, 0, 0, 0}, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Diff(tt.input1, tt.input2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diff(%v, %v) = %v; want %v", tt.input1, tt.input2, got, tt.want)
			}
		})
	}
}

func TestParseInputFile(t *testing.T) {
	expectedInput1 := []int{12, 23, 13}
	expectedInput2 := []int{27, 42, 37}

	got1, got2 := ParseInputFile("testinput.txt")

	if !reflect.DeepEqual(got1, expectedInput1) {
		t.Errorf("got %v want %v", got1, expectedInput1)
	}

	if !reflect.DeepEqual(got2, expectedInput2) {
		t.Errorf("got %v want %v", got2, expectedInput2)
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Positive Numbers", []int{1, 2, 3, 4, 5}, 15},
		{"Zeros", []int{0, 0, 0, 0, 0}, 0},
		{"Negative Numbers", []int{-1, -2, -3, -4, -5}, -15},
		{"Mixed Numbers", []int{-1, -2, -3, -4, 15}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.input)
			if got != tt.expected {
				t.Errorf("Sum(%v) = %d; want %d", tt.input, got, tt.expected)
			}
		})
	}
}
