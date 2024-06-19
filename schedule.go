package main

import (
	"encoding/xml"
	"io"
	"net/http"
)

const (
	scheduleXMLURL = "https://www.opensouthcode.org/conferences/opensouthcode2024/schedule.xml"
)

type Schedule struct {
	Conference Conference `xml:"conference"`
	Days       []Day      `xml:"day"`
}

type Day struct {
	Date  string `xml:"date,attr"`
	Rooms []Room `xml:"room"`
}

type Room struct {
	Name   string  `xml:"name,attr"`
	Events []Event `xml:"event"`
}

type Event struct {
	Id          string    `xml:"id,attr"`
	Guid        string    `xml:"guid,attr"`
	Date        string    `xml:"date"`
	Start       string    `xml:"start"`
	Duration    string    `xml:"duration"`
	Room        string    `xml:"room"`
	Type        string    `xml:"type"`
	Language    string    `xml:"language"`
	Slug        string    `xml:"slug"`
	Title       string    `xml:"title"`
	Subtitle    string    `xml:"subtitle"`
	Abstract    string    `xml:"abstract"`
	Track       string    `xml:"track"`
	Description string    `xml:"description"`
	Persons     []Person  `xml:"persons>person"`
	Recording   Recording `xml:"recording"`
}

type Recording struct {
	License string `xml:"license"`
	Optout  string `xml:"optout"`
}

type Person struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Conference struct {
	Acronym          string `xml:"acronym"`
	Title            string `xml:"title"`
	Start            string `xml:"start"`
	End              string `xml:"end"`
	Days             int    `xml:"days"`
	TimeslotDuration string `xml:"timeslot_duration"`
}

func getSchedule() (*Schedule, error) {
	var schedule Schedule
	resp, err := http.Get(scheduleXMLURL)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &schedule)
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}
