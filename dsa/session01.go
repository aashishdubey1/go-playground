package dsa

// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
// Return the running sum of nums.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [1,3,6,10]
// Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].


func FindRunningSum(nums []int) []int {

	runningSum := []int{}
	sum := 0

	for _ , val := range nums {
		sum += val
		runningSum = append(runningSum,sum)
	}

	return runningSum

}