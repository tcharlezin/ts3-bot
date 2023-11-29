package bot

import (
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"time"
	"ts3-claimedBot/cmd/models"
	"ts3-claimedBot/cmd/services/tibiadata"
	"ts3-claimedBot/cmd/teamspeak"
)

func (bot *Bot) HandleMembers() {

	for {
		log.Println("[HandleMembers] Verificando players...")

		for _, guildName := range bot.GetGuildsEnemies() {

			var membersGuildDatabaseIds []uint
			membersGuildDatabaseIds = bot.Repository.GetIdsMembersInsideGuildFromDatabase(guildName)

			responseObject, err := tibiadata.CreateGuildResponse(guildName)
			if err != nil {
				log.Println("Error verifying player: ", err)
				continue
			}

			var memberIdBasedInSite []uint
			for _, member := range responseObject.Guilds.Guild.Members {
				idMember := bot.Repository.GetIdCharacterFromName(member.Name)
				if idMember > 0 {
					memberIdBasedInSite = append(memberIdBasedInSite, idMember)
				}
			}

			for _, idCharacterFromSite := range memberIdBasedInSite {
				if !slices.Contains(membersGuildDatabaseIds, idCharacterFromSite) {
					bot.NotifyMemberLeaveGuild(idCharacterFromSite, guildName)
				}
			}

			// TODO: Se nao esta no banco de dados, verificar data de join para alertar novo membro
		}

		time.Sleep(TIMER_RECHECK_MEMBER)
	}
}

func (bot *Bot) NotifyMemberLeaveGuild(idCharacter uint, guild string) {
	var character models.Character
	result := bot.Repository.Db.First(&character, "id = ?", idCharacter)

	if result.RowsAffected == 0 {
		return
	}

	message := teamspeak.Message{
		Client:  bot.Client,
		Msg:     fmt.Sprintf("Player %s saiu da guild %s.", character.Name, guild),
		Channel: 1,
	}

	log.Println(message.Msg)

	err := message.SendMessageChannel()
	if err != nil {
		return
	}
}
