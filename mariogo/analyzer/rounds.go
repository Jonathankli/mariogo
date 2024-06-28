package analyzer

import (
	"jkli/mariogo/mariogo"
	"jkli/mariogo/mariogo/pixel"
	"time"
)

func (ga *GameAnalyzer) AnalyzeRounds() {

	ga.capture.MatchSetting(0.92, 0.1)

	roundReverencePixels := []pixel.Pixel{
		{X: 140, Y: 316},
		{X: 1204, Y: 316},
		{X: 140, Y: 676},
		{X: 1204, Y: 676},
	}

	doneReverencePixels := []pixel.Pixel{
		{X: 203, Y: 121},
		{X: 843, Y: 121},
		{X: 203, Y: 480},
		{X: 843, Y: 480},
	}

	now := time.Now()

	if ga.capture.Matches(pixel.StartRace) {
		ga.initRoundTimes()
		ga.roundStartedAt = now
		ga.exactStartFound = true
	}

	for i := 0; i < ga.playerCount; i++ {
		ref := roundReverencePixels[i]

		newRound := ga.playerRounds[i]
		if ga.playerRounds[i] == 0 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, pixel.RoundOne)) {
			newRound = 1
			ga.initRoundTimes()
		}
		if ga.playerRounds[i] == 1 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, pixel.RoundTwo)) {
			newRound = 2
		}
		if ga.playerRounds[i] == 2 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, pixel.RoundThree)) {
			newRound = 3
		}
		if ga.playerRounds[i] == 3 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, pixel.RoundFour)) {
			newRound = 4
		}
		if ga.playerRounds[i] == 4 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, pixel.RoundFive)) {
			newRound = 5
		}
		if ga.playerRounds[i] == 5 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, pixel.RoundSix)) {
			newRound = 6
		}
		if ga.playerRounds[i] == 6 && ga.capture.Matches(pixel.GetAbsolutePixels(ref, pixel.RoundSeven)) {
			newRound = 7
		}

		done := false
		if newRound > 0 {
			done = ga.capture.Matches(pixel.GetAbsolutePixels(doneReverencePixels[i], pixel.Done))
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
