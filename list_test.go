package asd

import (
	"testing"
)

func TestAppend(t *testing.T) {
	var l List
	l.Append(3)
	l.Append(5)

	result := collectValues(l)
	expected := []int{3, 5}
	cmp(t, result, expected)

	t.Log("values", result)
}

func TestAppend_nilRoot(t *testing.T) {
	var l List
	l.Append(4)

	result := collectValues(l)
	expected := []int{4}
	cmp(t, result, expected)
}
func TestInsert(t *testing.T) {
	var l List

	l.Insert(0, 42)
	result := collectValues(l)
	expected := []int{42}
	cmp(t, result, expected)

	l.Insert(0, 41)
	result = collectValues(l)
	expected = []int{41, 42}
	cmp(t, result, expected)
}

func TestClear(t *testing.T) {
	var l List
	l.Append(1)
	l.Append(2)
	l.Append(3)

	var result = collectValues(l)
	var expected = []int{1, 2, 3}
	cmp(t, result, expected)

	l.Clear()
	cmp(t, collectValues(l), []int{})
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
