package dsa

/*
1. Given an array, arr[] of n integers, and an integer element x, find whether element x is present in the array. Return the index of the first occurrence of x in the array, or -1 if it doesn't exist.
Input: arr[] = [1, 2, 3, 4], x = 3
Output: 2
Explanation: For array [1, 2, 3, 4], the element to be searched is 3. Since 3 is present at index 2, the output is 2.
*/

func FindEl(nums []int, num int) int {
	var index int
	for i , val := range nums {
		if val == num {
			index = i
			break
		}
	}
	return  index
}


/* 
2. 


*/