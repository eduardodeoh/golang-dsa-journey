package arrays

/*
You are given an array 'arr' containing 'n' integers. You are also given an integer 'num', and your
task is to find whether 'num' is present in the array or not. If 'num' is present in the array,
return 0-based index of the first ocurrence of 'num'. Else, return -1.

Source:
https://leetcode.com/problems/max-consecutive-ones/description/

Complexity Analysis
    Time complexity: O(n), where n is the length of the input array nums. We iterate through the array once, and the time complexity is linear with respect to the size of the input.

    Space complexity: O(1), as we use only a constant amount of extra space to store maxCount and currentCount, regardless of the size of the input array.

*/

func MaxConsecutiveOnes(nums []int) int {
	maxCount, currentCount := 0, 0

	for _, item := range nums {
		if item == 1 {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 0
		}
	}
	return maxCount
}

func MaxConsecutiveOnesAlternative(nums []int) int {
	maxCount, currentCount := 0, 0

	for _, item := range nums {
		currentCount = currentCount*item + item
		if currentCount > maxCount {
			maxCount = currentCount
		}
	}
	return maxCount
}
