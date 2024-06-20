package analyzer

import (
	"fmt"
	"jkli/mariogo/mariogo"
	"jkli/mariogo/mariogo/pixel"
	"time"

	"gocv.io/x/gocv"
)

const (
	Idle           = iota
	Loading        = iota
	Racing         = iota
	Pause          = iota
	RoundResults   = iota
	InterimResults = iota
	EndResults     = iota
)

type GameAnalyzer struct {
	state                 int
	stateUpdatedAt        time.Time
	capture               *mariogo.Capture
	currentRound          int
	playerCount           int
	running               bool
	observers             []mariogo.Observer
	nextRoundName         string
	playerNamesRegistered bool
}

func NewGameAnalyzer() *GameAnalyzer {
	return &GameAnalyzer{
		state:                 Idle,
		capture:               mariogo.NewCapture(),
		playerCount:           0,
		currentRound:          0,
		running:               true,
		playerNamesRegistered: false,
	}
}

func (ga *GameAnalyzer) AddObserver(o mariogo.Observer) {
	ga.observers = append(ga.observers, o)
}

func (ga *GameAnalyzer) NotifyObservers(callback func(mariogo.Observer)) {
	for _, o := range ga.observers {
		callback(o)
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
		startTime := time.Now()
		ga.capture.NextFrame()

		// gocv.IMWrite(fmt.Sprintf("images/frame_%v.png", ga.frame), frame)
		// frame++

		ga.updateState()

		fmt.Println("Time: ", time.Since(startTime))

		time.Sleep(100 * time.Millisecond)
	}
}

func (ga *GameAnalyzer) updateState() {
	newState := ga.state

	switch ga.state {
	case Idle: // -> Racing

		// New round name
		if ga.nextRoundName == "" && ga.capture.Matches(pixel.IntroPage) {
			ga.getRoundName()
		}

		if started, player := ga.gameStarted(); started {
			ga.playerCount = player
			ga.currentRound = 1
			newState = Racing

			ga.NotifyObservers(func(o mariogo.Observer) {
				o.PlayerCount(player)
				o.NewRound(ga.getNextRoundName())
			})
		}
	case Loading:
		// TODO
	case Racing: // -> pause
		if !ga.isRacing() {
			newState = Pause
		}
	case Pause: // -> roundResults | racing
		// back to racing
		if ga.isRacing() {
			newState = Racing
		}

		// round results
		placements, ok := ga.GetRoundResult()
		if ok {
			newState = RoundResults
			ga.NotifyObservers(func(o mariogo.Observer) {
				o.RoundResults(placements)
			})
		}
	case RoundResults: // -> inertimResult | racing

		// New round name
		if ga.nextRoundName == "" && ga.capture.Matches(pixel.IntroPage) {
			ga.getRoundName()
		}
		// new round
		if ga.isRacing() {
			newState = Racing
			ga.currentRound++
			ga.NotifyObservers(func(o mariogo.Observer) {
				o.NewRound(ga.getNextRoundName())
			})
		}

		// interim results
		if results, ok := ga.getInterimResults(); ok {
			newState = InterimResults
			ga.NotifyObservers(func(o mariogo.Observer) {
				o.InterimResults(results)
			})
		}

	case InterimResults: // -> racing | endResults

		if ga.nextRoundName == "" && ga.capture.Matches(pixel.IntroPage) {
			ga.getRoundName()
		}
		// new round
		if ga.isRacing() {
			newState = Racing
			ga.currentRound++
			ga.NotifyObservers(func(o mariogo.Observer) {
				o.NewRound(ga.getNextRoundName())
			})
		}

		// end results
		if ga.capture.Matches(pixel.EndResults) {
			newState = EndResults
		}

	case EndResults:
		// TODO
		newState = Idle
		ga.NotifyObservers(func(o mariogo.Observer) {
			o.StateChange(ga.state, newState)
		})
	}

	if newState != ga.state {
		ga.NotifyObservers(func(o mariogo.Observer) {
			o.StateChange(ga.state, newState)
		})
		gocv.IMWrite(fmt.Sprintf("stateChanges/%v_%v-%v.png", time.Now().Format("20060102150405"), ga.state, newState), *ga.capture.Frame)
		ga.state = newState
		ga.stateUpdatedAt = time.Now()
	}
}
