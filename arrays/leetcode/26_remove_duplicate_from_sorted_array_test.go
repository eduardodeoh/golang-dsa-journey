package arrays

import (
	"testing"
)

type removeDuplicateTableTests struct {
	nums         []int
	arr_expected []int
	num_expected int
}

func buildRemoveDuplicateTableTests() []removeDuplicateTableTests {
	return []removeDuplicateTableTests{
		{[]int{0, 0, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5},
		{[]int{1, 1, 2}, []int{1, 2}, 2},
	}
}

func TestRemoveDuplicateFromSortedArrayAlt1(t *testing.T) {
	for i, test := range buildRemoveDuplicateTableTests() {
		if got := RemoveDuplicateFromSortedArrayAlt1(test.nums); got != test.num_expected {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.num_expected, got)

		}
		for i := 0; i < len(test.arr_expected); i++ {
			if test.nums[i] != test.arr_expected[i] {
				t.Fatalf("Failed test case #%d. Want %v got %v", i, test.arr_expected, test.nums)
			}
		}
	}
}
