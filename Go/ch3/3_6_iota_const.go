// Write const declarations for KB, MB, up through YB as compactly as you can

package main

import "fmt"

const (
	_ = (1 << (10 * iota))
	KB
	MB
	GB
	TB
	PB
)

func main() {
	fmt.Println("Data Units in Bytes")
	fmt.Println("KB :", KB)
	fmt.Println("MB :", MB)
	fmt.Println("GB :", GB)
	fmt.Println("TB :", TB)
	fmt.Println("PB :", PB)
}
