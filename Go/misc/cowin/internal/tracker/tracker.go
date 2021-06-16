// main package for getting vaccine slot availability information from cowin.org and if more than one slots are available to book, then sending message to appropriate user by a means of telegram bot.
package tracker

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

	mylog "example.com/logger"
	types "example.com/types"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var Bot *tgbotapi.BotAPI

// bot constructor
func init() {
	bot, err := getTBot()
	Bot = bot
	if err != nil {
		fmt.Println("bot initialization failed")
		os.Exit(-1)
	}
}

// driver function
func Track() {
	for {
		err := GetSlotInfo()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Duration(types.WaitTime)) // follow 100 requests per 5 minutes limit by cowin.gov
	}
}

// 1. construct and parse url
// 2. fetch by http.Get + json decode
// 3. sending msg to bot
func GetSlotInfo() error {
	gsi_err := "GetSessionInfo: failed to get vaccine info %v"
	url, err := BuildQuery() // parse url
	if err != nil {
		mylog.Println(err)
		return fmt.Errorf(gsi_err, err)
	}
	data, err := FetchV2(url) // fetch url + decode json
	if err != nil {
		mylog.Println(err)
		return fmt.Errorf(gsi_err, err)
	}
	if reflect.ValueOf(data).IsZero() || len(data.Center) == 0 {
		return errors.New(gsi_err)
	} else {
		msg, counter := Filter(data)        // discard unnecessary data
		err := MessageHandler(msg, counter) // send msg if the vaccines are available to book
		if err != nil {
			mylog.Println(err)
			return fmt.Errorf(gsi_err, err)
		}
	}
	return nil
}

// initialize and validate bot
func getTBot() (*tgbotapi.BotAPI, error) {
	if len(types.Token) == 0 {
		return nil, errors.New("getTBot: could not find bot token")
	}
	bot, err := tgbotapi.NewBotAPI(types.Token)
	//bot.Debug = true
	if err != nil {
		return nil, fmt.Errorf("getTBot: error initializing bot: %v", err)
	}
	return bot, err
}

//send vaccine information to telegram bot
func MessageHandler(rmsg map[string]map[string]string, counter map[string]int) error {
	if counter[types.Available] == 0 {
		return nil
	}
	var write strings.Builder
	i := 1
	for key, val := range rmsg {
		if key == types.Session+mylog.Str(i) {
			for key1, val1 := range val {
				write.WriteString(fmt.Sprintf("%s\t%s\n", key1, val1))
			}
			write.WriteString("--- ")
		}
		i++
	}
	msg := write.String()
	SendMessage(Bot, msg)
	return nil
}

// sends message to registered id
func SendMessage(bot *tgbotapi.BotAPI, Info string) error {
	if types.StopFlag {
		fmt.Println("warning! bot has recieved an ack to stop")
		os.Exit(-1)
		return nil
	}
	msg := tgbotapi.NewMessage(types.SAMIRCID, Info)
	//msg1 := tgbotapi.NewMessage(types.GROUPID, Info)
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)
	//_, err = bot.Send(msg1)
	if err != nil {
		return fmt.Errorf("sendmessage: message sending failed: %v", err)
	}
	return nil
}

// checks for any msg from bot
func StopACK(bot *tgbotapi.BotAPI) {
	if types.StopFlag {
		return
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		types.StopFlag = false
	}
	for update := range updates {
		if update.Message != nil {
			types.StopFlag = true
		}
	}
}

// get required information from []struct to map
func Filter(data types.Meta) (map[string]map[string]string, map[string]int) {
	dlen := len(data.Center)
	slen := len(data.Center[dlen-1].Session)
	slotlen := len(data.Center[dlen-1].Session[slen-1].Slots)
	final := make(map[string]map[string]string)
	counter := make(map[string]int)
	counter[types.Available] = 0
	for i := 0; i < dlen; i++ {
		session := types.Session + mylog.Str(i+1)
		ret := make(map[string]string)
		for j := 0; j < slen; j++ {
			ret[types.AvailableCapacity] = mylog.Str(data.Center[i].Session[j].AvailableCapacity)
			ret[types.MinAge] = mylog.Str(data.Center[i].Session[j].MinAge)
			ret[types.Vaccine] = data.Center[i].Session[j].Vaccine
			ret[types.Name] = data.Center[i].Name
			if data.Center[i].Session[j].AvailableCapacity > 0 {
				counter[types.Available] = 1
			}
			for k := 0; k < slotlen; k++ {
				_slot := types.Slot + mylog.Str(k+1)
				ret[_slot] = data.Center[i].Session[j].Slots[k]
			}
			j++
		}
		final[session] = ret
	}
	counter[types.SessionCount] = slen
	counter[types.SlotCount] = slotlen
	types.FinalMsg = final
	return final, counter
}

// simple http.Get + decode json response
func Fetch(url string) (types.Meta, error) {
	var CoMeta types.Meta
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println(err, resp.StatusCode)
		return types.Meta{}, fmt.Errorf("fetch v1: unable to fetch URL: %v", err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&CoMeta)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println(err, resp.StatusCode)
		return types.Meta{}, fmt.Errorf("fetch v1: unable to decode json response: %v", err)

	}
	return CoMeta, nil
}

// custom http.Get for changing header + decode json response
func FetchV2(url string) (types.Meta, error) {
	CoMeta := &types.Meta{}
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	log.Println(url)
	if err != nil {
		return types.Meta{}, fmt.Errorf("fetch v2: unable to fetch url: %v", err)
	}

	// faking browser
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1 Safari/605.1.15")
	resp, err := client.Do(request)

	if resp == nil {
		return types.Meta{}, err
	}

	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println(err, resp.StatusCode)
		return types.Meta{}, fmt.Errorf("fetch v2: unable to fetch url: %v", err)
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&CoMeta)

	if err != nil {
		return types.Meta{}, fmt.Errorf("fetch v2: unable to decode json response %v", err)
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
	base, err := url.Parse(types.URL)
	if err != nil {
		return "", fmt.Errorf("buildquery: unable to parse url: %v", err)
	}
	base.Path += types.URLPATH
	params := url.Values{}
	params.Add(types.DATEQUERY, date)
	params.Add(types.PINQUERY, types.PINCODE)
	base.RawQuery = params.Encode()
	url := base.String()
	return url, nil
}
