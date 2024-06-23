package mariogo

import (
	"gorm.io/gorm"
)

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
	Index               int
	TrackName           string
	Placements          []RoundPlacement
	PlacementChangeLogs []PlacementChangeLog
	Game                Game
	GameID              *uint
}

type PlacementChangeLog struct {
	gorm.Model
	Round   Round
	RoundID uint
	Time    uint
	Player1 *uint
	Player2 *uint
	Player3 *uint
	Player4 *uint
}

type RoundTime struct {
	gorm.Model
	Round    Round
	RoundID  uint
	Player   Player
	PlayerID uint
	Time     uint
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
	Number       int
	Game         Game
	GameID       uint
	FallbackName *string
	Character    *Character
	CharacterID  *uint
	Person       *Person
	PersonID     *uint
}

type Character struct {
	gorm.Model
	Name    string
	Weight  string
	Image   string
	Persons []Person
}

type Person struct {
	gorm.Model
	Name        uint
	Character   *Character
	CharacterID *uint
}

func (g *Game) GetPlayerByPosition(position int) *Player {
	for _, player := range g.Players {
		if player.Number == position {
			return &player
		}
	}
	return nil
}

func (g *Game) GetCurrentRound() *Round {
	if len(g.Rounds) == 0 {
		return nil
	}
	return &g.Rounds[len(g.Rounds)-1]
}
