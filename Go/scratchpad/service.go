package main

import (
	"fmt"
	"math/rand"
)

func Service() bool {
	r := rand.Perm(5)
	return r[0] > 2
}

func GetService() {
	MAX_QUEUE := 10
	status := make(chan bool, MAX_QUEUE)

	for i := 0; i < MAX_QUEUE; i++ {
		status <- Service()
	}

	go func() {
		for i := 0; i < MAX_QUEUE; i++ {
			fmt.Println(<-status)
		}
	}()
}

func main() {
	for {
		GetService()
	}
}
