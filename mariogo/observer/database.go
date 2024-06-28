package observer

import (
	"fmt"
	"jkli/mariogo/mariogo"
	"jkli/mariogo/mariogo/analyzer"
	"time"

	"gorm.io/gorm/clause"
)

type Database struct {
	game            *mariogo.Game
	gameInitialized bool
}

func NewDatabase(player int) *Database {
	return &Database{
		gameInitialized: false,
	}
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

func (d *Database) Finish() {
	d.game.Finished = true
	mariogo.DB.Save(&d.game)
	d.game = nil
	d.gameInitialized = false
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

	var roundPlacements []mariogo.Placement

	for i, position := range placements {
		if position == 0 {
			break
		}
		player := g.game.Players[i]
		roundPlacement := mariogo.Placement{
			Round:    round,
			PlayerID: player.ID,
			Position: position,
		}
		roundPlacements = append(roundPlacements, roundPlacement)
	}

	mariogo.DB.Create(&roundPlacements)
	g.updateGame()
}

func (d *Database) PlayerCount(count int) {
	if !d.gameInitialized {
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

	if !d.gameInitialized && to == analyzer.Racing {
		d.CreateGame()
	}

	if to == analyzer.EndResults {
		d.Finish()
	}

}

func (d *Database) CreateGame() {
	d.gameInitialized = true

	game := mariogo.Game{
		Finished: false,
	}

	mariogo.DB.Create(&game)
	mariogo.DB.First(&d.game)
}

func (d *Database) RoundFinished(player int, round int, time time.Duration) {

	playerModel := d.game.GetPlayerByPosition(player)
	roundModel := d.game.GetCurrentRound()

	if roundModel == nil || playerModel == nil {
		return
	}

	roundTime := mariogo.RoundTime{
		PlayerID: playerModel.ID,
		RoundID:  roundModel.ID,
		Time:     uint(time.Milliseconds()),
	}

	mariogo.DB.Save(&roundTime)
	d.updateGame()
}

func (d *Database) PlayerFinishedRace(player int, time time.Duration) {
}

func (d *Database) PlacementsChanged(old [4]int, new [4]int, roundTime time.Duration) {

	if len(d.game.Rounds) == 0 {
		return
	}

	var p1, p2, p3, p4 *uint
	round := d.game.Rounds[len(d.game.Rounds)-1]

	if new[0] != 0 {
		place := uint(new[0])
		p1 = &place
	}

	if new[1] != 0 {
		place := uint(new[1])
		p2 = &place
	}

	if new[2] != 0 {
		place := uint(new[2])
		p3 = &place
	}

	if new[3] != 0 {
		place := uint(new[3])
		p4 = &place
	}

	log := mariogo.PlacementChangeLog{
		RoundID: round.ID,
		Time:    uint(roundTime.Milliseconds()),
		Player1: p1,
		Player2: p2,
		Player3: p3,
		Player4: p4,
	}

	mariogo.DB.Create(&log)

	d.updateGame()
}
