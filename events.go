package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Result struct {
	Group    string
	Title    string
	Time     int64
	Location string
}

type Results []Result

type MeetupRes struct {
	Results []struct {
		UtcOffset int `json:"utc_offset"`
		Venue     struct {
			Zip                  string  `json:"zip"`
			Country              string  `json:"country"`
			LocalizedCountryName string  `json:"localized_country_name"`
			City                 string  `json:"city"`
			Address1             string  `json:"address_1"`
			Name                 string  `json:"name"`
			Lon                  float64 `json:"lon"`
			ID                   int     `json:"id"`
			State                string  `json:"state"`
			Lat                  float64 `json:"lat"`
			Repinned             bool    `json:"repinned"`
		} `json:"venue,omitempty"`
		Headcount      int    `json:"headcount"`
		Visibility     string `json:"visibility"`
		WaitlistCount  int    `json:"waitlist_count"`
		Created        int64  `json:"created"`
		MaybeRsvpCount int    `json:"maybe_rsvp_count"`
		Description    string `json:"description"`
		EventURL       string `json:"event_url"`
		YesRsvpCount   int    `json:"yes_rsvp_count"`
		Duration       int    `json:"duration"`
		Name           string `json:"name"`
		ID             string `json:"id"`
		Time           int64  `json:"time"`
		Updated        int64  `json:"updated"`
		Group          struct {
			JoinMode string  `json:"join_mode"`
			Created  int64   `json:"created"`
			Name     string  `json:"name"`
			GroupLon float64 `json:"group_lon"`
			ID       int     `json:"id"`
			Urlname  string  `json:"urlname"`
			GroupLat float64 `json:"group_lat"`
			Who      string  `json:"who"`
		} `json:"group"`
		Status string `json:"status"`
	} `json:"results"`
	Meta struct {
		Next        string `json:"next"`
		Method      string `json:"method"`
		TotalCount  int    `json:"total_count"`
		Link        string `json:"link"`
		Count       int    `json:"count"`
		Description string `json:"description"`
		Lon         string `json:"lon"`
		Title       string `json:"title"`
		URL         string `json:"url"`
		ID          string `json:"id"`
		Updated     int64  `json:"updated"`
		Lat         string `json:"lat"`
	} `json:"meta"`
}

const (
	wwcURL         = "https://api.meetup.com/2/events?offset=0&format=json&limited_events=False&group_urlname=WWCWichita&photo-host=public&page=20&fields=&order=time&status=upcoming&desc=false&sig_id=73273692&sig=4111c5adf6695f954bd7ae7dfd86896970b451f6"
	devictURL      = "http://api.meetup.com/2/events?status=upcoming&order=time&limited_events=False&group_urlname=devict&desc=false&offset=0&photo-host=public&format=json&page=20&fields=&sig_id=73273692&sig=9cdd3af6b5a26eb664fe5abab6e5cf7bfaaf090e"
	makeictURL     = "https://api.meetup.com/2/events?offset=0&format=json&limited_events=False&group_urlname=MakeICT&photo-host=public&page=20&fields=&order=time&desc=false&status=upcoming&sig_id=15434981&sig=5da76a33f42c53199e5d7f97a3ed5340f3cc2e61"
	openwichitaURL = "https://api.meetup.com/2/events?offset=0&format=json&limited_events=False&group_urlname=openwichita&photo-host=public&page=20&fields=&order=time&desc=false&status=upcoming&sig_id=15434981&sig=25dca881d2d1fc821fe708f3687c83f451c1b683"
)

func (slice Results) Len() int {
	return len(slice)
}

func (slice Results) Less(i, j int) bool {
	return slice[i].Time < slice[j].Time
}

func (slice Results) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func PrintEvents() {
	w := tablewriter.NewWriter(os.Stdout)
	allRes := GetEvents()
	sort.Sort(allRes)

	for _, event := range allRes[:20] {
		var title, loc, grp string

		// Convert date
		y, m, d := time.Unix(0, event.Time*1000000).Date()
		date := strconv.Itoa(y) + "-" + strconv.Itoa(int(m)) + "-" + strconv.Itoa(d)

		// Trim title of event
		if len(event.Title) > 20 {
			title = event.Title[:19] + "..."
		} else {
			title = event.Title
		}

		// Trim title of location
		if len(event.Location) > 19 {
			loc = event.Location[:18] + "..."
		} else {
			loc = event.Location
		}

		// Shorten Women Who Code Wichita title
		if event.Group == "Women Who Code Wichita" {
			grp = "WWC Wichita"
		} else {
			grp = event.Group
		}

		// Append data to table
		w.Append([]string{date, grp, title, loc})
	}
	w.Render() // Renders table
}

func GetEvents() Results {
	allRes := Results{}
	list := []MeetupRes{}

	wwcData := GetMeetupResults(wwcURL)
	devictURL := GetMeetupResults(devictURL)
	makeictURL := GetMeetupResults(makeictURL)
	openwichitaURL := GetMeetupResults(openwichitaURL)
	list = append(list, wwcData)
	list = append(list, devictURL)
	list = append(list, makeictURL)
	list = append(list, openwichitaURL)

	for _, item := range list {
		for _, res := range item.Results {
			allRes = append(allRes, Result{
				res.Group.Name,
				res.Name,
				res.Time,
				res.Venue.Name,
			})
		}
	}
	return allRes
}

func GetMeetupResults(URL string) MeetupRes {
	var res MeetupRes
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		if err := json.Unmarshal(body, &res); err != nil {
			panic(err)
		}
	}
	return res
}
