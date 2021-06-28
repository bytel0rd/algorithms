package main

import (
	"fmt"
	"testing"
)

func TestQue(t *testing.T) {

	items := [...]int{1, 2, 5, 6, 7, 8}

	que := NewQue()

	for i := 0; i < len(items); i++ {
		entry := NewQueEntry(items[i])
		que.Add(entry)
	}

	fmt.Printf("stack length %v \n", que.GetLength())

	que.Remove()
	que.Remove()

	next := que.Peek()

	fmt.Printf("next value: %v \n expected lenght %v \n", next.Data, que.GetLength())

}
