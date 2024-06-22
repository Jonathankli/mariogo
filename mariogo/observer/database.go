package observer

import (
	"fmt"
	"jkli/mariogo/mariogo"
	"jkli/mariogo/mariogo/analyzer"
	"time"

	"gorm.io/gorm/clause"
)

type Database struct {
	game *mariogo.Game
}

func NewDatabase(player int) *Database {
	return &Database{}
}

func (g *Database) updateGame() {
	mariogo.DB.Preload(clause.Associations).Preload("Players.Person.DefaultCharacter").First(&g.game)
}

func (g *Database) Abort(message string) {
	g.game.Finished = true
	msg := "Game aborted"

	if message != "" {
		msg += ": " + message
	}

	g.game.Error = &msg

	mariogo.DB.Save(&g.game)
	g.updateGame()
}

func (g *Database) Finish() {
	g.game.Finished = true
	mariogo.DB.Save(&g.game)
	g.updateGame()
}

func (g *Database) NewRound(trackName string) {
	round := mariogo.Round{
		Index:     len(g.game.Rounds) + 1,
		TrackName: trackName,
		Game:      *g.game,
	}

	mariogo.DB.Create(&round)
	g.updateGame()
}

func (g *Database) RoundResults(placements [4]int) {

	round := g.game.Rounds[len(g.game.Rounds)-1]

	var roundPlacements []mariogo.RoundPlacement

	for i, position := range placements {
		if position == 0 {
			break
		}
		player := g.game.Players[i]
		roundPlacement := mariogo.RoundPlacement{
			Round:    round,
			PlayerID: player.ID,
			Position: position,
		}
		roundPlacements = append(roundPlacements, roundPlacement)
	}

	mariogo.DB.Create(&roundPlacements)
	g.updateGame()
}

func (g *Database) InterimResults(placements [4]int) {

	if len(g.game.Placements) > 0 {
		for i, placement := range g.game.Placements {
			// TODO: make sur player is right player
			placement.Position = placements[i]
			mariogo.DB.Save(&placement)
		}
		g.updateGame()
		return
	}

	var interimPlacements []mariogo.Placement

	for i, position := range placements {
		if position == 0 {
			break
		}
		player := g.game.Players[i]
		placement := mariogo.Placement{
			GameID:   g.game.ID,
			PlayerID: player.ID,
			Position: position,
		}
		interimPlacements = append(interimPlacements, placement)
	}

	mariogo.DB.Create(&interimPlacements)
	g.updateGame()
}

func (d *Database) PlayerCount(count int) {
	if d.game == nil {
		d.CreateGame()
	}

	var players []mariogo.Player
	for i := 0; i < count; i++ {
		name := fmt.Sprintf("Player %d", i+1)
		players = append(players, mariogo.Player{
			GameID:       d.game.ID,
			FallbackName: &name,
			Number:       i + 1,
		})
	}
	mariogo.DB.Create(&players)
	d.updateGame()
}

func (d *Database) PlayerName(player int, name string) {

	for _, playerModel := range d.game.Players {
		if playerModel.Number == player {
			playerModel.FallbackName = &name
			mariogo.DB.Save(&playerModel)
			break
		}
	}

	d.updateGame()
}

func (d *Database) StateChange(from int, to int) {

	if d.game == nil && to == analyzer.Racing {
		d.CreateGame()
	}

	if d.game != nil && to == analyzer.EndResults {
		d.game = nil
	}

}

func (d *Database) CreateGame() {
	game := mariogo.Game{
		Finished: false,
	}

	mariogo.DB.Create(&game)
	mariogo.DB.First(&d.game)
}

func (d *Database) RoundFinished(player int, round int, time time.Duration, finished bool) {
}

func (d *Database) PlacementsChanged(places [4]int) {
}
