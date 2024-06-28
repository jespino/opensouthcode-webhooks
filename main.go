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
				locat, err := time.LoadLocation("Europe/Madrid")
				if err != nil {
					log.Fatal(err)
				}
				eventDateTime, err := time.ParseInLocation("2006-01-02T15:04:05Z", strings.Replace(event.Date[0:16], " ", "T", 1)+":00Z", locat)
				if err != nil {
					log.Printf("Error al parsear fecha de evento: %v", err)
					continue
				}
				// This is a trick to change the date to today always
				// eventDateTimeToday := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), eventDateTime.Hour(), eventDateTime.Minute(), 0, 0, time.Now().Location())
				if eventDateTime.After(time.Now().Add(14*time.Minute).Add(10*time.Second)) && eventDateTime.Before(time.Now().Add(15*time.Minute).Add(10*time.Second)) {
					fmt.Println("AVISO: ", event.Title, "DATE: ", event.Date)
					sendEvent(event, "### Aviso de charla dentro de 15 minutos:")
				}
			}
		}
	}
	log.Printf("Hora objetivo: ", time.Now().Add(15*time.Minute).Add(10*time.Second))
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
