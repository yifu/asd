package asd

import (
	"fmt"
	"log"
	"strings"
)

// BSTNode is a binary search tree node.
type BSTNode struct {
	value       int
	left, right *BSTNode
}

// Insert a node inside the bst.
func (root *BSTNode) Insert(v int) *BSTNode {
	if root == nil {
		return &BSTNode{v, nil, nil}
	}
	if v <= root.value {
		root.left = root.left.Insert(v)
	} else {
		root.right = root.right.Insert(v)
	}
	return root
}

// FromHRInputToBST expect an input as given from the Hacker Rank website and populates a BST tree from It.
// (Note it should be places inside a asdutil package).
func FromHRInputToBST(str string) (root *BSTNode) {
	r := strings.NewReader(str)

	n := 0
	if _, err := fmt.Fscanln(r, &n); err != nil {
		log.Panic(err)
	}

	for i := 0; i < n; i++ {
		var v int
		fmt.Fscan(r, &v)
		root = root.Insert(v)
	}
	return
}

// TopOfView gives the top of a view of the bst.
func (root *BSTNode) TopOfView() []int {
	if root == nil {
		return []int{}
	}
	var result []int
	result = append(result, root.left.topOfViewLeft()...)
	result = append(result, []int{root.value}...)
	result = append(result, root.right.topOfViewRight()...)
	return result
}

func (root *BSTNode) topOfViewLeft() []int {
	if root == nil {
		return []int{}
	}
	var result []int
	result = append(result, root.left.topOfViewLeft()...)
	result = append(result, []int{root.value}...)
	return result
}

func (root *BSTNode) topOfViewRight() []int {
	if root == nil {
		return []int{}
	}
	var result []int
	result = append(result, []int{root.value}...)
	result = append(result, root.right.topOfViewRight()...)
	return result
}
