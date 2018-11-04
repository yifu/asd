package asd

import "fmt"

// Key is the type to use for the search key.Key
// TODO rename as SearchKey ?
type Key int

// Keyer is an interface for Hashable objects.
type Keyer interface {
	Key() Key
}

// HashTable implements a simple open-chaining hash table.
type HashTable struct {
	table []List
}

// TODO make() function. There must be an initial length to the hastable. And we must upgrade it when the ratio is going over a threshold: (n / k) > 2 for instance.

// New creates a new HashTable.
func New() *HashTable {
	return &HashTable{make([]List, 10)}
}

func (ht *HashTable) hash(sk Key) int {
	if len(ht.table) == 0 {
		panic("The hash table has not been initialised.")
	}
	return int(sk) % len(ht.table)
}

// Insert inserts or updates an element in the hastable
func (ht *HashTable) Insert(v Keyer) {
	sk := v.Key()
	pos := ht.hash(sk)
	l := &ht.table[pos]
	node, err := l.Search(func(v interface{}) bool {
		return v.(Keyer).Key() == sk
	})
	if err != nil {
		//fmt.Println("Insert")
		l.Push(v)
	} else {
		//fmt.Println("Insert-Update")
		node.value = v
	}
}

// Search lookups for an element with Key k in the Hashtable.
func (ht *HashTable) Search(v Keyer) (*Node, error) {
	sk := v.Key()
	pos := ht.hash(sk)
	l := ht.table[pos]
	return l.Search(func(v interface{}) bool {
		return v.(Keyer).Key() == sk
	})
}

// Display the hashtable
func (ht *HashTable) Display() {
	for _, l := range ht.table {
		//fmt.Println("l=", l)
		l.Visit(func(v *Node) { fmt.Println(v) })
	}
}
