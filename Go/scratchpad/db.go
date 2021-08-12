package main

import (
	"fmt"
	"reflect"
	"sort"
)

type database map[string]struct{}

func (d database) Query(niddle string) bool {
	_, ok := d[niddle]
	return ok
}

func (d database) Insert(word string) {
	var empty struct{}
	d[word] = empty
}

func (d database) queryAnagram(niddle string) bool {
	m := make(map[byte]int)
	tmp := make(map[byte]int)
	var found bool

	for i := 0; i < len(niddle); i++ {
		m[niddle[i]]++
	}

	for k, _ := range d {
		found = true
		for i := 0; i < len(k); i++ {
			tmp[k[i]]++
		}
		if len(tmp) != len(m) {
			continue
		}
		for tk, _ := range m {
			if m[tk] != tmp[tk] {
				found = false
				tmp = make(map[byte]int)
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}

func (d database) queryAnagramBySort(niddle string) bool {
	nt := []byte(niddle)
	sort.Slice(nt, func(i, j int) bool {
		return nt[i] < nt[j]
	})
	for k, _ := range d {
		tmp := []byte(k)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		if reflect.DeepEqual(nt, tmp) {
			return true
		}
	}

	return false
}

func old() {
	// database = make(map[string]struct{})
	// Insert("hello")
	// Insert("world")
	// Insert("xyz")
	// fmt.Println(Query("xyz"))
	// fmt.Println(CheckAnagramBySort("xyzd"))
}

func main() {
	var d database
	d = make(map[string]struct{})
	d.Insert("hello")
	d.Insert("world")
	d.Insert("xyz")
	fmt.Println(d)
}
