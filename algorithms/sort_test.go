package algorithms_test

import (
	"fmt"
	"testing"

	algo "github.com/hum/ds-algo/algorithms"
)

var tests = []struct {
	input []int
	expected []int
} {
	{[]int{7, 5, 4, 3, 6, 8, 1, 2, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	{[]int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}, []int{0, 1, 2, 4, 5, 6, 44, 63, 87, 99, 283}},
	{[]int{1}, []int{1}},
	{[]int{1, 2}, []int{1, 2}},
	{[]int{2, 1}, []int{1, 2}},
	{[]int{-100, 99, 85, -1, 0, 15}, []int{-100, -1, 0, 15, 85, 99}},
}

func testSort(t *testing.T, sortFunc func([]int) []int) error {
	for _, tt := range tests {
		result := sortFunc(tt.input)
		if !eq(tt.expected, result) {
			return fmt.Errorf("Got=%v, expected=%v", result, tt.expected)
		}
	}
	return nil
}

func TestQuickSort(t *testing.T) {
	for _, tt := range tests {
		result := algo.QuickSort(tt.input, 0, len(tt.input)-1)
		if !eq(tt.expected, result) {
			t.Fatalf("Quick Sort failed. Got=%v, expected=%v", result, tt.expected)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	if err := testSort(t, algo.BubbleSort); err != nil {
		t.Fatalf("Bubble Sort failed. %v", err)
	}
}

func TestInsertionSort(t *testing.T) {
	if err := testSort(t, algo.InsertionSort); err != nil {
		t.Fatalf("Insertion Sort failed. %v", err)
	}
}

func TestMergeSort(t *testing.T) {
	if err := testSort(t, algo.MergeSort); err != nil {
		t.Fatalf("Merge Sort failed. %v", err)
	}
}

func TestSelectionSort(t *testing.T) {
	if err := testSort(t, algo.SelectionSort); err != nil {
		t.Fatalf("Selection Sort failed. %v", err)
	}
}

func eq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
