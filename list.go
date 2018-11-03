package asd

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
