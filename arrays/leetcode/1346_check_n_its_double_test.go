package arrays

import (
	"testing"
)

type checkNAndItsDoubleTableTests struct {
	nums     []int
	expected bool
}

func buildCheckNAndItsDoubleTableTests() []checkNAndItsDoubleTableTests {
	return []checkNAndItsDoubleTableTests{
		{[]int{10, 2, 5, 3}, true},
		{[]int{3, 1, 7, 11}, false},
		{[]int{5, 2, 10, 3}, true},
	}
}

func TestCheckNAndItsDouble(t *testing.T) {
	for i, test := range buildCheckNAndItsDoubleTableTests() {
		if got := CheckNAndItsDoubleAlt1(test.nums); got != test.expected {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.expected, got)

		}
	}
}
