package tibiadata

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateGuildResponse(guildName string) (*GuildResponse, error) {
	response, err := http.Get(fmt.Sprintf("https://api.tibiadata.com/v3/guild/%s", guildName))
	if err != nil {
		log.Fatal("[CreateGuildResponse] Error calling TibiaData API #Guild ", err)
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("[CreateGuildResponse]  calling CreateGuildResponse io.ReadAll ", err)
		return nil, err
	}

	var responseObject GuildResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}

	return &responseObject, nil
}
