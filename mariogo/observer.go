package mariogo

type Observer interface {
	StateChange(from int, to int)
	PlayerCount(count int)
	PlayerName(player int, name string)
	NewRound(name string)
	RoundResults(placements [4]int)
	InterimResults(placements [4]int)
	Abort(message string)
}
