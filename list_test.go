package asd

import (
	"testing"
)

func collectValues(root *List) []interface{} {
	values := make([]interface{}, 0)
	if root == nil {
		return values
	}

	for root != nil {
		values = append(values, root.value)
		root = root.next
	}
	return values
}

func cmp(t *testing.T, result []interface{}, expected []int) {
	if len(result) != len(expected) {
		t.Error("Bad len. Expecting ", expected, "got ", result)
	}
	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Error("Bad value at ", i, ". Expecting ", expected, "got ", result)
		}
	}
}

func TestInsert(t *testing.T) {
	root := new(List)
	*root = List{3, nil}
	Insert(&root, 5)

	var result = collectValues(root)
	expected := []int{3, 5}
	cmp(t, result, expected)

	t.Log("values", result)
}

func TestInsert_nilRoot(t *testing.T) {
	var root *List
	Insert(&root, 4)

	var result = collectValues(root)
	expected := []int{4}
	cmp(t, result, expected)
}
