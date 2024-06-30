package mariogo

import (
	"time"

	"github.com/corona10/goimagehash"
)

type PlayerPlacement struct {
	Position     int
	IsBot        bool
	PlayerNumber int
	IconHash     *goimagehash.ImageHash
}

type Observer interface {
	StateChange(from int, to int)
	PlayerCount(count int)
	PlayerName(player int, name string)
	NewRound(name string)
	PlacementsChanged(old [4]int, new [4]int, roundTime time.Duration)
	RoundFinished(player int, round int, time time.Duration)
	PlayerFinishedRace(player int, time time.Duration)
	RoundResults(placements []PlayerPlacement)
	Abort(message string)
}
