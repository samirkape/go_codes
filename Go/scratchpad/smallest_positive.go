package main

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

func Solution_arr(A []int) int {
	isAllNeg := true

	set := make(map[int]struct{}, len(A))
	var empty struct{}

	for _, e := range A {
		set[e] = empty
		if e > 0 {
			isAllNeg = false
		}
	}

	if isAllNeg {
		return 1
	}

	nodups := make([]int, len(set))

	i := 0
	for k, _ := range set {
		nodups[i] = k
		i++
	}

	sort.Ints(nodups)

	for i := 0; i+1 < len(nodups); i++ {
		if nodups[i+1]-nodups[i] != 1 {
			return nodups[i] + 1
		}
	}
	return nodups[len(nodups)-1] + 1
}

func Solution1(S string) bool {
	aPresent := false
	bPresent := false
	j := 0

	for i := 0; i < len(S); i++ {
		if S[i] == 'a' {
			aPresent = true
		} else if S[i] == 'b' {
			bPresent = true
			for j = i; j < len(S); j++ {
				if S[j] == 'a' {
					return false
				}
			}
			break
		}
	}
	if aPresent == false {
		return true
	} else if bPresent == false {
		return true
	}
	return true
}

func Solution2(S string) string {
	var occurrences [26]int
	for _, ch := range S {
		occurrences[int(ch)-int('a')]++
	}

	var best_char uint8 = 'a'
	var best_res int = 0
	var i int

	for i = 0; i < 26; i++ {
		if occurrences[i] > best_res {
			best_char = uint8(int('a') + i)
			best_res = occurrences[i]
		}
	}

	return string(best_char)
}

func getBinary(S string) *big.Int {
	place := big.NewInt(1)
	dec := big.NewInt(0)
	two := big.NewInt(2)
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '1' {
			dec.Add(dec, place)
			// fmt.Println(dec)
		}
		place.Mul(place, two)
	}
	return dec
}

func Solution(S string) int {

	num := getBinary(S)
	count := 0
	zero := big.NewInt(0)
	tmp := big.NewInt(0)
	two := big.NewInt(2)
	one := big.NewInt(1)

	// fmt.Println(num.Sub(num, one))

	for num.Cmp(zero) != 0 {
		tmp = tmp.Mod(num, two)
		if tmp.Cmp(zero) == 0 {
			num.Div(num, big.NewInt(2))
		} else {
			num.Sub(num, one)
		}
		// fmt.Println(big.NewInt(0).Sub(num, one))
		count++
	}
	return count
}

func main() {
	S := strings.Repeat("1010", 400000/4)
	// S = "1111010101111"
	fmt.Println(Solution(S))
}
