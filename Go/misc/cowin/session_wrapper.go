package main

import "time"

func main() {
	t := float64(time.Minute / 100)
	for {
		GetSessionInfo()
		time.Sleep(time.Duration(t))
	}
}
