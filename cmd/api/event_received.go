package main

import (
	"golang.org/x/exp/slices"
	"log"
	"strings"
	"ts3-claimedBot/cmd/teamspeak"
)

func (app *Config) HandleEventsReceived() {

	log.Println("[HandleEventsReceived] Ready to receive events...")

	for notification := range app.Bot.Client.Notifications() {

		if notification.Type != "textmessage" {
			continue
		}

		user := teamspeak.User{
			Id:   notification.Data["invokerid"],
			Uuid: notification.Data["invokeruid"],
			Name: notification.Data["invokername"],
		}

		msg := notification.Data["msg"]
		splittedMsg := strings.Split(msg, " ")
		command := splittedMsg[0]
		var argument string

		for _, val := range splittedMsg[1:] {
			argument = argument + " " + val
		}

		argument = strings.Trim(argument, " ")

		if !slices.Contains(app.Bot.GetAllowedCommands(), command) {
			log.Println("[HandleEventsReceived] Command not allowed!")
			continue
		}

		switch command {
		case "!info":
			go app.Bot.Info(argument)
			break
		case "!register":
			go app.Bot.Register(user)
			break
		case "!check":
			go app.Bot.Check(user)
			break
		default:
			continue
		}
	}
}
