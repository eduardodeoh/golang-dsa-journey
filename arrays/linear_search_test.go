package arrays

import (
	"reflect"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		n        int
		num      int
		arr      []int
		expected int
	}{
		{5, 4, []int{6, 7, 8, 4, 1}, 3},
		{4, 2, []int{2, 5, 6, 2}, 0},
		{16, 19, []int{9, 23, 1, 15, 22, 13, 24, 25, 5, 17, 8, 14, 18, 20, 7, 2}, -1},
	}
	for i, test := range tests {
		if got := LinearSearch(test.n, test.num, test.arr); !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.expected, got)
		}
	}
}
