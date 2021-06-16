package main

import (
	mylog "example.com/logger"
	tracker "example.com/tracker"
)

func main() {
	go mylog.RouteLog()             // :8081/log will print debug logs
	go tracker.StopACK(tracker.Bot) // this will keep checking for stop ack in background.
	tracker.Track()
}
