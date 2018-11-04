package asd

import "testing"

type Person struct {
	name string
	age  int
	ssid int
}

func (p *Person) Key() Key {
	return Key(p.ssid)
}
func TestHashtable_Insert(t *testing.T) {
	ht := New()

	personList := []*Person{
		&Person{"toto", 3, 0},
		&Person{"Jean", 34, 123},
		&Person{"Sophie", 23, 35},
		&Person{"Avner", 35, 124},
		&Person{"Sardoche", 31, 125},
		&Person{"Ismaïl", 27, 125}}

	for _, p := range personList {
		ht.Insert(p)
	}
	for _, p := range personList {
		if _, err := ht.Search(p); err != nil {
			t.Fail()
		}
	}
}
