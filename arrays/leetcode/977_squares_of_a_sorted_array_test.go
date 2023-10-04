package arrays

import (
	"reflect"
	"testing"
)

type squaresOfSortedArrayTableTests struct {
	arr      []int
	expected []int
}

func buildSquaresOfSortedArrayTableTests() []squaresOfSortedArrayTableTests {
	return []squaresOfSortedArrayTableTests{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-4, -2, -1}, []int{1, 4, 16}},
		{[]int{1, 2, 4}, []int{1, 4, 16}},
	}
}

func TestSortedSquares(t *testing.T) {
	for i, test := range buildSquaresOfSortedArrayTableTests() {
		if got := SortedSquares(test.arr); !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.expected, got)
		}
	}
}

func TestSortedSquaresApproachTwoPointers(t *testing.T) {
	for i, test := range buildSquaresOfSortedArrayTableTests() {
		if got := SortedSquaresApproachTwoPointers(test.arr); !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.expected, got)
		}
	}
}
