package setup

import (
	"ts3-claimedBot/cmd/bot"
	"ts3-claimedBot/cmd/models"
)

func SetupBot() *bot.Bot {

	b := bot.Bot{
		Repository: models.Repository{
			Db: SetupDatabase(),
		},
	}

	return &b
}
