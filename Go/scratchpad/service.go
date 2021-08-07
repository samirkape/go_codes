package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func Service() bool {
	r := rand.Perm(5)
	return r[0] > 2
}

func GetService(pool *sync.Pool, MAX_QUEUE int) {

	var status chan bool

	go func() {
		for i := 0; i < MAX_QUEUE; i++ {
			fmt.Println(<-status)
		}
		defer pool.Put(status)

	}()

	func() {
		status = pool.Get().(chan bool)
		for i := 0; i < MAX_QUEUE; i++ {
			status <- Service()
		}
	}()
}

func main() {
	var numPoolCreated int = 1
	MAX_QUEUE := 1000

	pool := &sync.Pool{
		New: func() interface{} {
			status := make(chan bool, numPoolCreated*MAX_QUEUE)
			numPoolCreated++
			return status
		},
	}

	pool.Put(pool.New())
	pool.Put(pool.New())
	pool.Put(pool.New())
	pool.Put(pool.New())

	for {
		GetService(pool, numPoolCreated*MAX_QUEUE)
	}
}
