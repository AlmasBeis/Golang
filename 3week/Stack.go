package main

import "fmt"

type Node struct {
	value int
	next  *Node
}
type Stack struct {
	top  *Node
	size int
}

func (s *Stack) Push(value int) {
	node := &Node{value: value, next: s.top}
	s.top = node
	s.size++
}

func (s *Stack) Pop() int {
	if s.size == 0 {
		return 0
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value
}

func (s *Stack) Peek() int {
	if s.size == 0 {
		return 0
	}
	return s.top.value
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Clear() {
	for i := 0; i < s.Size(); i++ {
		s.Pop()
	}
}
func (s *Stack) Contains() bool {
	if s.Size() == 0 {
		return false
	}
	return true
}

func (s *Stack) Increment() {
	newNode := s.top
	for newNode != nil {
		newNode.value++
		newNode = newNode.next
	}
}
func (s *Stack) Print() {
	newNode := s.top
	for newNode != nil {
		fmt.Print(newNode.value, " ")
		newNode = newNode.next
	}
}

func (s *Stack) PrintReverse() {
	st := &Stack{}
	newNode := s.top
	for newNode != nil {
		st.Push(newNode.value)
		newNode = newNode.next
	}
	for st.size != 0 {
		fmt.Print(st.Pop(), " ")
	}
	fmt.Println()
}

func main() {
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Size())

	value := stack.Peek()
	fmt.Println(value)
	stack.Increment()
	value1 := stack.Peek()
	fmt.Println(value1)
	value = stack.Pop()
	fmt.Println(value)
	fmt.Println(stack.Size())
	stack.Pop()
	stack.Pop()
	fmt.Println(stack.Size())
	stack.Push(2)
	stack.Push(3)
	stack.Increment()
	stack.Print()
	println()
	stack.PrintReverse()
}
