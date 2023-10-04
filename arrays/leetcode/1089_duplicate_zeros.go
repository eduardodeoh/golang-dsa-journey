package arrays

import (
	"container/list"
)

/*
Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written. Do the above modifications to the input array in place and do not return anything.



Example 1:

Input: arr = [1,0,2,3,0,4,5,0]
Output: [1,0,0,2,3,0,0,4]
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

Example 2:

Input: arr = [1,2,3]
Output: [1,2,3]
Explanation: After calling your function, the input array is modified to: [1,2,3]



Constraints:

    1 <= arr.length <= 104
    0 <= arr[i] <= 9

Source: https://leetcode.com/problems/duplicate-zeros/description/
*/

/*
Complexity Analysis

	Time Complexity: O(N^2)
	Space complexity: O(1)
*/
func DuplicateZerosAlt1(arr []int) {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
		}
	}
}

/*
Complexity Analysis

	Time Complexity: O(N)
	Space complexity: O(N)
*/
func DuplicateZerosAlt2(arr []int) {
	queue := list.New()
	for i, item := range arr {
		if item == 0 {
			queue.PushBack(0)
			queue.PushBack(0)
		} else {
			queue.PushBack(item)
		}
		arr[i] = queue.Front().Value.(int)
		queue.Remove(queue.Front())

	}
}
