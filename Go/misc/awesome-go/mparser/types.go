// types is a part of mparser, responsible for maintaining types, variables and constants
package mparser

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const FILE = "./awesome.md"

var DBURI string

func init() {
	// load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Print(".env file found")
	}
	DBURI = os.Getenv("ATLAS_URI")
}

type config struct {
	PackageDBName string
	UserDBName    string
	UserDBColName string
	MongoURL      string
}

var Config *config

func init() {
	Config = &config{
		PackageDBName: "packagedb",
		UserDBName:    "usersdb",
		UserDBColName: "requestctr",
		MongoURL:      os.Getenv("ATLAS_URI"),
	}
}

type Category struct {
	Title          string
	PackageDetails []Package
	RawLines       []string // * [How To Code in Go eBook](https://www.digitalocean.com/community/books/how-to-code-in-go-ebook) - A 600 page introduction to Go aimed at first time developers.
	SubTitle       string
	Count          int
}

type Package struct {
	Name string `bson:"name" json:"name"`
	URL  string `bson:"url" json:"url"`
	Info string `bson:"info" json:"info"`
	// ID   primitive.ObjectID `bson:"_id" json:"id,omitempty"`
}
