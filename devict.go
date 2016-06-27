package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type v struct {
	country                string
	localized_country_name string
	city                   string
	address_1              string
	name                   string
	lat                    float64
	lon                    float64
	id                     int
	state                  string
	repinned               bool
}

type g struct {
	join_mode string
	created   int
	name      string
	group_lon float64
	id        int
	urlname   string
	group_lat float64
	who       string
}

type res struct {
	utc_offset       int
	venue            v
	headcount        int
	visibility       string
	waitlist_count   int
	created          int
	maybe_rsvp_count int
	description      string
	event_url        string
	yes_rsvp_count   int
	duration         int
	name             string
	id               string
	time             int
	updated          string
	group            g
	status           string
}

func main() {
	response, err := http.Get("https://api.meetup.com/2/events?offset=0&format=json&limited_events=False&group_urlname=WWCWichita&photo-host=public&page=20&fields=&order=time&status=upcoming&desc=false&sig_id=73273692&sig=4111c5adf6695f954bd7ae7dfd86896970b451f6")
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		var dat map[string]interface{}
		var list []res
		body, _ := ioutil.ReadAll(response.Body)

		if err := json.Unmarshal(body, &dat); err != nil {
			panic(err)
		}
		for _, thing := range dat["results"] {
			list = append(list, thing.(res))
		}
		list = dat["results"].([]res)
		fmt.Println(list[0])
	}
}
