package main

import (
	"fmt"
	"sync"
)

func gr_a(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("In goroutine a")
}

func gr_b(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("In goroutine b")
}

func gr_c(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("In goroutine c")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go gr_a(&wg)
	go gr_b(&wg)
	wg.Wait()
	wg.Add(1)
	go gr_c(&wg)
	wg.Wait()
}
