package main

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {

	items := [...]int{1, 2, 5, 6, 7, 8}

	stack := NewStackDS()

	for i := 0; i < len(items); i++ {
		stack.Push(StackNode{
			Data: items[i],
			Next: nil,
		})
	}

	fmt.Printf("stack length %v \n", stack.GetLength())

	stack.Pop()
	stack.Pop()
	stack.Pop()

	latest := stack.Peek()

	fmt.Printf("lastest value: %v \n expected lenght %v \n", latest.Data, stack.GetLength())

}
