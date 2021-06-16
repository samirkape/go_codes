package main

import (
	"context"
	mp "mparser"
)

func main() {
	var data []interface{}
	file := mp.FileHandle(mp.FILE)
	defer file.Close()
	final := mp.GetSlice(file)
	mp.Split(final)
	client := mp.DbConnect()
	defer client.Disconnect(context.Background())

	for i := 0; i < len(final); i++ {
		e := final[i]
		title := e.Title
		data = append(data, final)
		mp.WriteData(client, title, data)
	}
}
