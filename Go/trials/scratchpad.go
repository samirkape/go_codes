// Rough codes separated by the functions

package main

import (
	"fmt"
	"reflect"
)

var m = make(map[string]int)

func main() {
	FunctionValues()
	// mMapExp()
	// rReflect()
	// sStruct()
	//slice_append_variac()
}
func add(x, y int) int { return x + y }
func FunctionValues() {
	var f func(int, int) int
	fmt.Printf("%T\n", f)
	f = add
	fmt.Printf("%T\n", f)
}

func sStruct() {
	type s_try struct{ a, b int }
	pp := s_try{b: 2}
	fmt.Println(pp)
}

func rReflect() {
	hw := "Samir Kape"
	s1 := hw[:5]
	s2 := hw[0:5]
	fmt.Println(reflect.DeepEqual(s1, s2))
}

func mMapExp() {

	hw := "Samir Kape"
	sl := mStringtoSlice(hw)
	// str1 := mSlicetoString(sl)
	mAdd(sl)
	ml := m
	fmt.Println(ml)
}

func mAdd(ls []string) {
	m[mSlicetoString(ls)]++
}

func mCount(ls []string) int {
	return m[mSlicetoString(ls)]
}

func SliceAppendVariac() {
	s := []int{1, 2, 3, 4, 5}
	s = append(s, s...) // 3 dots are neccesary if you want to append a slice into the slice
}

func mStringtoSlice(s string) []string {
	var slice []string
	for _, i := range s {
		slice = append(slice, string(i))
	}
	return slice
}

func mSlicetoString(sltr []string) string {
	return fmt.Sprintf("%q", sltr)
}

// func slice_to_string2(sltr []string) string { // Does not work
// 	return string(sltr)
// }
