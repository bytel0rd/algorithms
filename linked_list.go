package main

type Node struct {
	data int
	next *Node
}

type SingleLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (list *SingleLinkedList) Append(node Node) {

	if list.length == 0 {

		list.head = &node
		list.tail = &node
		list.length++

		return
	}

	currentTail := list.tail

	currentTail.next = &node

	list.tail = &node

	list.length++
}

func (list *SingleLinkedList) Remove(node Node) {

	if list.head == nil {
		return
	}

	if list.head.data == node.data {
		list.head = list.head.next
		list.length--
		return
	}

	parentNode := list.head

	currentNode := list.head.next

	for i := 0; i < list.length; i++ {

		if currentNode != nil && currentNode.data == node.data {

			parentNode.next = currentNode.next

			if currentNode.next == nil {
				list.tail = parentNode
			}

			list.length--
			return
		}

		parentNode = currentNode
		currentNode = currentNode.next

	}

}

func (list *SingleLinkedList) Includes(node Node) bool {

	currentNode := list.head

	for currentNode.next != nil {
		if currentNode.data == node.data {
			return true
		}

		currentNode = currentNode.next
	}

	return false
}

func (list *SingleLinkedList) GetLength() int {
	return list.length
}

func (list *SingleLinkedList) Pop() {

	if list.head == nil {
		return
	}

	if list.head.next == nil {
		list.head = nil
		list.length--
		return
	}

	var nextToTailNode *Node = list.head.next

	for i := 0; i < list.length; i++ {

		if nextToTailNode.next.next == nil {
			nextToTailNode.next = nil
			list.length--
			return
		} else {
			nextToTailNode = nextToTailNode.next
		}

	}
}

func (list *SingleLinkedList) Shfit() {

	if list.head == nil {
		return
	}

	if list.head.next == nil {
		list.head = nil
	} else {
		list.head = list.head.next
	}

	list.length--

}

func (list *SingleLinkedList) UnShift(node Node) {

	currentHead := list.head

	node.next = currentHead

	list.head = &node

}

func (list *SingleLinkedList) isValidIndex(index int) bool {
	return index > 0 && index < list.length
}

func (list *SingleLinkedList) Get(index int) (*Node, bool) {

	if list.isValidIndex(index) {
		return nil, false
	}

	currentNode := list.head

	for i := 0; i < index+1; i++ {

		if i == index {
			return currentNode, true
		}

		currentNode = currentNode.next
	}

	return nil, false

}

func (list *SingleLinkedList) Set(index int, node Node) {

	if list.isValidIndex(index) {
		return
	}

	currentNode := list.head

	for i := 0; i < index+1; i++ {

		if i == index {
			currentNode.data = node.data
			return
		}

		currentNode = currentNode.next
	}

}

func (list *SingleLinkedList) Reverse() {

	currentNode := list.head
	list.tail = list.head

	var nextNode *Node = nil
	var prevNode *Node = nil

	for i := 0; i < list.length; i++ {

		// storing the next chain before overiding
		nextNode = currentNode.next

		// changing the next node to previous node, for 1 next = nil.
		currentNode.next = prevNode

		// make the current node the previous node for the next iteration
		prevNode = currentNode

		// make the nextNode the new current node for the next iteration
		currentNode = nextNode
	}

	list.head = prevNode

}
