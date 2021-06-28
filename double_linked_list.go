package main

type DoubleLinkNode struct {
	Data interface{}
	Next *DoubleLinkNode
	Prev *DoubleLinkNode
}

// Note: Head will always a previous node as nil
// while tail will always have next as nil also

type DoubleLinkList struct {
	head   *DoubleLinkNode
	tail   *DoubleLinkNode
	length int
}

func (list *DoubleLinkList) GetLength() int {
	return list.length
}

func (list *DoubleLinkList) Append(node DoubleLinkNode) {

	// for index 0, the head and tail is the same node and also
	// note the previous for the tail is it self and the next of the head is itself
	if list.head == nil {

		list.head = &node

		list.head.Next = list.tail
		list.tail.Prev = list.head

		list.length++

		return
	}

	currentNode := list.head

	for i := 0; i < list.length; i++ {

		if currentNode.Next == nil {

			node.Prev = currentNode
			currentNode.Next = &node

			break
		}

		// increment the current node to the next node for the next iteration
		currentNode = currentNode.Next
	}

	list.tail = &node

	list.length++

}

func (list *DoubleLinkList) Remove(node DoubleLinkNode) {

	currentNode := list.head
	nextNode := list.head.Next
	prevNode := list.head.Prev

	// Important: Note the iteration is forward iteration

	for i := 0; i < list.length; i++ {

		// edge cases for both head and tail nodes
		// remeber the next node is nil for the head and the tail previous is nil
		if currentNode.Data == node.Data {

			list.length--

			// tail's node nextNode will be nil here
			if nextNode != nil {

				nextNode.Prev = prevNode

			} else {

				// handling the case that the tail node  next is nil
				// setting prevNode.Next as nil removes it
				prevNode.Next = nil
				list.tail = prevNode

			}

			// remember that the list.head prev will be nil
			if prevNode != nil {

				prevNode.Next = nextNode

			} else {

				// handling the head edge case that the prev is nil
				nextNode.Prev = nil
				list.head = nextNode

			}

			break
		}

		// avoid nil case, which means we have gotten to the end of the list
		// if not, make the nextNode the new currentNod for the next iteration
		if nextNode != nil {
			currentNode = nextNode
		}

	}

}

func (list *DoubleLinkList) IndexOf(node DoubleLinkNode) int {

	index := -1

	currentNode := list.head

	// Note this is a forward based iteration
	for i := 0; i < list.length; i++ {

		// defensive coding againt any nil iteration, by any freak accident it happens, break
		if currentNode == nil {
			break
		}

		if currentNode.Data == node.Data {
			index = i
			break
		}

		// am making an assumption that the length will not allow the iteration of nil to happen.
		currentNode = currentNode.Next
	}

	return index

}

// removes the head of the list
func (list *DoubleLinkList) Shift() {

	if list.length == 0 {
		return
	}

	if list.length == 1 {

		list.head = nil
		list.tail = nil

		list.length--
		return
	}

	currentHead := list.head
	nextNode := currentHead.Next

	nextNode.Prev = nil
	list.head = nextNode

	list.length--

}

// pop removes the tail of the list
func (list *DoubleLinkList) Pop() {

	if list.length == 0 {
		return
	}

	if list.length == 1 {

		list.head = nil
		list.tail = nil

		list.length--
		return
	}

	currentTail := list.tail
	prevNode := currentTail.Prev

	prevNode.Next = nil
	list.tail = prevNode

	list.length--

}

func (list *DoubleLinkList) Set(index int, node DoubleLinkNode) {

	currentNode := list.head

	// Note this is a forward based iteration
	for i := 0; i < list.length; i++ {

		// defensive coding againt any nil iteration, by any freak accident it happens, break
		if currentNode == nil {
			break
		}

		if i == index {
			currentNode.Data = node.Data
			break
		}

		// am making an assumption that the length will not allow the iteration of nil to happen.
		currentNode = currentNode.Next
	}

}

func (list *DoubleLinkList) Get(index int) *DoubleLinkNode {

	var node *DoubleLinkNode = nil

	currentNode := list.head

	// Note this is a forward based iteration
	for i := 0; i < list.length; i++ {

		// defensive coding againt any nil iteration, by any freak accident it happens, break
		if currentNode == nil {
			break
		}

		if i == index {
			node = currentNode
			break
		}

		// am making an assumption that the length will not allow the iteration of nil to happen.
		currentNode = currentNode.Next
	}

	return node
}

//TODO: Implement join two different linkedlist
// func (list *DoubleLinkList) Join(index int) *DoubleLinkNode {

// 	var node *DoubleLinkNode = nil

// 	currentNode := list.head

// 	// Note this is a forward based iteration
// 	for i := 0; i < list.length; i++ {

// 		// defensive coding againt any nil iteration, by any freak accident it happens, break
// 		if currentNode == nil {
// 			break
// 		}

// 		if i == index {
// 			node = currentNode
// 			break
// 		}

// 		// am making an assumption that the length will not allow the iteration of nil to happen.
// 		currentNode = currentNode.Next
// 	}

// 	return node
// }
