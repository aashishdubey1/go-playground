package stack

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	index := len(s.items) - 1
	value := s.items[index]
	s.items = s.items[:index]
	return value, true
}

func (s *Stack) Peek() (int, bool) {

	if len(s.items) == 0 {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func RunStack() {

	stack := Stack{}

	stack.Push(5)
	stack.Push(10)
	stack.Push(15)
	stack.Push(16)

	fmt.Println(stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack)

}
