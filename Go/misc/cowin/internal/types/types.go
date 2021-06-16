// all the required data structures are maintained here
package types

import (
	"os"
	"time"
)

const (
	SAMIRCID    = 1346530914
	GROUPID     = -557832891
	URL         = "https://cdn-api.co-vin.in/"
	URLPATH     = "api/v2/appointment/sessions/public/calendarByPin"
	PINCODE     = "423601"
	PINQUERY    = "pincode"
	DATEQUERY   = "date"
	HostAddress = "8081"
	WaitTime    = float64((time.Minute * 5) / 100)
)

const (
	AvailableCapacity = "Available Capacity: "
	MinAge            = "Minimum Age: "
	Vaccine           = "Vaccine: "
	Name              = "Name: "
	Available         = "Available"
	Session           = "Session"
	Slot              = "Slot"
	SessionCount      = "SessionCount"
	SlotCount         = "SlotCount"
)

type Meta struct {
	Center []Centers `json:"centers,omitempty"`
}

type Centers struct {
	Name     string     `json:"name,omitempty"`
	Address  string     `json:"address,omitempty"`
	PinCode  int        `json:"pincode,omitempty"`
	From     string     `json:"from,omitempty"`
	To       string     `json:"to,omitempty"`
	Capacity int        `json:"available_capacity,omitempty"`
	AgeLimit int        `json:"min_age_limit,omitempty"`
	Session  []Sessions `json:"sessions"`
}

type Sessions struct {
	Date              string   `json:"date,omitempty"`
	AvailableCapacity int      `json:"available_capacity,omitempty"`
	MinAge            int      `json:"min_age_limit,omitempty"`
	Vaccine           string   `json:"vaccine,omitempty"`
	Slots             []string `json:"slots,omitempty"`
}

type Needed struct {
	NumberOfSessions  int
	NumberOfSlots     int
	Name              string
	AvailableCapacity int
	MinAge            int
	Vaccine           string
	Slots             []string
}

var Token = os.Getenv("TOKEN")
var FinalMsg map[string]map[string]string
var StopFlag = false
