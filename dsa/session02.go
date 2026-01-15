package dsa

import (
	"math"
)

/// - Left and Right Sum Differences
/*
You are given a 0-indexed integer array nums of size n.
Define two arrays leftSum and rightSum where:
leftSum[i] is the sum of elements to the left of the index i in the array nums. If there is no such element, leftSum[i] = 0.
rightSum[i] is the sum of elements to the right of the index i in the array nums. If there is no such element, rightSum[i] = 0.
Return an integer array answer of size n where answer[i] = |leftSum[i] - rightSum[i]|.
*/

func LeftRightDifference(nums []int) []int {

	leftSum := []int{}
	rightSum := []int{}
  diff := []int{}
	sumL := 0


	if len(nums) == 1 {
		return []int{0}
	}

	for i,v := range nums  {

		if i == len(nums) - 1 {
			sumL = 0
			break
		}
		if i == 0 {
			leftSum = append(leftSum, 0)
		}
		sumL += v
		leftSum = append(leftSum, sumL)
	}
	for i := len(nums) - 1 ; i>=0; i--{
		if i == 0 {
			break
		}
		if i == len(nums) - 1 {
			rightSum = append([]int{0}, rightSum...)
		}
		sumL += nums[i]
		rightSum = append([]int{sumL}, rightSum...)
	}

	for i:=0 ; i<len(nums); i++ {
		diff = append(diff, int(math.Abs(float64(leftSum[i] - rightSum[i]))))
	}

	 return diff 

}


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


// You are given an integer array arr[]. The task is to find the sum of it.
// Input: arr[] = [1, 2, 3, 4]
// Output: 10
// Explanation: 1 + 2 + 3 + 4 = 10.

func FindSum(nums []int) int {
	sum := 0

	for _, val := range nums {
		sum += val
	}

	return sum
}

// Given an array of positive integers arr[], return the second largest element from the array. If the second largest element doesn't exist then return -1.
// Note: The second largest element should not be equal to the largest element.
// Input: arr[] = [12, 35, 1, 10, 34, 1]
// Output: 34
// Explanation: The largest element of the array is 35 and the second largest element is 34.

func FindSecondLargest(num []int) (int,int) {

	largest := math.MinInt 
	secondLargest := math.MinInt 

	for _, val := range num{
		if val > largest {
			secondLargest = largest
			largest = val
		}else if val > secondLargest && val != largest {
			secondLargest = val
		}
	}

	if secondLargest == math.MinInt {
		return largest,-1
	}

	return largest,secondLargest
}

/* The smallest and second smallest element
Given an array, arr[] of integers, your task is to return the smallest and second smallest element in the array. If the smallest and second smallest do not exist, return -1.
Input: arr[] = [2, 4, 3, 5, 6]
Output: [2, 3] 
Explanation: 2 and 3 are respectively the smallest and second smallest elements in the array.
Input: arr[] = [2, 4, 3, 5, 6]
Output: [2, 3] 
Explanation: 2 and 3 are respectively the smallest and second smallest elements in the array.
1 ≤ arr.size ≤105
1 ≤ arr[i] ≤ 105
*/

func FindSmallestAndSecondSmallest(nums []int) []int{

	smallest := math.MaxInt
	secondSmallest := math.MaxInt

	for _, val:= range nums {
		if val < smallest {
			secondSmallest = smallest
			smallest = val
		}else if val < secondSmallest && val != smallest {
			secondSmallest = val
		}
	}

	if secondSmallest == math.MaxInt {
		return []int{-1}
	}

	return []int{smallest,secondSmallest}

}