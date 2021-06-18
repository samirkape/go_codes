// types is a part of mparser, responsible for maintaining types, variables and constants
package mybot

import (
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Package[i] is a master DS for a category.
// it contains short description, title and all subpackages of a category.
var pkgs []Package

// mongodb URI to establish connection with database.
var DBURI string

// mongodb daatabse name
const DbName = "packagedb"
const UserDbName = "usersdb"

// local markdown file for parsing
const FILE = "./awesome.md"

// init mongo db instance
func init() {
	pkgs = make([]Package, 0)
	DBURI = os.Getenv("ATLAS_URI")
}

const (
	CMDStart          = "/start"
	CMDListCategories = "/listcategories"
	CMDListPackages   = "/selectentry"
)

// Below structs are used for parsing the incoming POST request from telegram bot.
// root level structure
type ReceiveMessage struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message struct
// holds information about complete message that includes chat id, msg text etc.
type Message struct {
	MessageID int        `json:"message_id"`
	From      From       `json:"from"`
	Chat      Chat       `json:"chat"`
	Date      int        `json:"date"`
	Text      string     `json:"text"`
	Entities  []Entities `json:"entities"`
}

// From struct
// holds information about the sender.
type From struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	UserName     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

// Chat struct
// it holds the meta data of chat that includes id, msg type (text, image, etc.)
type Chat struct {
	ID                          int    `json:"id"`
	FirstName                   string `json:"first_name"`
	UserName                    string `json:"username"`
	Type                        string `json:"type"`
	Title                       string `json:"title"`
	AllMembersAreAdministrators bool   `json:"all_members_are_administrators"`
}

// Entities struct
// Unused, written for future use
type Entities struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

// below data structures are related to parsing of markdown file
// root structure
type Package struct {
	Details Meta
}

// root structure of a category meta data
type Meta struct {
	Title    string
	Line     LineMeta
	SubTitle string
	Count    int
}

// this structure holds a information for multiple single lines.
// i.e it stores multiple raw lines related with package that belong to certain category.
type LineMeta struct {
	LinkDetails []SplitLink
	FullLink    []string
}

// this is final structure of parser which will also be use for inserting package into database.
type SplitLink struct {
	Name string             `bson:"name" json:"name"`
	URL  string             `bson:"url" json:"url"`
	Info string             `bson:"info" json:"info"`
	ID   primitive.ObjectID `bson:"_id" json:"id,omitempty"`
}
