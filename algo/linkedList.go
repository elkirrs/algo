package algo

import (
	"fmt"
	"math/rand"
)

type NodeList struct {
	Value int
	Next  *NodeList
}

type List struct {
	Size int
	Root *NodeList
}

func LinkedList() {
	list := &List{}

	iterations := rand.Intn(10)
	for i := 1; i <= iterations; i++ {
		list.add(rand.Intn(10) + i)
	}

	list.printList()
}

func (list *List) add(value int) {
	if list.Size == 0 {
		list.Root = &NodeList{
			Value: value,
			Next:  nil,
		}
		list.Size++
		return
	}

	node := list.Root
	for node.Next != nil {
		node = node.Next
	}

	newNode := &NodeList{
		Value: value,
		Next:  nil,
	}
	node.Next = newNode
	list.Size++
}

func (list *List) getSize() int {
	return list.Size
}

func (list *List) printList() {
	var linkedList []int

	for node := list.Root; node != nil; node = node.Next {
		linkedList = append(linkedList, node.Value)
	}

	fmt.Println(linkedList)
}
