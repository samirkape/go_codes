package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	SAMIRCID         = 1346530914
	GROUPID          = 557832891
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

//var Token = os.Getenv("TOKEN")
var Token = "1890317276:AAE0IpeZ7hCX-FzQsTXnX1g3eBPdST2ZveQ"
var FinalMsg map[string]map[string]string

func main() {
	//go RouteLog()
	Wait()
}

func Wait() {
	t := float64(time.Hour / 100)
	go RouteLog()
	for {
		GetSessionInfo()
		time.Sleep(time.Duration(t))
	}
}

func RouteLogWriter() gin.HandlerFunc {
	write := strings.Builder{}
	fmt.Printf("%s", "\nInside Logger\n")
	return func(c *gin.Context) {
		i := 0 + 1
		for key, val := range FinalMsg {
			if key == "Session"+str(i) {
				for key1, val1 := range val {
					write.WriteString(fmt.Sprintf("%s\t%s\n", key1, val1))
				}
			}
			i++
		}
		write.WriteString(fmt.Sprintf("%v\n", time.Now()))
		c.String(http.StatusOK, write.String())
	}
}

func ConsoleLog(rmsg map[string]map[string]string) {
	i := 0 + 1
	for key, val := range rmsg {
		if key == "Session"+str(i) {
			for key1, val1 := range val {
				Println(key1 + "\t" + val1)
			}
		}
		i++
	}
}

func RouteLog() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/log", RouteLogWriter())
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, string("**hi!**"))
	})
	router.Run(":" + "80")
}

func GetSessionInfo() {
	url := BuildQuery()
	data, err := FetchV2(url)
	if err != nil {
		Println("Unable to parse URL")
		return // don't exit, rerun if request is rejected
	}
	if reflect.ValueOf(data).IsZero() {
		Println("No Data Yet")
	} else {
		msg, counter := Filter(data)
		SendMessage(msg, counter)
	}
}

func getTBot() (*tgbotapi.BotAPI, error) {
	if len(Token) == 0 {
		Println("Provide Token")
		return nil, errors.New("could not find Bot token")
	}
	bot, err := tgbotapi.NewBotAPI(Token)
	//bot.Debug = true
	if err != nil {
		Println("Error getting bot")
		return nil, err
	}
	return bot, err
}

func SendMessage(rmsg map[string]map[string]string, counter map[string]int) error {
	bot, err := getTBot()
	if err != nil {
		Println("Error getting bot")
		return errors.New("could not find Bot")
	}
	var write strings.Builder
	i := 1
	if counter["Available"] == 0 {
		return nil
	}
	for key, val := range rmsg {
		if key == "Session"+str(i) {
			for key1, val1 := range val {
				write.WriteString(fmt.Sprintf("%s\t%s\n", key1, val1))
			}
		}
		i++
	}
	msg := tgbotapi.NewMessage(SAMIRCID, write.String())
	msg1 := tgbotapi.NewMessage(GROUPID, write.String())
	msg.ParseMode = "markdown"
	_, err = bot.Send(msg)
	_, err = bot.Send(msg1)

	if err != nil {
		Println("Error Sending Message")
		return err
	}
	return nil
}

func Filter(data Meta) (map[string]map[string]string, map[string]int) {
	dlen := len(data.Center)
	slen := len(data.Center[dlen-1].Session)
	slotlen := len(data.Center[dlen-1].Session[slen-1].Slots)
	final := make(map[string]map[string]string)
	counter := make(map[string]int)
	counter["Available"] = 0

	for i := 0; i < dlen; i++ {
		session := "Session" + str(i+1)
		ret := make(map[string]string)
		for j := 0; j < slen; j++ {
			ret["Available Capacity: "] = str(data.Center[i].Session[j].AvailableCapacity)
			ret["Minimum Age: "] = str(data.Center[i].Session[j].MinAge)
			ret["Vaccine: "] = data.Center[i].Session[j].Vaccine
			ret["Name: "] = data.Center[i].Name
			if data.Center[i].Session[j].AvailableCapacity > 0 {
				counter["Available"] = 1
			}
			for k := 0; k < slotlen; k++ {
				_slot := "Slot" + str(k+1)
				ret[_slot] = data.Center[i].Session[j].Slots[k]
			}
			j++
		}
		final[session] = ret
	}
	counter["SessionCount"] = slen
	counter["SlotCount"] = slotlen
	FinalMsg = final
	return final, counter
}

