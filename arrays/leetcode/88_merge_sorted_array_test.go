package arrays

import (
	"reflect"
	"testing"
)

type mergeSortedArraysTableTests struct {
	nums1    []int
	m        int
	nums2    []int
	n        int
	expected []int
}

func buildMergeSortedArraysTableTests() []mergeSortedArraysTableTests {
	return []mergeSortedArraysTableTests{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{2, 2, 3, 0, 0, 0}, 3, []int{1, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}
}

func TestMergeSortedArraysAlt1(t *testing.T) {
	for i, test := range buildMergeSortedArraysTableTests() {
		if got := MergeSortedArraysAlt1(test.nums1, test.m, test.nums2, test.n); !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.expected, got)
		}
	}
}

func TestMergeSortedArraysAlt2(t *testing.T) {
	for i, test := range buildMergeSortedArraysTableTests() {
		if got := MergeSortedArraysAlt2(test.nums1, test.m, test.nums2, test.n); !reflect.DeepEqual(got, test.expected) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.expected, got)
		}
	}
}

func TestMergeSortedArraysAlt3(t *testing.T) {
	for i, test := range buildMergeSortedArraysTableTests() {
		MergeSortedArraysAlt3(test.nums1, test.m, test.nums2, test.n)
		if !reflect.DeepEqual(test.nums1, test.expected) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.expected, test.nums1)
		}
	}
}
