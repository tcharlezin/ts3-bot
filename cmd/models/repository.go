package models

import (
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func (repo *Repository) GetDeathIdFromIdCharacterAndTime(idCharacter uint, time string) uint {
	var death Death
	result := repo.Db.First(&death, "id_character = ? AND time = ?", idCharacter, time)

	if result.RowsAffected == 0 {
		return 0
	}

	return death.ID
}

func (repo *Repository) GetIdsMembersInsideGuildFromDatabase(guildName string) []uint {
	var guild Guild
	result := repo.Db.First(&guild, "name = ?", guildName)

	if result.RowsAffected == 0 {
		return nil
	}

	var charactersInGuild []CharacterGuild
	result = repo.Db.Where("id_guild IN (?)", guild.ID).Find(&charactersInGuild)

	if result.RowsAffected == 0 {
		return nil
	}

	var idsMembersDatabase []uint
	for _, characterInGuild := range charactersInGuild {
		idsMembersDatabase = append(idsMembersDatabase, characterInGuild.IdCharacter)
	}

	return idsMembersDatabase
}

func (repo *Repository) GetIdCharacterFromName(name string) uint {
	var character Character
	result := repo.Db.First(&character, "name = ?", name)

	if result.RowsAffected == 0 {
		return 0
	}

	return character.ID
}

func (repo *Repository) SearchCharacterByName(characterName string) *Character {
	var character Character
	result := repo.Db.First(&character, "name = ?", characterName)

	if result.RowsAffected > 0 {
		return &character
	}

	return nil
}

func (repo *Repository) SearchGuildByName(guildName string) *Guild {
	var guild Guild
	result := repo.Db.First(&guild, "name = ?", guildName)

	if result.RowsAffected > 0 {
		return &guild
	}

	return nil
}
