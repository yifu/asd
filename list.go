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

// Visit each node and apply the function.
func (l *List) Visit(f func(n *Node)) {
	p := l.head
	for p != nil {
		f(p)
		p = p.next
	}
}

// Append does append a new node at the end of the list.
func (l *List) Append(values ...interface{}) {
	for _, v := range values {
		if l.head == nil {
			l.head = &Node{v, nil}
			continue
		}
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = &Node{v, nil}
	}
}

// Push pushes a value at the front of the list.
func (l *List) Push(v interface{}) {
	l.head = &Node{v, l.head}
}

// Insert inserts an element at position pos.
func (l *List) Insert(pos int, v interface{}) error {
	if pos == 0 {
		l.head = &Node{v, l.head}
		return nil
	}

	p := l.head
	i := pos - 1
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

// Predicate is the type of function which return true when the value v satisfy the test.
type Predicate func(v interface{}) bool

// Search for an element given the provided predicate
func (l *List) Search(pred Predicate) (*Node, error) {
	p := l.head
	for p != nil {
		if pred(p.value) {
			return p, nil
		}
		p = p.next
	}
	return nil, fmt.Errorf("Value not found")
}
