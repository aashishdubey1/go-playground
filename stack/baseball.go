package stack

import (
	"fmt"
	"strconv"
)

func RunBaseball() {

	// inputs := []string{"1", "2", "+", "C", "5", "D"}
	inputs := []string{"5", "D", "+", "C"}

	stack := Stack{}

	for _, v := range inputs {
		if v == "D" {
			el, _ := stack.Peek()
			stack.Push(el * 2)
		} else if v == "C" {
			stack.Pop()
		} else if v == "+" {
			top, _ := stack.Peek()
			secondLast := stack.items[len(stack.items)-2]
			stack.Push(top + secondLast)
		} else {
			sv, e := strconv.Atoi(v)
			if e != nil {
				fmt.Println(e)
			}
			stack.Push(sv)
		}
	}

	total := 0

	for _, v := range stack.items {
		total += v
	}

	fmt.Println(stack)
	fmt.Println(total)

}
