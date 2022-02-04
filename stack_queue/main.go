package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}

	toRemove := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
}
