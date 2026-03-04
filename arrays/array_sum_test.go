package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	actual := Sum(numbers)
	expected := 15

	if actual != expected {
		t.Errorf("expected %d but actual is %d", expected, actual)
	}
}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{1, 2, 3, 4, 5}, []int{0, 9})
	expected := []int{15, 9}

	if !slices.Equal(actual, expected) {
		t.Errorf("expected %v but actual is %v", expected, actual)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, actual, expected []int) {
		t.Helper()
		if !slices.Equal(actual, expected) {
			t.Errorf("got %v want %v", actual, expected)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checkSums(t, actual, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		checkSums(t, actual, expected)
	})

	t.Run("safely sum slice of length 1", func(t *testing.T) {
		actual := SumAllTails([]int{1}, []int{2})
		expected := []int{0, 0}
		checkSums(t, actual, expected)
	})
}
