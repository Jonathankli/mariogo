package analyzer

import (
	"fmt"
	"jkli/mariogo/mariogo"
	"jkli/mariogo/mariogo/pixel"
	"time"

	"gocv.io/x/gocv"
)

const (
	idle          = iota
	loading       = iota
	racing        = iota
	pause         = iota
	roundResults  = iota
	interimResult = iota
	endResults    = iota
)

type GameAnalyzer struct {
	state             int
	stateUpdatedAt    time.Time
	gameModel         mariogo.Game
	capture           *mariogo.Capture
	currentRound      int
	playerCount       int
	running           bool
	nextRoundCaptured bool
	observer          *mariogo.GameObserver
}

func NewGameAnalyzer() *GameAnalyzer {
	return &GameAnalyzer{
		state:             idle,
		capture:           mariogo.NewCapture(),
		playerCount:       0,
		currentRound:      0,
		nextRoundCaptured: false,
		running:           true,
	}
}

func (ga *GameAnalyzer) Stop() {
	ga.running = false
	defer ga.capture.Stop()
}

func (ga *GameAnalyzer) Run() {
	fmt.Println("Start game analyzer")
	// frame := 0
	for ga.running {
		// startTime := time.Now()
		ga.capture.NextFrame()

		// gocv.IMWrite(fmt.Sprintf("images/frame_%v.png", ga.frame), frame)
		// frame++

		ga.updateState()

		// fmt.Println("Time:", time.Since(startTime))

		time.Sleep(100 * time.Millisecond)
	}
}

func (ga *GameAnalyzer) updateState() {
	newState := ga.state

	switch ga.state {
	case idle: // -> loading
		if !ga.nextRoundCaptured && ga.capture.Matches(pixel.IntroPage) {
			ga.getRoundName()
			ga.nextRoundCaptured = true
		}
		if started, player := ga.gameStarted(); started {
			ga.playerCount = player
			ga.currentRound = 1
			newState = racing
			ga.observer = mariogo.NewGameObserver(player)
			ga.observer.NewRound(ga.currentRound, nil)

			fmt.Println("Game started with", player, "players")
		}
	case loading:
		// TODO
	case racing: // -> pause
		if !ga.isRacing() {
			newState = pause
			fmt.Println("Game paused")
		}
	case pause: // -> roundResults | racing
		// back to racing
		if ga.isRacing() {
			newState = racing
			fmt.Println("Game resumed")
		}

		// round results
		placements, ok := ga.GetRoundResult()
		if ok {
			fmt.Println("Round results:", placements)
			ga.observer.RoundResults(placements)
			newState = roundResults
		}
	case roundResults: // -> inertimResult | racing

		if !ga.nextRoundCaptured && ga.capture.Matches(pixel.IntroPage) {
			ga.getRoundName()
			ga.nextRoundCaptured = true
		}
		// new round
		if ga.isRacing() {
			newState = racing
			ga.currentRound++
			ga.observer.NewRound(ga.currentRound, nil)
			fmt.Println("New round started")
		}

		// interim results
		if results, ok := ga.getInterimResults(); ok {
			newState = interimResult
			ga.observer.InterimResults(results)
			fmt.Println("Interim results:", results)
		}

	case interimResult: // -> racing | endResults

		if !ga.nextRoundCaptured && ga.capture.Matches(pixel.IntroPage) {
			ga.getRoundName()
			ga.nextRoundCaptured = true
		}
		// new round
		if ga.isRacing() {
			newState = racing
			ga.currentRound++
			ga.observer.NewRound(ga.currentRound, nil)
			fmt.Println("New round started")
		}

		// end results
		if ga.capture.Matches(pixel.EndResults) {
			newState = endResults
		}

	case endResults:
		// TODO
		fmt.Println("Game ended")
		ga.observer.Finish()
		newState = idle
	}

	if newState != ga.state {
		gocv.IMWrite(fmt.Sprintf("stateChanges/%v_%v-%v.png", time.Now().Format("20060102150405"), ga.state, newState), *ga.capture.Frame)
		ga.state = newState
		ga.stateUpdatedAt = time.Now()

	}
}
