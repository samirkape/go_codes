// Find basename and dirname of a path string

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := os.Args[1]
	base := basename(path)
	dir := dirname(path)
	fmt.Println("Path Components:", base, dir)
}

func basename(path string) string {
	slash_idx := strings.LastIndex(path, "/")
	if slash_idx == -1 {
		fmt.Println("Path not found...")
		return path
	} else {
		fmt.Printf("\nOriginal Path : %s\n", path)
		fmt.Printf("basename : %s\n\n", path[slash_idx+1:])
		return path[slash_idx+1:]
	}
}

func dirname(path string) string {
	slash_idx := strings.LastIndex(path, "/")
	if slash_idx == -1 {
		fmt.Println("Path not found...")
		return path
	} else {
		fmt.Printf("\nOriginal Path : %s\n", path)
		fmt.Printf("basename : %s\n\n", path[:slash_idx])
		return path[:slash_idx]
	}
}
