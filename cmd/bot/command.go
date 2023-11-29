package bot

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"ts3-claimedBot/cmd/services/tibiadata"
	"ts3-claimedBot/cmd/teamspeak"
)

func (b *Bot) GetAllowedCommands() []string {
	return []string{
		"!info",
		"!register",
		"!check",
	}
}

func (b *Bot) Info(argument string) {

	response, err := http.Get("https://api.tibiadata.com/v3/character/" + argument)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	var responseObject tibiadata.CharacterResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal(err)
		return
	}

	character := responseObject.Characters.Character

	var msg string

	if character.Level == 0 {
		msg = fmt.Sprintf("Character not exist.")
	} else {
		msg = fmt.Sprintf("Name: %s - Level: %v - Vocation: %s", character.Name, character.Level, character.Vocation)
	}

	/*
		for _, death := range responseObject.Characters.Deaths {
			deathMessage := fmt.Sprintf("[%s] %s", death.Time, death.Reason)
			fmt.Println(deathMessage)
		}
	*/

	message := teamspeak.Message{
		Client:  b.Client,
		Msg:     msg,
		Channel: 1,
	}
	err = message.SendMessageChannel()

	if err != nil {
		log.Fatal(err)
	}
}

func (b *Bot) Register(user teamspeak.User) {

	uuid, _ := uuid.NewUUID()

	msg := fmt.Sprintf("Add in the character comment: %s", uuid.String())

	message := teamspeak.Message{
		Client:  b.Client,
		User:    &user,
		Msg:     msg,
		Private: true,
	}

	err := message.SendMessageToUser()
	if err != nil {
		log.Fatal(err)
	}

	message.Msg = fmt.Sprintf("After, confirm with !check")
	err = message.SendMessageToUser()
	if err != nil {
		log.Fatal(err)
	}
}

func (b *Bot) Check(user teamspeak.User) {

	msg := fmt.Sprintf("Confirmed!")

	message := teamspeak.Message{
		Client:  b.Client,
		User:    &user,
		Msg:     msg,
		Private: true,
	}

	err := message.SendMessageToUser()
	if err != nil {
		log.Fatal(err)
	}
}
