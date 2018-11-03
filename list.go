package asd

// List implements a single list structure.
type List struct {
	value interface{}
	next  *List
}

// Insert does insert a new node.
func Insert(head **List, v interface{}) {
	if *head == nil {
		*head = new(List)
		*(*head) = List{v, nil}
		return
	}
	node := List{v, (*head).next}
	(*head).next = &node
}
