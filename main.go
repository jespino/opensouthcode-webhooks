package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

func notifyEvents(schedule *Schedule) {
	for _, day := range schedule.Days {
		for _, room := range day.Rooms {
			for _, event := range room.Events {
				eventDateTime, err := time.Parse(time.RFC3339, strings.Replace(event.Date[0:16], " ", "T", 1)+":00Z")
				if err != nil {
					log.Printf("Error al parsear fecha de evento: %v", err)
					continue
				}
				// This is a trick to change the date to today always
				// eventDateTimeToday := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), eventDateTime.Hour(), eventDateTime.Minute(), 0, 0, time.Now().Location())
				if eventDateTime.After(time.Now().Add(15*time.Minute).Add(10*time.Second).UTC()) && eventDateTime.Before(time.Now().Add(15*time.Minute).Add(10*time.Second).UTC()) {
					fmt.Println("NOW", time.Now(), "event date", event.Date)
					sendEvent(event, "### Aviso de charla dentro de 15 minutos:")
				}
			}
		}
	}
}

func main() {
	schedule, err := getSchedule()
	if err != nil {
		log.Fatalln(err)
	}
	notifyEvents(schedule)
	// event := schedule.Days[0].Rooms[0].Events[0]
	// sendEvent(event, "### Aviso de charla dentro de 15 minutos:")

	c := cron.New()
	_, err = c.AddFunc("* * * * *", func() {
		notifyEvents(schedule)
	})
	if err != nil {
		log.Fatalln(err)
	}

	c.Start()
	defer c.Stop()

	for {
		time.Sleep(1 * time.Hour)
	}
}
