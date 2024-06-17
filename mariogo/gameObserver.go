package mariogo

import "gorm.io/gorm/clause"

type GameObserver struct {
	game        Game
	playerCount int
}

func NewGameObserver(player int) *GameObserver {

	game := Game{
		Finished: false,
	}

	DB.Create(&game)
	DB.First(&game)

	return &GameObserver{
		game:        game,
		playerCount: player,
	}
}

func (g *GameObserver) updateGame() {
	DB.Preload(clause.Associations).Preload("Players.Person.DefaultCharacter").First(&g.game)
}

func (g *GameObserver) Abort(message string) {
	g.game.Finished = true
	msg := "Game aborted"

	if message != "" {
		msg += ": " + message
	}

	g.game.Error = &msg

	DB.Save(&g.game)
	g.updateGame()
}

func (g *GameObserver) Finish() {
	g.game.Finished = true
	DB.Save(&g.game)
	g.updateGame()
}

func (g *GameObserver) NewRound(index int, trackName *string) {
	round := Round{
		Index:     index,
		TrackName: trackName,
		Game:      g.game,
	}

	DB.Create(&round)
	g.updateGame()
}

func (g *GameObserver) RoundResults(placements [4]int) {

	round := g.game.Rounds[len(g.game.Rounds)-1]

	var roundPlacements []RoundPlacement

	for i, position := range placements {
		if position == 0 {
			break
		}
		player := g.game.Players[i]
		roundPlacement := RoundPlacement{
			Round:    round,
			PlayerID: player.ID,
			Position: position,
		}
		roundPlacements = append(roundPlacements, roundPlacement)
	}

	DB.Create(&roundPlacements)
	g.updateGame()
}

func (g *GameObserver) InterimResults(placements [4]int) {

	if len(g.game.Placements) > 0 {
		for i, placement := range g.game.Placements {
			// TODO: make sur player is right player
			placement.Position = placements[i]
			DB.Save(&placement)
		}
		g.updateGame()
		return
	}

	var interimPlacements []Placement

	for i, position := range placements {
		if position == 0 {
			break
		}
		player := g.game.Players[i]
		placement := Placement{
			GameID:   g.game.ID,
			PlayerID: player.ID,
			Position: position,
		}
		interimPlacements = append(interimPlacements, placement)
	}

	DB.Create(&interimPlacements)
	g.updateGame()
}

func (g *GameObserver) RegisterPlayer(player int, name string) {

	playerModel := Player{
		GameID:       g.game.ID,
		Number:       player,
		FallbackName: &name,
	}

	DB.Save(&playerModel)
	g.updateGame()
}

func (g *GameObserver) GetRegisteredPlayer() int {
	return len(g.game.Players)
}
