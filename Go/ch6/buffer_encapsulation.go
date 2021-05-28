package main

import "fmt"

type Buffer struct {
	buf     []byte
	initial [64]byte
	/* ... */
}

// Grow expands the buffer's capacity, if necessary,
// to guarantee space for another n bytes. [...]
func (b *Buffer) Grow(n int) {
	if b.buf == nil {
		b.buf = b.initial[:0] // use preallocated space initially
	}
	if len(b.buf)+n > cap(b.buf) {
		buf := make([]byte, b.Len(), 2*cap(b.buf)+n)
		copy(buf, b.buf)
		b.buf = buf
	}
}

func (b *Buffer) Len() int {
	if b.buf == nil {
		return 0
	} else {
		return len(b.buf)
	}
}

func (b *Buffer) Cap() int {
	if b.buf == nil {
		return 0
	} else {
		return cap(b.buf)
	}
}

func main() {
	var b = Buffer{}
	b.Grow(5)
	fmt.Println(b.Len())
	fmt.Println(b.Cap())
}
