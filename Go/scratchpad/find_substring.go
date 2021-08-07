package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindIndex(sen1, sen2 string) int {
	var record []int
	var foundFlag bool = true
	var tmpIndex int

	s1 := strings.Fields(sen1)
	s2 := strings.Fields(sen2)

	for index := 0; index < len(s1); index++ {
		s1token := s1[index]
		if s1token == s2[0] {
			tmpIndex = index
			j := index + 1
			for i := 1; i < len(s2); i++ {
				if s1[j] != s2[i] {
					foundFlag = false
					index = j
					break
				}
				j++
			}
			if foundFlag {
				record = append(record, tmpIndex)
				index = j - 1
			}
			foundFlag = true
		}
	}

	if len(record) == 0 {
		return -1
	}

	sort.Ints(record)

	return record[len(record)-1]
}

func FindIndexV2(sen1, sen2 string) int {
	// var record []int
	var foundFlag bool = true
	// var tmpIndex int
	var res int

	s1 := strings.Fields(sen1)
	s2 := strings.Fields(sen2)

	for index := 0; index < len(s1); index++ {
		s1token := s1[index]
		if s1token == s2[0] {
			//tmpIndex = index
			j := index + 1
			for i := 1; i < len(s2); i++ {
				if s1[j] != s2[i] {
					foundFlag = false
					index = j
					break
				}
				j++
			}
			if foundFlag {
				res = j
				index = j - 1
			}
			foundFlag = true
		}
	}

	// if len(record) == 0 {
	// 	return -1
	// }

	//sort.Ints(record)

	return res - len(s2)
}

func main() {
	sen1 := `some people over age 60 have few indicate hypothyroidism in an if any, symptoms of an some people underactive thyroid gland (hypothyroidism), while others experience the same symptoms some people younger people do. Still others have hypothyroidism symptoms that are not typical at all, making the diagnosis even more difficult. Any of the following signs and symptoms can indicate hypothyroidism in an older person. `
	sen2 := "some people"
	f := strings.Fields(sen1)
	_ = f
	index := FindIndex(sen1, sen2)
	fmt.Println(index)
}
