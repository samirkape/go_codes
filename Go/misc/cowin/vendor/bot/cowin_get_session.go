package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func GetSessionInfo() {
	url := BuildQuery()
	data, err := FetchV2(url)
	if err != nil {
		Println("Unable to parse URL", os.Stdout)
		os.Exit(-1)
	}
	if reflect.ValueOf(data).IsZero() {
		Println("No Data Yet", os.Stdout)
	} else {
		// implement SMS logic here
		final, counter := Filter(data)
		SendMessage(final, counter)
		fmt.Println(final)
	}
}

func getTBot() (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(Token)
	bot.Debug = true

	if err != nil {
		Println("Error getting bot", os.Stdout)
		return nil, err
	}
	return bot, err
}

func SendMessage(rmsg map[string]map[string]string, counter map[string]int) {
	bot, _ := getTBot()
	var write strings.Builder
	i := 1
	if counter["Available"] == 0 {
		return
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
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)

	if err != nil {
		Println("Error Sending Message", os.Stdout)
		return
	}
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
	return final, counter
}

func str(in interface{}) string {
	return fmt.Sprint(in)
}

func Fetch(url string) []string {
	var CoMeta Meta
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		Println("Unable to Fetch URL", os.Stdout)
		return nil
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&CoMeta)
	if err != nil {
		Println("Unable to Load Response", os.Stdout)
		return nil
	}
	return nil //TBI
}

func FetchV2(url string) (Meta, error) {
	CoMeta := &Meta{}
	//eg := Example{}
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// faking browser
	request.Header.Set("User-Agent", "User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.76 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil || resp.StatusCode != http.StatusOK {
		Println("Fetch v2: Unable to Fetch URL", os.Stdout)
		return Meta{}, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&CoMeta)
	//body, err := ioutil.ReadAll(resp.Body)
	//json.Unmarshal(body, &CoMeta)
	if err != nil {
		Println("Unable to Load Response", os.Stdout)
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
		Println("Unable to parse URL", os.Stdout)
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

func Println(str string, where io.Writer) {
	fmt.Fprintf(where, "%s\n", str)
}