func str(in interface{}) string {
	return fmt.Sprint(in)
}

func Fetch(url string) (Meta, error) {
	var CoMeta Meta
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		Println("Fetch v2: Unable to Fetch URL")
		log.Println(err, resp.StatusCode)
		return Meta{}, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&CoMeta)
	if err != nil || resp.StatusCode != http.StatusOK {
		Println("Fetch v2: Unable to Fetch URL")
		log.Println(err, resp.StatusCode)
		return Meta{}, err
	}
	return CoMeta, nil
}

func FetchV2(url string) (Meta, error) {
	CoMeta := &Meta{}
	//localAddr, err := net.ResolveIPAddr("ip", "103.112.10.253")
	//
	//if err != nil {
	//	panic(err)
	//}
	//localTCPAddr := net.TCPAddr{
	//	IP: localAddr.IP,
	//}
	//d := net.Dialer{
	//	LocalAddr: &localTCPAddr,
	//	Timeout:   30 * time.Second,
	//	KeepAlive: 30 * time.Second,
	//}
	//tr := &http.Transport{
	//	Proxy:               http.ProxyFromEnvironment,
	//	Dial:                d.Dial,
	//	TLSHandshakeTimeout: 10 * time.Second,
	//}
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	log.Println(url)
	if err != nil {
		log.Fatalln(err)
	}
	// faking browser
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Safari/605.1.15")
	//request.Header.Set("User-Agent", "User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.76 Safari/537.36")

	resp, err := client.Do(request)

	if err != nil || resp.StatusCode != http.StatusOK {
		Println("Fetch v2: Unable to Fetch URL")
		log.Println(err, resp.StatusCode)
		return Meta{}, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&CoMeta)
	//body, err := ioutil.ReadAll(resp.Body)
	//json.Unmarshal(body, &CoMeta)
	if err != nil {
		Println("Unable to Load Response")
		return Meta{}, err
	}
	return *CoMeta, nil
}

func dummyJson() *bytes.Reader {
	input := []byte(`{"centers": [{"center_id": 116271, "name": "Sanvatsar PHC", "address": "Sanvatsar", "state_name": "Maharashtra", "district_name": "Ahmednagar", "block_name": "Kopargaon", "pincode": 423601, "lat": 19, "long": 74, "from": "09:00:00", "to": "17:00:00", "fee_type": "Free", "sessions": [{"session_id": "ab670e27-4e05-487b-b282-bde3a8904061", "date": "08-05-2021", "available_capacity": 0, "min_age_limit": 45, "vaccine": "COVISHIELD", "slots": ["09:00AM-11:00AM", "11:00AM-01:00PM", "01:00PM-03:00PM", "03:00PM-05:00PM"]}]}]}`)
	r := bytes.NewReader(input)
	return r
}

func GetDate() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf(strconv.Itoa(day) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(year))
}

func BuildQuery() string {
	date := GetDate()
	base, err := url.Parse(URL)
	if err != nil {
		Println("Unable to parse URL")
		return ""
	}
	base.Path += URLPATH
	params := url.Values{}
	params.Add(DATEQUERY, date)
	params.Add(PINQUERY, PINCODE)
	base.RawQuery = params.Encode()
	url := base.String()
	return url
}

func Println(str string) {
	fmt.Fprintf(os.Stdout, "%s\n", str)
}
