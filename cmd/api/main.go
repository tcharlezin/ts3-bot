package main

import (
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/multiplay/go-ts3"
	"sync"
	"ts3-claimedBot/cmd/bot"
	"ts3-claimedBot/cmd/setup"
)

type Config struct {
	Client *ts3.Client
	Bot    *bot.Bot
}

func main() {

	app := Config{
		Client: setup.SetupTeamspeak(),
		Bot:    setup.SetupBot(),
	}

	app.Bot.Client = app.Client

	var wg sync.WaitGroup
	wg.Add(1)

	go app.HandleEventsReceived()
	go app.Bot.HandleMembers()
	go app.Bot.HandleMemberDeath()

	wg.Wait()
}
