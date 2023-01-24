package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (list *LinkedList) Insert(node *Node) {
	second := list.head
	list.head = node
	list.head.next = second
	list.length++
}

func (list *LinkedList) InsertAtIndex(node *Node, index int) {
	newNode := node
	if list.length == 0 {
		return
	}

	if index == 0 {
		list.Insert(newNode)
		return
	}

	if index < 0 {
		return
	}

	length := 0
	currentNode := list.head

	for length != index-1 {
		if currentNode.next == nil {
			break
		}
		currentNode = currentNode.next
		length++
	}

	// Example
	// 5  -->  4
	newNode.next = currentNode.next

	// 6  -->  5
	currentNode.next = newNode

	list.length++
}

func (list *LinkedList) RemoveAtIndex(index int) {

	if list.length == 0 {
		return
	}

	if index < 0 {
		return
	}

	length := 0
	currentNode := list.head

	for length != index-1 {
		if currentNode.next == nil {
			break
		}
		currentNode = currentNode.next
		length++
	}

	// example
	// next node is 5, set that to the next and next which is 6
	// and decrement the list length
	currentNode.next = currentNode.next.next

	list.length--
}

func (list *LinkedList) RemoveHead() {
	if list.length == 0 {
		return
	}

	newHead := list.head.next
	list.head = newHead
	list.length--
}

func (list *LinkedList) RemoveLastIndex() {
	if list.length == 0 {
		return
	}

	list.length--
}

func (list *LinkedList) PrintList() {

	toPrint := list.head

	for list.length != 0 {
		fmt.Println(toPrint.data)
		toPrint = toPrint.next
		list.length--
	}
}

func (list *LinkedList) PrintAtIndex(index int) {
	toPrint := list.head

	temporaryLength := list.length

	if index > list.length-1 || index < 0 {
		return
	}

	for temporaryLength-1 != index {
		temporaryLength--
		toPrint = toPrint.next
	}

	fmt.Println(toPrint.data)
}

func AddTwoNumbers(l1 *Node, l2 *Node) {

}

func main() {
	list := LinkedList{}

	count := 3

	for count != 0 {
		list.Insert(&Node{data: count})
		count--
	}

	AddTwoNumbers(&Node{data: 1, next: &Node{data: 2 * 3}}, &Node{data: 2 + 3, next: &Node{data: 2 * 2}})

}
