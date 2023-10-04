package arrays

import (
	"sort"
)

/*
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

Example 1:

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].

Example 2:

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]

Constraints:

    1 <= nums.length <= 104
    -104 <= nums[i] <= 104
    nums is sorted in non-decreasing order.


Source: https://leetcode.com/problems/squares-of-a-sorted-array (#977)
*/

/*
Complexity Analysis

	Time complexity: O(n log n)
	Space complexity: O(1)
*/
func SortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] *= nums[i]
	}
	sort.Ints(nums)

	return nums
}

/*
Complexity Analysis

	Time complexity: O(n)
	Space complexity: O(n)
	Auxiliary complexity: O(1)
*/
func SortedSquaresApproachTwoPointers(nums []int) []int {
	result := make([]int, len(nums))
	arraySize := len(nums) - 1
	headIndex, tailIndex := 0, arraySize

	for i := tailIndex; i >= 0; i-- {
		if abs(nums[headIndex]) > abs(nums[tailIndex]) {
			result[i] = nums[headIndex] * nums[headIndex]
			headIndex++
		} else {
			result[i] = nums[tailIndex] * nums[tailIndex]
			tailIndex--
		}
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
