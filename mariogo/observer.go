package mariogo

import "time"

type Observer interface {
	StateChange(from int, to int)
	PlayerCount(count int)
	PlayerName(player int, name string)
	NewRound(name string)
	RoundFinished(player int, round int, time time.Duration, finished bool)
	RoundResults(placements [4]int)
	InterimResults(placements [4]int)
	Abort(message string)
}
