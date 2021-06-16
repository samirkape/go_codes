// types is a part of mparser, responsible for maintaining types, variables and constants
package mparser

import (
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const FILE = "./awesome.md"

var DBURI string

func init() {
	DBURI = os.Getenv("ATLAS_URI")
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
	Name string             `bson:"name" json:"name"`
	URL  string             `bson:"url" json:"url"`
	Info string             `bson:"info" json:"info"`
	ID   primitive.ObjectID `bson:"_id" json:"id,omitempty"`
}
