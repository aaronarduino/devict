package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type ResultsTwo struct {
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
}

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

func main() {
	dat_1 := GetMeetupResults("https://api.meetup.com/2/events?offset=0&format=json&limited_events=False&group_urlname=WWCWichita&photo-host=public&page=20&fields=&order=time&status=upcoming&desc=false&sig_id=73273692&sig=4111c5adf6695f954bd7ae7dfd86896970b451f6")
	fmt.Println(reflect.TypeOf(dat_1.Results[0]))
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
