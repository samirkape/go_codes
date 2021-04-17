// Implement errorf using variadic functions

package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	_, fn, line, _ := runtime.Caller(2) // 0 = This file , 1 = Go runtime, 2 = ASM
	errorf(line, fn, "ERROR")
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d:", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
