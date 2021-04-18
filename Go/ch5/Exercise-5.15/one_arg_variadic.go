/* Exercise 5.15: Write variadic functions max and min, analogous to sum. What should these functions do when called with no arguments? Write variants that require at least one argument. */

package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println(min(1, 2, 3, 4, 5))
	fmt.Println(max(1, 2))
}

func min(arg ...int) (tMin int, err error) {
	if len(arg) < 1 {
		err := errors.New("min: error: atleast 1 argument needed")
		return -1, err
	}
	tMin = math.MaxInt64
	for _, e := range arg {
		if e < tMin {
			tMin = e
		}
	}
	return
}

func max(arg ...int) (tMax int, err error) {
	if len(arg) < 1 {
		err := errors.New("max: error: atleast 1 argument needed")
		return -1, err
	}
	tMax = math.MinInt64
	for _, e := range arg {
		if e > tMax {
			tMax = e
		}
	}
	return
}
