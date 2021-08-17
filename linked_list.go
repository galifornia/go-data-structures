package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Prepend(n *Node) {
	n.next = l.head
	l.length++
	l.head = n
}

func (l *LinkedList) Print() {
	var p = l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ", p.value)
		p = p.next
	}
	fmt.Println()
}

// Insert: accepts an index from 1 and a value int
func (l *LinkedList) Insert(index, value int) {
	// Abort if index is not valid
	if index > l.length+1 || index <= 0 {
		return
	}

	p := l.head
	// loop until the place right before the desired spot (index - 2)
	for i := 0; i < index-2; i++ {
		p = p.next
	}

	// Save next node
	aux := p.next
	p.next = &Node{value: value, next: aux}
	l.length++
}

// Delete: remove node at index (starts from 1)
func (l *LinkedList) Delete(index int) {
	// Abort if index is not valid
	if index > l.length || index <= 0 {
		return
	}

	if index == 1 {
		l.head = l.head.next
		l.length--
		return
	}

	p := l.head
	// loop until the place right before the node that is going away
	for i := 0; i < index-2; i++ {
		p = p.next
	}
	p.next = p.next.next
	l.length--
}

// Get: gets value of node at given index
func (l *LinkedList) Get(index int) int {
	// Abort if index is not valid
	if index > l.length || index <= 0 {
		return -1
	}

	p := l.head
	for i := 0; i < index-1; i++ {
		p = p.next
	}
	return p.value
}

// DeleteWithValue: deletes all nodes that contain the v input value
func (l *LinkedList) DeleteWithValue(v int) {
	// Abort if list is empty
	if l.length == 0 {
		return
	}

	// Head is not value to delete
	for l.head.value == v {
		l.head = l.head.next
		l.length--

		// If there are no more items exit function
		if l.length == 0 {
			return
		}
	}

	// Two pointers to jump point previous node.next to curent node.next
	p := l.head
	p2 := l.head.next

	// Loop until end of the list
	for p2 != nil {
		// If current value is v jump over unwanted node & decrease list length by 1
		if p2.value == v {
			p.next = p2.next
			p2 = p2.next
			l.length--
		}

		// Move both pointer a step forward
		p = p2
		if p2 != nil {
			p2 = p2.next
		}
	}
}

// func main() {
// 	// !FIXME: use proper tests
// 	list := LinkedList{}

// 	node1 := &Node{value: 3}
// 	node2 := &Node{value: 13}
// 	node3 := &Node{value: 5}
// 	node4 := &Node{value: 8}
// 	node5 := &Node{value: 15}
// 	node6 := &Node{value: 8}
// 	node8 := &Node{value: 8}
// 	node7 := &Node{value: 25}

// 	list.prepend(node1)
// 	list.prepend(node2)
// 	list.prepend(node3)
// 	list.prepend(node4)
// 	list.prepend(node5)
// 	list.prepend(node6)
// 	list.prepend(node7)
// 	list.prepend(node8)

// 	list.print()
// 	fmt.Println(list.Get(5))
// 	fmt.Println(list.Get(15))

// 	list.Insert(3, 77)
// 	list.Insert(0, 666)
// 	list.Insert(10, 112)

// 	list.print()

// 	list.DeleteWithValue(5)
// 	list.DeleteWithValue(8)
// 	list.DeleteWithValue(666)
// 	list.DeleteWithValue(15)
// 	list.DeleteWithValue(13)
// 	// list.DeleteWithValue(25)
// 	// list.DeleteWithValue(13)
// 	list.DeleteWithValue(3)
// 	list.DeleteWithValue(1000)
// 	list.DeleteWithValue(666)

// 	list.print()

// 	list.Delete(3)
// 	list.print()

// 	fmt.Println(list.Get(2))

// }
