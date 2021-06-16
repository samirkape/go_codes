package main

import (
	"context"
	mp "mparser"
	bot "mybot"
	"os"
)

// raw function calls for fetching and
// sending a package URL to a telegram bot
func main() {

	// Parse markdown file
	// file := mp.FileHandle(mp.FILE)
	// defer file.Close()
	// final := mp.GetSlice(file)
	// mp.Split(final)

	// write parsed package meta to database
	// mp.DbWritePkgs(final, client, mp.DbName)
	// mp.RemoveDuplicates(client, mp.DbName)

	// Connect to database and fetch a document from collection list
	client := mp.GetDbClient()
	defer client.Disconnect(context.Background())
	colls := mp.ListCollections(client, mp.DbName)
	Package, err := mp.FindDoc(client, mp.DbName, colls[0])
	if err != nil {
		os.Exit(-1)
	}

	// send sample message to telegram bot
	bot.SendMessage(Package.URL)
}
