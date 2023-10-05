package arrays

import (
	"sort"
	"testing"
)

type removeElementTableTests struct {
	nums         []int
	val          int
	arr_expected []int
	val_expected int
}

func buildRemoveElementTableTests() []removeElementTableTests {
	return []removeElementTableTests{
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4}, 5},
		{[]int{3, 2, 2, 3}, 3, []int{2, 2}, 2},
	}
}

func TestRemoveElement(t *testing.T) {
	for i, test := range buildRemoveElementTableTests() {
		if got := RemoveElementAlt1(test.nums, test.val); got != test.val_expected {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.val_expected, got)

		}
		tempNums := test.nums[:test.val_expected]
		sort.Ints(tempNums)
		sort.Ints(test.arr_expected)
		for i := 0; i < len(test.arr_expected); i++ {
			if tempNums[i] != test.arr_expected[i] {
				t.Fatalf("Failed test case #%d. Want %v got %v", i, test.arr_expected, tempNums)
			}
		}
	}
}

func TestRemoveElementAlt2(t *testing.T) {
	for i, test := range buildRemoveElementTableTests() {
		if got := RemoveElementAlt2(test.nums, test.val); got != test.val_expected {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.val_expected, got)

		}
		tempNums := test.nums[:test.val_expected]
		sort.Ints(tempNums)
		sort.Ints(test.arr_expected)
		for i := 0; i < len(test.arr_expected); i++ {
			if tempNums[i] != test.arr_expected[i] {
				t.Fatalf("Failed test case #%d. Want %v got %v", i, test.arr_expected, tempNums)
			}
		}
	}
}
