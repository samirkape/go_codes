package main

import (
	"bytes"
	"fmt"
	"os"
)

type Storage struct {
	buf []byte
}

func (s *Storage) Write(p []byte) (int, error) {
	s.buf = p
	return len(s.buf), nil
}

func Printf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stdout, format, args...) // Writing to stdout
}

func Sprintf(format string, args ...interface{}) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, format, args...) //
	return buf.String()
}

func eSprintf(format string, args ...interface{}) string {
	var buffer Storage
	fmt.Fprintf(&buffer, format, args...) // buffer should have Write() method to be qualified for input as a io.Writer
	return string(buffer.buf)
}

func main() {
	fmt.Println(eSprintf("Samir %s", "Kape"))
}
