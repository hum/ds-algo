package algorithms_test

import (
  "testing"

  algo "github.com/hum/ds-algo/algorithms"
)

func TestSelectionSort(t *testing.T) {
  values := []int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0}
	expected := []int{0, 1, 2, 4, 5, 6, 44, 63, 87, 99, 283}

	algo.SelectionSort(values)

	if !eq(values, expected) {
		t.Fatalf("Slice is not sorted in the expected order. Expected=%v, got=%v", expected, values)
	}

	values = []int{1}
	expected = []int{1}
	algo.SelectionSort(values)

	if !eq(values, expected) {
		t.Fatalf("Slice has been modified. Expected=%v, got=%v", expected, values)
	}

	values = []int{1, 2}
	expected = []int{1, 2}
	algo.SelectionSort(values)

	if !eq(values, expected) {
		t.Fatalf("Slice has been modified. Expected=%v, got=%v", expected, values)
	}

	values = []int{2, 1}
	expected = []int{1, 2}
	algo.SelectionSort(values)

	if !eq(values, expected) {
		t.Fatalf("Slice of the size %d has not been sorted properly. Expected=%v, got=%v", len(values), expected, values)
	}

	values = []int{-100, 99, 85, -1, 0, 15}
	expected = []int{-100, -1, 0, 15, 85, 99}
	algo.SelectionSort(values)

	if !eq(values, expected) {
		t.Fatalf("Slice is not sorted in the expected order. Expected=%v, got=%v", expected, values)
	}
}
