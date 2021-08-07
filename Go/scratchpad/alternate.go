package main

import (
	"fmt"
)

func chanconf(ch chan<- int) {
	fmt.Println(ch)
}

func main() {

	ch := make(chan int)
	lock := make(chan bool)

	e, o := 1, 1

	go func() {
		for o <= 100 {
			<-ch
			if o%2 == 1 {
				fmt.Println(o)
			}
			o++
		}
	}()

	go func() {
		for e <= 100 {
			ch <- 1
			if e%2 == 0 {
				fmt.Println(e)
			}
			e++
		}
		close(lock)
	}()

	<-lock
}
