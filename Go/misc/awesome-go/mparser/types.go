// types is a part of mparser, responsible for maintaining types, variables and constants
package mparser

import (
	"os"
)

const FILE = "./awesome.md"

var DBURI string

func init() {
	DBURI = os.Getenv("ATLAS_URI")
	DBURI = "mongodb+srv://samirkape:Vpceh31en@cluster0.csfcu.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
}

type Package struct {
	Details Meta
}

var pkgs = make([]Package, 0)

type Meta struct {
	Title    string
	Line     LineMeta
	SubTitle string
	Count    int
}

type LineMeta struct {
	LinkDetails []SplitLink
	FullLink    []string
}

type SplitLink struct {
	Name string `bson:"name" json:"name"`
	URL  string `bson:"url" json:"url"`
	Info string `bson:"subtitle" json:"subtitle"`
}
