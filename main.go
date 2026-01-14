package main

import (
	"fmt"

	"github.com/aashishdubey1/go-playground/dsa"
)

func main(){
	fmt.Println("Starting")
	nums := []int{1, 2, 3, 4}
	res := dsa.FindRunningSum(nums)

	fmt.Print(res)

}