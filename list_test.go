package asd

import (
	"testing"
)

func TestAppend(t *testing.T) {
	var l List
	l.Append(3)
	l.Append(5)

	var result = collectValues(l)
	expected := []int{3, 5}
	cmp(t, result, expected)

	t.Log("values", result)
}

func TestAppend_nilRoot(t *testing.T) {
	var l List
	l.Append(4)

	var result = collectValues(l)
	expected := []int{4}
	cmp(t, result, expected)
}

func TestInsert(t *testing.T) {
	var l List

	l.Insert(0, 42)
	var result = collectValues(l)
	expected := []int{42}
	cmp(t, result, expected)
}

func collectValues(l List) []interface{} {
	values := make([]interface{}, 0)

	p := l.head
	for p != nil {
		values = append(values, p.value)
		p = p.next
	}
	return values
}

func cmp(t *testing.T, result []interface{}, expected []int) {
	if len(result) != len(expected) {
		t.Error("Bad len. Expecting ", expected, "got ", result)
		return
	}

	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Error("Bad value at ", i, ". Expecting ", expected, "got ", result)
			return
		}
	}
}
