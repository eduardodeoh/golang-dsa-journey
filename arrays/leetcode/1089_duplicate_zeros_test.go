package arrays

import (
	"reflect"
	"testing"
)

func TestDuplicateZerosAlt1(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
	}
	for i, test := range tests {
		DuplicateZerosAlt1(test.arr)
		if !reflect.DeepEqual(test.arr, test.expected) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.expected, test.arr)
		}
	}
}

func TestDuplicateZerosAlt2(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
	}
	for i, test := range tests {
		DuplicateZerosAlt2(test.arr)
		if !reflect.DeepEqual(test.arr, test.expected) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.expected, test.arr)
		}
	}
}
