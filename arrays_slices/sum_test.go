package arrays_slices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Array of five integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		expected := 15
		result := Sum(numbers)

		if expected != result {
			t.Errorf("Got %d want %d given %v", result, expected, numbers)
		}
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2}

		expected := 3
		result := Sum(numbers)

		if expected != result {
			t.Errorf("Got %d want %d given %v", result, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	expected := []int{3, 9}
	result := SumAll([]int{1, 2}, []int{0, 9})

	if !slices.Equal(expected, result) {
		t.Errorf("Got %v want %v", result, expected)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, expected, result []int) {
		t.Helper()
		if !slices.Equal(expected, result) {
			t.Errorf("Got %v want %v", result, expected)
		}
	}

	t.Run("Make sums of some slices", func(t *testing.T) {
		expected := []int{2, 9}
		result := SumAllTails([]int{1, 2}, []int{0, 9})

		checkSums(t, expected, result)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		result := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		checkSums(t, expected, result)
	})
}
