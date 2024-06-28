package main

import "github.com/mattermost/mattermost/server/public/model"

const (
	IncommingWebhookURL = "https://opensouthcode.cloud.mattermost.com/hooks/<ID>"
	Channel             = "programaschedule"
)

func sendEvent(event Event, message string) {
	authorName := ""
	authorLink := ""
	for _, person := range event.Persons {
		if authorLink == "" {
			authorLink = "https://www.opensouthcode.org/users/" + person.Id
		}
		if authorName != "" {
			authorName += ", "
		}
		authorName += person.Name
	}
	attachment := model.SlackAttachment{
		AuthorName: authorName,
		AuthorLink: authorLink,
		Title:      event.Title,
		TitleLink:  "https://www.opensouthcode.org/conferences/opensouthcode2024/program/proposals/" + event.Id,
		Text:       event.Abstract,
		Fields: []*model.SlackAttachmentField{
			{
				Title: "Date",
				Value: event.Date[0:10],
				Short: model.SlackCompatibleBool(true),
			},
			{
				Title: "Start",
				Value: event.Start,
				Short: model.SlackCompatibleBool(true),
			},
			{
				Title: "Room",
				Value: event.Room,
				Short: model.SlackCompatibleBool(true),
			},
			{
				Title: "Duration",
				Value: event.Duration,
				Short: model.SlackCompatibleBool(true),
			},
			{
				Title: "Type",
				Value: event.Type,
				Short: model.SlackCompatibleBool(true),
			},
			{
				Title: "Language",
				Value: event.Language,
				Short: model.SlackCompatibleBool(true),
			},
		},
	}
	sendMessage(IncommingWebhookURL, Channel, message, []*model.SlackAttachment{&attachment})
}
