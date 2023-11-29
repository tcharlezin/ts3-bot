package bot

import (
	"github.com/multiplay/go-ts3"
	"log"
	"strconv"
	"time"
	"ts3-claimedBot/cmd/models"
	"ts3-claimedBot/cmd/services/tibiadata"
	"ts3-claimedBot/cmd/teamspeak"
)

const TIMER_RECHECK_MEMBER = 10 * time.Second

type Bot struct {
	Client               *ts3.Client
	Repository           models.Repository
	World                string
	EnemiesGuildResponse []*tibiadata.GuildResponse
}

func (app *Bot) Notify(msg string) {

	message := teamspeak.Message{
		Client:  app.Client,
		Msg:     msg,
		Channel: 1,
	}

	err := message.SendMessageChannel()
	if err != nil {
		log.Println("[Notify] Error sending notification: ", err)
	}
}

func (bot *Bot) GetGuildsEnemies() []string {
	return []string{
		"Rangers",
		"Rangers Academy",
	}
}

func (bot *Bot) Setup() {

	bot.World = "Belobra"

	for _, guildName := range bot.GetGuildsEnemies() {
		guildResponse, err := tibiadata.CreateGuildResponse(guildName)
		if err != nil {
			log.Fatal("[Setup] ", err)
		}

		bot.EnemiesGuildResponse = append(bot.EnemiesGuildResponse, guildResponse)
	}

	for _, guildResponse := range bot.EnemiesGuildResponse {

		guildName := guildResponse.Guilds.Guild.Name

		var guild models.Guild
		result := bot.Repository.Db.First(&guild, "name = ?", guildName)

		if result.RowsAffected == 0 {
			guild.Name = guildName
			result = bot.Repository.Db.Create(&guild)

			if result.Error != nil || result.RowsAffected == 0 {
				log.Fatal("[Setup] Error creating guild. ", result.Error)
			}
		}

		for _, member := range guildResponse.Guilds.Guild.Members {

			var character models.Character
			result := bot.Repository.Db.First(&character, "name = ?", member.Name)

			if result.RowsAffected == 0 {

				responseObject, err := tibiadata.CreateCharacterResponse(member.Name)
				if err != nil {
					log.Println("[Setup] Error loading character information: ", err)
					continue
				}

				foundByFormerName := false
				for _, formerName := range responseObject.Characters.Character.FormerNames {
					result := bot.Repository.Db.First(&character, "name = ?", formerName)
					if result.RowsAffected > 0 {
						foundByFormerName = true
						break
					}
				}

				if !foundByFormerName {

					character.Name = member.Name
					character.Vocation = member.Vocation
					character.World = bot.World
					character.Level = member.Level
					character.FormerNames = ""
					character.FormerWorlds = ""

					result = bot.Repository.Db.Create(&character)

					if result.Error != nil || result.RowsAffected == 0 {
						log.Fatal("[Setup] Error creating character. ", result.Error)
					}
				}

				for _, siteDeath := range responseObject.Characters.Deaths {

					deathId := bot.Repository.GetDeathIdFromIdCharacterAndTime(character.ID, siteDeath.Time)

					if deathId > 0 {
						continue
					}

					death := models.Death{
						IdCharacter: character.ID,
						Time:        siteDeath.Time,
						Reason:      siteDeath.Reason,
						Level:       strconv.Itoa(siteDeath.Level),
					}

					result = bot.Repository.Db.Create(&death)

					if result.Error != nil || result.RowsAffected == 0 {
						log.Fatal("[Setup] Error creating death. ", result.Error)
					}
				}
			}

			var characterInGuild models.CharacterGuild
			result = bot.Repository.Db.Where("id_character = ? AND id_guild >= ?", character.ID, guild.ID).Find(&characterInGuild)

			if result.RowsAffected == 0 {
				characterInGuild.IdCharacter = character.ID
				characterInGuild.IdGuild = guild.ID
				characterInGuild.JoinedAt = member.Joined

				result = bot.Repository.Db.Create(&characterInGuild)

				if result.Error != nil || result.RowsAffected == 0 {
					log.Fatal("[Setup] Error adding character in guild. ", result.Error)
				}
			}
		}
	}
}
