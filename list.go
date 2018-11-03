package asd

import "fmt"

// Node implements a Node in a linked list.
type Node struct {
	value interface{}
	next  *Node
}

// List implement a linked list.
type List struct {
	head *Node
}

// Append does append a new node at the end of the list.
func (l *List) Append(v interface{}) {
	if l.head == nil {
		l.head = &Node{v, nil}
		return
	}
	p := l.head
	for p.next != nil {
		p = p.next
	}
	p.next = &Node{v, nil}
}

// Insert inserts an element at position pos.
func (l *List) Insert(pos int, v interface{}) error {
	if l.head == nil {
		if pos != 0 {
			return fmt.Errorf("%v is not a valid position", pos)
		}
		l.head = &Node{v, nil}
		return nil
	}
	p := l.head
	var i int
	for i = 0; i < pos && p.next != nil; i++ {
		p = p.next
	}
	if i != pos {
		return fmt.Errorf("%v is not a valid position", pos)
	}
	p.next = &Node{v, nil}
	return nil
}
