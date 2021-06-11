package main

import (
	"builder"
	"bytes"
	mylog "cwlogger"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	go mylog.RouteLog() // :8081/log will print debug logs
	Polling()
}

func Polling() {
	t := float64(time.Hour / 100)
	for {
		err := GetSlotInfo()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Duration(t)) // follow 100 requests per hour limit by cowin.gov
	}
}

// 1. build and parse url
// 2. fetch by http.Get + json decode
// 3. sending msg to bot
func GetSlotInfo() error {
	gsi_err := errors.New("GetSessionInfo: failed to get vaccine info")
	url, err := BuildQuery() // parse url
	if err != nil {
		mylog.Println(err)
		return gsi_err
	}
	data, err := FetchV2(url) // fetch url + decode json
	if err != nil {
		mylog.Println(err)
		return gsi_err
	}
	if reflect.ValueOf(data).IsZero() || len(data.Center) == 0 {
		return gsi_err
	} else {
		msg, counter := Filter(data)     // discard unnecessary data
		err := SendMessage(msg, counter) // send msg if the vaccines are available to book
		if err != nil {
			mylog.Println(err)
			return gsi_err
		}
	}
	return nil
}

// initialize and validate bot
func getTBot() (*tgbotapi.BotAPI, error) {
	if len(builder.Token) == 0 {
		return nil, errors.New("getTBot: could not find bot token")
	}
	bot, err := tgbotapi.NewBotAPI(builder.Token)
	//bot.Debug = true
	if err != nil {
		return nil, errors.New("getTBot: error getting bot")
	}
	return bot, err
}

//send vaccine information to telegram bot
func SendMessage(rmsg map[string]map[string]string, counter map[string]int) error {
	if counter["Available"] == 0 {
		return nil
	}
	bot, err := getTBot()
	if err != nil {
		return errors.New("SendMessage: could not find Bot")
	}
	var write strings.Builder
	i := 1
	for key, val := range rmsg {
		if key == "Session"+mylog.Str(i) {
			for key1, val1 := range val {
				write.WriteString(fmt.Sprintf("%s\t%s\n", key1, val1))
			}
		}
		i++
	}

	msg := tgbotapi.NewMessage(builder.SAMIRCID, write.String())
	msg1 := tgbotapi.NewMessage(builder.GROUPID, write.String())
	msg.ParseMode = "markdown"
	_, err = bot.Send(msg)
	_, err = bot.Send(msg1)

	if err != nil {
		return errors.New("SendMessage: message sending failed")
	}
	return nil
}

// get required information from []struct to map
func Filter(data builder.Meta) (map[string]map[string]string, map[string]int) {
	dlen := len(data.Center)
	slen := len(data.Center[dlen-1].Session)
	slotlen := len(data.Center[dlen-1].Session[slen-1].Slots)
	final := make(map[string]map[string]string)
	counter := make(map[string]int)
	counter["Available"] = 0
	for i := 0; i < dlen; i++ {
		session := "Session" + mylog.Str(i+1)
		ret := make(map[string]string)
		for j := 0; j < slen; j++ {
			ret["Available Capacity: "] = mylog.Str(data.Center[i].Session[j].AvailableCapacity)
			ret["Minimum Age: "] = mylog.Str(data.Center[i].Session[j].MinAge)
			ret["Vaccine: "] = data.Center[i].Session[j].Vaccine
			ret["Name: "] = data.Center[i].Name
			if data.Center[i].Session[j].AvailableCapacity > 0 {
				counter["Available"] = 1
			}
			for k := 0; k < slotlen; k++ {
				_slot := "Slot" + mylog.Str(k+1)
				ret[_slot] = data.Center[i].Session[j].Slots[k]
			}
			j++
		}
		final[session] = ret
	}
	counter["SessionCount"] = slen
	counter["SlotCount"] = slotlen
	builder.FinalMsg = final
	return final, counter
}

// simple http.Get + decode
func Fetch(url string) (builder.Meta, error) {
	var CoMeta builder.Meta
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println(err, resp.StatusCode)
		return builder.Meta{}, errors.New("fetch v1: unable to fetch URL")
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&CoMeta)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println(err, resp.StatusCode)
		return builder.Meta{}, errors.New("fetch v1: unable to decode json response")

	}
	return CoMeta, nil
}

// custom http.Get for changing header + decode response
func FetchV2(url string) (builder.Meta, error) {
	CoMeta := &builder.Meta{}
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	log.Println(url)
	if err != nil {
		return builder.Meta{}, errors.New("fetch v2: unable to fetch url")
	}

	// faking browser
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Safari/605.1.15")
	resp, err := client.Do(request)

	if resp == nil {
		return builder.Meta{}, err
	}

	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println(err, resp.StatusCode)
		return builder.Meta{}, errors.New("fetch v2: unable to fetch url")
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&CoMeta)

	if err != nil {
		return builder.Meta{}, errors.New("fetch v2: unable to decode json response")
	}
	return *CoMeta, nil
}

// dummy input for testing
func dummyJson() *bytes.Reader {
	input := []byte(`{"centers": [{"center_id": 116271, "name": "Sanvatsar PHC", "address": "Sanvatsar", "state_name": "Maharashtra", "district_name": "Ahmednagar", "block_name": "Kopargaon", "pincode": 423601, "lat": 19, "long": 74, "from": "09:00:00", "to": "17:00:00", "fee_type": "Free", "sessions": [{"session_id": "ab670e27-4e05-487b-b282-bde3a8904061", "date": "08-05-2021", "available_capacity": 0, "min_age_limit": 45, "vaccine": "COVISHIELD", "slots": ["09:00AM-11:00AM", "11:00AM-01:00PM", "01:00PM-03:00PM", "03:00PM-05:00PM"]}]}]}`)
	r := bytes.NewReader(input)
	return r
}

// date needed for query
func GetDate() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf(strconv.Itoa(day) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(year))
}

// form query by paramters
func BuildQuery() (string, error) {
	date := GetDate()
	base, err := url.Parse(builder.URL)
	if err != nil {
		return "", errors.New("BuildQuery: unable to parse url")
	}
	base.Path += builder.URLPATH
	params := url.Values{}
	params.Add(builder.DATEQUERY, date)
	params.Add(builder.PINQUERY, builder.PINCODE)
	base.RawQuery = params.Encode()
	url := base.String()
	return url, nil
}
