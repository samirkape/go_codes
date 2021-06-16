package main

import (
	"context"
	mp "mparser"
)

func main() {
	file := mp.FileHandle(mp.FILE)
	defer file.Close()
	final := mp.GetSlice(file)
	mp.Split(final)
	client := mp.DbConnect()
	defer client.Disconnect(context.Background())
	mp.DbWritePkgList(final, client, mp.DbName)
}
