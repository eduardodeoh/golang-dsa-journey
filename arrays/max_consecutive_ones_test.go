package arrays

import (
	"reflect"
	"testing"
)

var tests = []struct {
	arr      []int
	expected int
}{
	{[]int{1, 0, 1, 1, 1}, 3},
	{[]int{0, 1, 0, 1, 1}, 2},
	{[]int{1, 0, 1, 1, 0, 1}, 2},
}

func TestMaxConsecutiveOnes(t *testing.T) {
	for i, test := range tests {
		if got := MaxConsecutiveOnes(test.arr); !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.expected, got)
		}

		if got := MaxConsecutiveOnesAlternative(test.arr); !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.expected, got)
		}
	}
}

func TestMaxConsecutiveOnesAlternative(t *testing.T) {
	for i, test := range tests {
		if got := MaxConsecutiveOnesAlternative(test.arr); !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.expected, got)
		}
	}
}
