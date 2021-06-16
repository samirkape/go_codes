package main

import (
	"context"
	mp "mparser"
	bot "mybot"
	"os"
)

func main() {
	file := mp.FileHandle(mp.FILE)
	defer file.Close()
	final := mp.GetSlice(file)
	mp.Split(final)
	client := mp.GetDbClient()
	defer client.Disconnect(context.Background())
	colls := mp.ListCollections(client, mp.DbName)
	Package, err := mp.FindDoc(client, mp.DbName, colls[0])
	if err != nil {
		os.Exit(-1)
	}
	bot.SendMessage(Package.URL)
	// One time operations
	// mp.DbWritePkgs(final, client, mp.DbName)
	// mp.RemoveDuplicates(client, mp.DbName)
}
