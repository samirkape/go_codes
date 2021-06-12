package types

const (
	SAMIRCID         = 1346530914
	GROUPID          = -557832891
	URL       string = "https://cdn-api.co-vin.in/"
	URLPATH   string = "api/v2/appointment/sessions/public/calendarByPin"
	PINQUERY  string = "pincode"
	PINCODE   string = "423601"
	DATEQUERY string = "date"
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

//var Token = os.Getenv("TOKEN")
var Token = "1890317276:AAE-gL26gJbOxDm_3Ndhjv-kWiYDN9yni0o"
var FinalMsg map[string]map[string]string
var StopFlag = false
