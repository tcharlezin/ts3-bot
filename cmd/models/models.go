package models

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name         string
	Vocation     string
	World        string
	Level        int
	FormerNames  string
	FormerWorlds string
}

type Guild struct {
	gorm.Model
	Name string
}

type Death struct {
	gorm.Model
	IdCharacter uint
	Time        string
	Reason      string
	Level       string
}

type CharacterGuild struct {
	gorm.Model
	IdCharacter uint
	IdGuild     uint
	JoinedAt    string
}
