package arrays

/*
You are given an array 'arr' containing 'n' integers. You are also given an integer 'num', and your
task is to find whether 'num' is present in the array or not. If 'num' is present in the array,
return 0-based index of the first ocurrence of 'num'. Else, return -1.

Source:
https://www.codingninjas.com/studio/guided-paths/data-structures-algorithms/content/607713/offering/9536703?leftPanelTab=0

Complexity Analysis
	Time Complexity: 	O(N) - N is the size of the array - For the worst case, we are going to iterate all the elements of the array once

   Space complexity: O(1), as we use only a constant amount of extra space to store index and i, regardless of the size of the input array.
*/

func LinearSearch(n int, num int, arr []int) int {
	index := -1

	for i := 0; i < n; i++ {
		if arr[i] == num {
			index = i
			break
		}
	}
	return index
}
