package main

import "fmt"

func TestLinkedList() {

	list := SingleLinkedList{}

	list.Append(Node{data: 6})
	list.Append(Node{data: 1})
	list.Append(Node{data: 8})
	list.Append(Node{data: 3})

	list.Reverse()
	fmt.Printf("list length %v \n", list.GetLength())

	list.Pop()
	fmt.Printf("list pop %v \n", list.GetLength())

	nodeTwoIndex, status := list.Get(2)

	if status && nodeTwoIndex != nil {
		fmt.Printf("list Get %v \n", nodeTwoIndex.data)
	}

}
