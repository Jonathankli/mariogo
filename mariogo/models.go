package mariogo

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Rounds     []Round //siehe type Round struct {}
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
	Game                Game `json:"-"`
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
	Round    Round `json:"-"`
	RoundID  uint
	Player   Player `json:"-"`
	PlayerID uint
	Time     uint
}
type RoundPlacement struct {
	gorm.Model
	Round    Round `json:"-"`
	RoundID  uint
	Player   Player `json:"-"`
	PlayerID uint
	Position int
}

type Placement struct {
	gorm.Model
	Game     Game `json:"-"`
	GameID   uint
	Player   Player `json:"-"`
	PlayerID uint
	Position int
}

type Player struct {
	gorm.Model
	Number       int
	Game         Game `json:"-"`
	GameID       uint
	FallbackName *string
	Character    *Character `json:"-"`
	CharacterID  *uint      `json:"-"`
	Person       *Person    `json:"-"`
	PersonID     *uint      `json:"-"`
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
	Name        string
	Character   *Character
	CharacterID *uint
}

type PlayerUpdateInput struct {
	FallbackName *string `json:"fallback_name"`
	CharacterID  *uint   `json:"character_id"`
	PersonID     *uint   `json:"person_id"`
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
