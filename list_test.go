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

func TestPush_List(t *testing.T) {
	var l List
	l.Push(1)
	cmp(t, collectValues(l), []int{1})
	l.Push(2)
	cmp(t, collectValues(l), []int{2, 1})
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

	l.Insert(1, 0)
	result = collectValues(l)
	expected = []int{41, 0, 42}
	cmp(t, result, expected)

	l.Insert(1, 1)
	result = collectValues(l)
	expected = []int{41, 1, 0, 42}
	cmp(t, result, expected)

	l.Insert(1, 2)
	result = collectValues(l)
	expected = []int{41, 2, 1, 0, 42}
	cmp(t, result, expected)

	l.Insert(2, 45)
	result = collectValues(l)
	expected = []int{41, 2, 45, 1, 0, 42}
	cmp(t, result, expected)

	l.Insert(6, 33)
	result = collectValues(l)
	expected = []int{41, 2, 45, 1, 0, 42, 33}
	cmp(t, result, expected)

	err := l.Insert(8, 66)
	if err != nil {
		t.Log("err: ", err)
	} else {
		t.Fail()
	}
	result = collectValues(l)
	expected = []int{41, 2, 45, 1, 0, 42, 33}
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

func TestSearch(t *testing.T) {
	var l List
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Append(6)

	cmp(t, collectValues(l), []int{1, 2, 3, 4, 5, 6})

	search := func(val int) (*Node, error) {
		return l.Search(func(v interface{}) bool {
			return v == val
		})
	}

	expectedResult := 4
	if result, err := search(expectedResult); err != nil || result.value != expectedResult {
		t.Fail()
	}

	if result, err := search(42); err == nil {
		t.Log("result", result, "err", err)
		t.Fail()
	}
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
