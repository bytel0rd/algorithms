package main

type StackNode struct {
	Data interface{}
	Next *StackNode
}

type Stack struct {

	// most recently added node.
	top *StackNode

	// the most oldest node.
	buttom *StackNode

	// the length of the total node.
	length int
}

func NewStackDS() Stack {
	return Stack{
		top:    nil,
		buttom: nil,
		length: 0,
	}
}

func (s *Stack) GetLength() int {
	return s.length
}

func (s *Stack) IsEmpty() bool {
	return s.length == 0
}

// return the top of the stack
func (s *Stack) Peek() *StackNode {
	return s.top
}

// i want to avoid looping

func (s *Stack) Pop() bool {

	if s.top == nil {
		return false
	}

	var prevNode *StackNode = nil
	var currentNode *StackNode = s.buttom

	for i := 0; i < s.length; i++ {

		if currentNode.Next == nil {

			// the previous will only be nil for the first iteration
			if prevNode != nil {
				// deletes the most recent
				prevNode.Next = nil
			} else {
				s.buttom = nil
			}
			// resets the top node to the prevNode
			s.top = prevNode
			s.length--

			return true
		}

		// changing previous and current node over the iteration
		prevNode = currentNode
		currentNode = currentNode.Next

	}

	return false

}

func (s *Stack) Push(node StackNode) {

	if s.length == 0 {
		s.top = &node
		s.buttom = &node
		s.length++

		return
	}

	s.top.Next = &node
	s.top = &node

	s.length++
}
