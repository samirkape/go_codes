package main

import "time"

func SlowFunc(result chan int) {
	time.Sleep(5 * time.Second)

}
