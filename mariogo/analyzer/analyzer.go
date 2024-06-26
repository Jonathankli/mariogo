package analyzer

import (
	"fmt"
	"jkli/mariogo/mariogo"
	"jkli/mariogo/mariogo/pixel"
	"os"
	"reflect"
	"strconv"
	"time"

	"gocv.io/x/gocv"
)

const (
	Idle         = iota
	Racing       = iota
	Pause        = iota
	RoundResults = iota
	EndResults   = iota
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
	playerNamesRegistered [4]bool
	playerRounds          [4]int
	placements            [4]int
	playerRoundTimes      [4]time.Time
	exactStartFound       bool
	enableDebugImages     bool
	enableDebugTimes      bool
	maxFPS                int
	roundStartedAt        time.Time
	keepStateChangeImages bool
}

func NewGameAnalyzer() *GameAnalyzer {
	// Load and generate hashes
	GeneratePlacmentHashes()
	LoadResultHashes()

	fps, err := strconv.Atoi(mariogo.Getenv("MAX_FPS", "30"))

	if err != nil {
		fmt.Println("Error parsing MAX_FPS")
		fps = 30
	}

	ga := &GameAnalyzer{
		capture:               mariogo.NewCapture(),
		enableDebugImages:     os.Getenv("DEBUG_IMAGES") == "true",
		enableDebugTimes:      os.Getenv("DEBUG_TIMES") == "true",
		keepStateChangeImages: os.Getenv("KEEP_STATE_CHANGE_IMAGES") == "true",
		maxFPS:                fps,
	}

	ga.DefaultState()

	return ga
}

func (ga *GameAnalyzer) DefaultState() {
	ga.state = Idle
	ga.stateUpdatedAt = time.Now()
	ga.currentRound = 0
	ga.playerCount = 0
	ga.nextRoundName = ""
	ga.playerNamesRegistered = [4]bool{false, false, false, false}
	ga.playerRounds = [4]int{0, 0, 0, 0}
	ga.placements = [4]int{0, 0, 0, 0}
	ga.playerRoundTimes = [4]time.Time{}
	ga.exactStartFound = false
	ga.running = true
}

func (ga *GameAnalyzer) AddObserver(o mariogo.Observer) {
	ga.observers = append(ga.observers, o)
}

func (ga *GameAnalyzer) NotifyObservers(callback func(mariogo.Observer)) {
	for _, o := range ga.observers {
		go func(o mariogo.Observer) {
			// Recover from observer panics
			defer ga.catchObserverError(o)

			// Call observer
			callback(o)
		}(o)
	}
}

func (ga *GameAnalyzer) Stop() {
	ga.running = false
	defer ga.capture.Stop()
}

func (ga *GameAnalyzer) catchObserverError(o mariogo.Observer) {
	if r := recover(); r != nil {
		t := reflect.TypeOf(o)
		fmt.Println("Observer error in:", t)
		// TODO: log error to file
	}
}

func (ga *GameAnalyzer) recover() {
	if r := recover(); r != nil {
		fmt.Println("GameAnalyzer error")
		// TODO: log error to file
		time.Sleep(time.Millisecond * 500)
		ga.Run() //rerun game analyzer
	}
}

func (ga *GameAnalyzer) Run() {

	// recover from panics
	defer ga.recover()

	fmt.Println("Start game gokart")
	frame := 0
	for ga.running {
		startTime := time.Now()
		ga.capture.NextFrame()

		if ga.enableDebugImages {
			gocv.IMWrite(fmt.Sprintf("images/frame_%v.png", frame), *ga.capture.Frame)
			frame++
		}

		ga.updateState()

		took := time.Since(startTime)
		if ga.enableDebugTimes {
			fmt.Println("Time:", took)
		}

		if took < time.Second/time.Duration(ga.maxFPS) {
			sleep := time.Second/time.Duration(ga.maxFPS) - took
			time.Sleep(sleep)
		}
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
			roundName := ga.getNextRoundName()
			ga.roundStartedAt = time.Now().Add(time.Second * 4) // TODO: Find exact time

			ga.NotifyObservers(func(o mariogo.Observer) {
				o.PlayerCount(player)
				o.NewRound(roundName)
			})
		}
	case Racing: // -> pause
		if !ga.isRacing() {
			newState = Pause
			break
		}
		ga.AnalyzeRounds()
		ga.AnalyzeCurrentPlacements()
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
	case RoundResults: // -> endResults | racing

		// New round name
		if ga.nextRoundName == "" && ga.capture.Matches(pixel.IntroPage) {
			ga.getRoundName()
		}
		// new round
		if ga.isRacing() {
			newState = Racing
			ga.currentRound++
			roundName := ga.getNextRoundName()
			ga.NotifyObservers(func(o mariogo.Observer) {
				o.NewRound(roundName)
			})
		}

		// end results
		if ga.capture.Matches(pixel.EndResults) {
			newState = EndResults
		}

	case EndResults:
		// TODO: do something
		newState = Idle
		ga.DefaultState()
	}

	if newState != ga.state {
		ga.NotifyObservers(func(o mariogo.Observer) {
			o.StateChange(ga.state, newState)
		})

		ga.stateUpdatedAt = time.Now()
		ga.state = newState

		if ga.keepStateChangeImages {
			go gocv.IMWrite(fmt.Sprintf("images/stateChanges/%v_%v-%v.png", ga.stateUpdatedAt.Format("20060102150405"), ga.state, newState), *ga.capture.Frame)
		}
	}
}
