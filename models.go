package main

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Rounds     []Round
	Players    []Player
	Finished   bool
	Error      *string
	Placements []Placement
}

type Round struct {
	gorm.Model
	Index      int
	TrackName  *string
	Placements []RoundPlacement
	Game       Game
	GameID     *uint
}

type RoundPlacement struct {
	gorm.Model
	Round    Round
	RoundID  uint
	Player   Player
	PlayerID uint
	Position int
}

type Placement struct {
	gorm.Model
	Game     Game
	GameID   uint
	Player   Player
	PlayerID uint
	Position int
}

type Player struct {
	gorm.Model
	Number   int
	Name     string
	Game     Game
	GameID   uint
	Person   *Person
	PersonID *uint
}

type Character struct {
	gorm.Model
	Name    string
	Image   string
	Persons []Person
}

type Person struct {
	gorm.Model
	Name        uint
	Character   *Character
	CharacterID *uint
}
