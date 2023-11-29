package bot

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"ts3-claimedBot/cmd/models"
	"ts3-claimedBot/cmd/services/tibiadata"
	"ts3-claimedBot/cmd/teamspeak"
)

func (app *Bot) HandleMemberDeath() {

	for {
		log.Println("[HandleMemberDeath] Verificando deaths...")

		deathsFound := 0

		var notificationsToSend []string

		for _, guildName := range app.GetGuildsEnemies() {

			var membersGuildDatabaseIds []uint
			membersGuildDatabaseIds = app.Repository.GetIdsMembersInsideGuildFromDatabase(guildName)

			for _, memberId := range membersGuildDatabaseIds {
				var character models.Character
				result := app.Repository.Db.First(&character, "id = ?", memberId)

				if result.RowsAffected == 0 {
					continue
				}

				playerName := character.Name
				playerId := character.ID

				responseObject, err := tibiadata.CreateCharacterResponse(playerName)
				if err != nil {
					log.Println("[HandleMemberDeath] Error loading character information: ", err)
					continue
				}

				for _, siteDeath := range responseObject.Characters.Deaths {

					deathId := app.Repository.GetDeathIdFromIdCharacterAndTime(playerId, siteDeath.Time)

					if deathId > 0 {
						continue
					}

					death := models.Death{
						IdCharacter: playerId,
						Time:        siteDeath.Time,
						Reason:      siteDeath.Reason,
						Level:       strconv.Itoa(siteDeath.Level),
					}

					result = app.Repository.Db.Create(&death)

					deathsFound++

					if result.Error != nil || result.RowsAffected == 0 {
						log.Println("[Setup] Error creating death. ", result.Error)
						continue
					}

					log.Println(fmt.Sprintf("[HandleMemberDeath] Sending notification %s...", playerName))
					message := fmt.Sprintf("[%s] %s - %s", death.Time, playerName, death.Reason)
					notificationsToSend = append(notificationsToSend, message)
				}
			}
		}

		for _, message := range notificationsToSend {
			log.Println(fmt.Sprintf("--- Sending... %s", message))
			app.NotifyMemberDeath(message)
		}

		log.Println(fmt.Sprintf("[HandleMemberDeath] Deaths %v found", deathsFound))
		time.Sleep(TIMER_RECHECK_MEMBER)
	}
}

func (app *Bot) NotifyMemberDeath(msg string) {

	message := teamspeak.Message{
		Client:  app.Client,
		Msg:     msg,
		Channel: 1,
	}

	err := message.SendMessageChannel()
	if err != nil {
		log.Println("[NotifyMemberDeath] Error sending notification: ", err)
	}
}
