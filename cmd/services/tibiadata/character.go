package tibiadata

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateCharacterResponse(player string) (*CharacterResponse, error) {
	response, err := http.Get("https://api.tibiadata.com/v3/character/" + player)
	if err != nil {
		log.Println("[CreateCharacterResponse] Error calling TibiaData API #Character ", err)
		return nil, err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("[CreateCharacterResponse] Error calling CharacterResponse io.ReadAll ", err)
		return nil, err
	}

	var responseObject CharacterResponse
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return nil, err
	}

	return &responseObject, nil
}
