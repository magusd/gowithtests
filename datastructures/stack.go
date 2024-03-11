package datastructures

import "errors"

type StackErr string

func (d StackErr) Error() string {
	return string(d)
}

const (
	ErrStackEmptyPop = StackErr("could not find remove top element from an empty stack")
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(item int) {
	n := Node{value: item}
	if s.top != nil {
		n.prev = s.top
	}
	s.top = &n
}

func (s *Stack) Pop() (int, error) {
	if s.top == nil {
		return 0, errors.New("Stack is empty")
	}
	n := s.top
	s.top = n.prev
	return n.value, nil
}

func (s *Stack) Top() int {
	n := s.top
	count := 0
	for {
		if n == nil {
			break
		}
		n = n.next
		count += 1
	}
	return count
}
