package mariogo

import (
	"sort"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Rounds   []Round
	Players  []Player
	Finished bool
	Error    *string
}

type Round struct {
	gorm.Model
	Index               int
	TrackName           string
	Placements          []Placement
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

type Placement struct {
	gorm.Model
	Round    Round `json:"-"`
	RoundID  uint
	Player   Player `json:"-"`
	PlayerID uint
	Position int
	Points   int `gorm:"-"` // Not stored in the database
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

func (g *Game) GetPlacements() []Placement {
	if len(g.Rounds) == 0 {
		return nil
	}

	placementMap := make(map[uint]*Placement)

	for _, round := range g.Rounds {
		for _, placement := range round.Placements {
			if placementMap[placement.PlayerID] == nil {
				placementMap[placement.PlayerID] = &placement
			} else {
				placementMap[placement.PlayerID].Points += placement.GetPoints()
			}
		}
	}

	placements := make([]Placement, 0, len(placementMap))
	for _, placement := range placementMap {
		placements = append(placements, *placement)
	}

	sort.Slice(placements, func(i, j int) bool {
		return placements[i].Points > placements[j].Points
	})

	return placements
}

func (p *Placement) GetPoints() int {
	switch p.Position {
	case 1:
		return 15
	case 2:
		return 12
	case 3:
		return 10
	case 4:
		return 9
	case 5:
		return 8
	case 6:
		return 7
	case 7:
		return 6
	case 8:
		return 5
	case 9:
		return 4
	case 10:
		return 3
	case 11:
		return 2
	case 12:
		return 1
	}
	return 0
}
