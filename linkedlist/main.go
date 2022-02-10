package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

func (l *linkedList) delete(v int) {
	// check emptylist
	if l.length == 0 {
		return
	}

	// check v is header
	if l.head.data == v {
		l.head = l.head.next
		l.length--
		return
	}

	toDelete := l.head
	for toDelete.next.data != v {
		// not found
		if toDelete.next.next == nil {
			return
		}
		toDelete = toDelete.next
	}

	toDelete.next = toDelete.next.next
	l.length--
}

func main() {
	myList := linkedList{}
	n1 := &node{data: 48}
	n2 := &node{data: 12}
	n3 := &node{data: 3}
	myList.prepend(n1)
	myList.prepend(n2)
	myList.prepend(n3)
	myList.printListData()
	myList.delete(12)
	myList.printListData()
}
