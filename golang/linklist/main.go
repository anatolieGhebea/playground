package main

import "fmt"

type node struct {
	data int
	next *node
}

type headOnlyLinkedList struct {
	head *node
}

type linkedList struct {
	head *node
	tail *node
}

func main() {
	fmt.Println("Testing head only linked list")
	myHeadOnlyList := headOnlyLinkedList{}
	myHeadOnlyList.headOnlyAppendNode(1)
	myHeadOnlyList.headOnlyAppendNode(2)
	myHeadOnlyList.headOnlyAppendNode(3)
	myHeadOnlyList.headOnlyAppendNode(4)
	myHeadOnlyList.headOnlyAppendNode(5)

	myHeadOnlyList.headOnlyPrintList()

	fmt.Println("Testing linked list")
	myList := linkedList{}
	myList.appendNode(1)
	myList.appendNode(2)
	myList.appendNode(3)
	myList.appendNode(4)
	myList.appendNode(5)

	myList.printList()

	fmt.Println("Deleting node 3")
	myList.deleteNode(3)

	fmt.Println("Deleting node 6 (not in the list)")
	myList.deleteNode(6)

	myList.printList()

	fmt.Println("Finding node 4")
	fmt.Println(myList.findNode(4))
}

func (l *headOnlyLinkedList) headOnlyAppendNode(data int) {
	newNode := &node{data: data}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (l *headOnlyLinkedList) headOnlyPrintList() {
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func (l *linkedList) appendNode(data int) {
	newNode := &node{data: data}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		l.tail = newNode
	}
}

func (l *linkedList) printList() {
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func (l *linkedList) deleteNode(data int) {
	current := l.head
	for current != nil {
		// since the list has only next node, we need to check the next node value in order to skip it if it matches the data
		if current.next != nil && current.next.data == data {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}
func (l *linkedList) findNode(data int) *node {
	current := l.head
	for current != nil {
		if current.data == data {
			return current
		}
		current = current.next
	}
	return nil
}
