package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	WaitForServer("gg.pp")
}

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	var timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	tn := time.Now().Before(deadline)
	for tries := 0; tn; tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
