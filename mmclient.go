package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/mattermost/mattermost/server/public/model"
)

func sendMessage(webhookUrl string, channelName string, message string, attachments []*model.SlackAttachment) error {
	request := model.IncomingWebhookRequest{
		Text:     message,
		Username: "OpenSouthCode Bot",
		// IconURL:     "https://yt3.ggpht.com/bb3S4C-s4BAYtXU7DMIecV5GXfnHDa1OY8STxGiZSvtr6vdVU9VQ2bVASxCTvwoZD57HRYLUkg=s88-c-k-c0x00ffffff-no-rj",
		IconURL:     "https://pbs.twimg.com/profile_images/712755802961014787/D2yArhYE_400x400.jpg",
		ChannelName: channelName,
		Props:       map[string]interface{}{},
		Attachments: attachments,
		Type:        "",
		IconEmoji:   "",
	}

	data, err := json.Marshal(request)
	if err != nil {
		return err
	}

	_, err = http.Post(webhookUrl, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return nil
}
