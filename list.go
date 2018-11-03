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
	if pos == 0 {
		l.head = &Node{v, l.head}
		return nil
	}

	p := l.head
	i := pos
	for i != 0 {
		if p.next == nil {
			return fmt.Errorf("%v is not a valid position for a %v long list", pos, pos-i)
		}
		p = p.next
		i--
	}

	p.next = &Node{v, p.next}
	return nil
}

// Clear clears a list.
func (l *List) Clear() {
	l.head = nil
}
