// Creating and deleting temp dirs

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	scope1()
	scope2()
}

func scope1() {
	var rmdirs []func()
	dirs := []string{"./xx", "./xy", "./xz"}
	for i := 0; i < len(dirs); i++ {
		os.MkdirAll(dirs[i], 0755) // OK
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dirs[i]) // NOTE: incorrect!
		})
	}
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}

func scope2() {
	var rmdirs []func()
	tempDirs := []string{"./xx", "./xy", "./xz"}

	for _, dir := range tempDirs { // For each iteration, range re-uses same address for dir
		d := dir // NOTE: necessary!
		fmt.Printf("dir=%v, *dir=%p\n", dir, &dir)
		fmt.Printf("d=%v, *d=%p\n", d, &d)
		err := os.MkdirAll(dir, 0755) // creates parent directories too
		if err != nil {
			log.Fatal(err)
		}
		rmdirs = append(rmdirs, func() {
			err := os.RemoveAll(dir)
			if err != nil {
				log.Fatal(err)
			}
		})
	}
	// ...do some work...
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}
