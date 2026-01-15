package main

import (
	"fmt"

	"github.com/aashishdubey1/go-playground/dsa"
)

func main(){
	fmt.Println("Starting")
	nums := []int{10,4,8,3}


	// fmt.Println(dsa.FindSum(nums))
	// fmt.Println(dsa.FindSecondLargest(nums))
	// fmt.Println(dsa.FindSmallestAndSecondSmallest(nums))
	fmt.Println(dsa.LeftRightDifference(nums))

}