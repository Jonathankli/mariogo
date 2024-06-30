package analyzer

import (
	"jkli/mariogo/mariogo"
	"jkli/mariogo/mariogo/pixel"
	"time"
)

var RoundPixels = [][]pixel.Pixel{
	pixel.RoundOne,
	pixel.RoundTwo,
	pixel.RoundThree,
	pixel.RoundFour,
	pixel.RoundFive,
	pixel.RoundSix,
	pixel.RoundSeven,
}

var RoundPixels2P = [][]pixel.Pixel{
	pixel.RoundOne2P,
	pixel.RoundTwo2P,
	pixel.RoundThree2P,
	pixel.RoundFour2P,
	pixel.RoundFive2P,
	pixel.RoundSix2P,
	pixel.RoundSeven2P,
}

var RoundReferencePixels = []pixel.Pixel{
	{X: 140, Y: 316},
	{X: 1204, Y: 316},
	{X: 140, Y: 676},
	{X: 1204, Y: 676},
}

var RoundReferencePixels2P = []pixel.Pixel{
	{X: 163, Y: 678},
	{X: 803, Y: 678},
}

var DoneReferencePixels = []pixel.Pixel{
	{X: 203, Y: 121},
	{X: 843, Y: 121},
	{X: 203, Y: 480},
	{X: 843, Y: 480},
}

var DoneReferencePixels2P = []pixel.Pixel{
	{X: 198, Y: 345},
	{X: 838, Y: 345},
}

func (ga *GameAnalyzer) AnalyzeRounds() {

	ga.capture.MatchSetting(0.92, 0.1)

	roundPixels := RoundPixels
	roundReferencePixels := RoundReferencePixels
	doneReferencePixels := DoneReferencePixels
	donePixel := pixel.Done

	if ga.playerCount == 2 {
		roundPixels = RoundPixels2P
		roundReferencePixels = RoundReferencePixels2P
		doneReferencePixels = DoneReferencePixels2P
		donePixel = pixel.Done2P
	}

	now := time.Now()

	if ga.capture.Matches(pixel.StartRace) {
		ga.initRoundTimes()
		ga.roundStartedAt = now
		ga.exactStartFound = true
	}

	for i := 0; i < ga.playerCount; i++ {
		ref := roundReferencePixels[i]

		newRound := ga.playerRounds[i]
		if ga.playerRounds[i] == 0 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, roundPixels[0])) {
			newRound = 1
			ga.initRoundTimes()
		}
		if ga.playerRounds[i] == 1 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, roundPixels[1])) {
			newRound = 2
		}
		if ga.playerRounds[i] == 2 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, roundPixels[2])) {
			newRound = 3
		}
		if ga.playerRounds[i] == 3 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, roundPixels[3])) {
			newRound = 4
		}
		if ga.playerRounds[i] == 4 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, roundPixels[4])) {
			newRound = 5
		}
		if ga.playerRounds[i] == 5 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, roundPixels[5])) {
			newRound = 6
		}
		if ga.playerRounds[i] == 6 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, roundPixels[6])) {
			newRound = 7
		}

		done := false
		if newRound > 0 {
			done = ga.capture.Matches(pixel.GetAbsolutePixels(doneReferencePixels[i], donePixel))
		}

		if newRound != ga.playerRounds[i] || done {
			if newRound > 1 || done {
				roundTime := now.Sub(ga.playerRoundTimes[i])
				sub := time.Duration(0)
				if newRound == 2 && ga.exactStartFound {
					sub += 4 * time.Second // TODO: Find exact time
				}
				if done {
					sub += 1 * time.Second // TODO: Find exact time
				}
				roundTime -= sub

				finishedRound := ga.playerRounds[i]
				ga.NotifyObservers(func(o mariogo.Observer) {
					o.RoundFinished(i+1, finishedRound, roundTime)
					if done {
						o.PlayerFinishedRace(i+1, now.Sub(ga.roundStartedAt)-sub)
					}
				})

				if done {
					newRound = 0
				}
			}
			ga.playerRoundTimes[i] = now
			ga.playerRounds[i] = newRound
		}
	}

	ga.capture.DefaultMatchSetting()
}

func (ga *GameAnalyzer) initRoundTimes() {
	now := time.Now()
	for i := 0; i < ga.playerCount; i++ {
		ga.playerRoundTimes[i] = now
	}
}
