package mariogo

import (
	"fmt"
	"image/color"
	"time"

	"github.com/lucasb-eyer/go-colorful"
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
	state          int
	stateUpdatedAt time.Time
	gameModel      Game
	webcam         *gocv.VideoCapture
	frame          int
	currentRound   int
	playerCount    int
	running        bool
	observer       *GameObserver
}

type Pixel struct {
	x int
	y int
	c color.RGBA
}

func NewGameAnalyzer() *GameAnalyzer {
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		panic("Error opening video capture device: 0")
	}

	return &GameAnalyzer{
		state:        idle,
		webcam:       webcam,
		frame:        0,
		playerCount:  0,
		currentRound: 0,
		running:      true,
	}
}

func (ga *GameAnalyzer) Stop() {
	ga.running = false
	defer ga.webcam.Close()
}

func (ga *GameAnalyzer) Run() {
	fmt.Println("Start!")
	for ga.running {
		// startTime := time.Now()
		frame := ga.GetCurrentFrame()

		// gocv.IMWrite(fmt.Sprintf("images/frame_%v.png", ga.frame), frame)
		// ga.frame++

		ga.updateState(frame)
		frame.Close()

		// fmt.Println("Time:", time.Since(startTime))

		time.Sleep(500 * time.Millisecond)
	}
}

func (ga *GameAnalyzer) GetCurrentFrame() gocv.Mat {
	img := gocv.NewMat()
	// defer img.Close()

	if ok := ga.webcam.Read(&img); !ok {
		panic("cannot read device 0")
	}

	if img.Empty() {
		panic("no image on device 0")
	}

	return img
}

func (ga *GameAnalyzer) Matches(frame gocv.Mat, pixels []Pixel) bool {
	matchCount := 0

	for _, pixel := range pixels {
		vec := frame.GetVecbAt(pixel.y, pixel.x)

		color := color.RGBA{vec[2], vec[1], vec[0], 255}
		color1, ok1 := colorful.MakeColor(pixel.c)
		color2, ok2 := colorful.MakeColor(color)

		if !ok1 || !ok2 {
			fmt.Printf("Error converting color: %v\n", vec)
			continue
		}

		if color1.DistanceCIEDE2000(color2) > 0.1 {
			continue
		}

		matchCount++
	}

	matchPercentage := float64(matchCount) / float64(len(pixels))

	return matchPercentage > 0.8
}

func (ga *GameAnalyzer) updateState(frame gocv.Mat) {
	newState := ga.state

	switch ga.state {
	case idle: // -> loading
		if started, player := ga.gameStarted(frame, newState); started {
			ga.playerCount = player
			ga.currentRound = 1
			newState = racing
			ga.observer = NewGameObserver()

			player := []string{}
			for i := 0; i < ga.playerCount; i++ {
				player = append(player, fmt.Sprintf("Player %v", i+1))
			}
			ga.observer.InitPlayer(player)
			ga.observer.NewRound(ga.currentRound, nil)

			fmt.Println("Game started with", player, "players")
		}
	case loading:
		// TODO
	case racing: // -> pause
		if !ga.isRacing(frame, newState) {
			newState = pause
			fmt.Println("Game paused")
		}
	case pause: // -> roundResults | racing
		// back to racing
		if ga.isRacing(frame, newState) {
			newState = racing
			fmt.Println("Game resumed")
		}

		// round results
		placements, ok := ga.getRoundResult(frame)
		if ok {
			fmt.Println("Round results:", placements)
			ga.observer.RoundResults(placements)
			newState = roundResults
		}
	case roundResults: // -> inertimResult | racing
		// new round
		if ga.isRacing(frame, newState) {
			newState = racing
			ga.currentRound++
			ga.observer.NewRound(ga.currentRound, nil)
			fmt.Println("New round started")
		}

		// interim results
		if results, ok := ga.getInterimResults(frame); ok {
			newState = interimResult
			ga.observer.InterimResults(results)
			fmt.Println("Interim results:", results)
		}

	case interimResult: // -> racing | endResults

		// new round
		if ga.isRacing(frame, newState) {
			newState = racing
			ga.currentRound++
			ga.observer.NewRound(ga.currentRound, nil)
			fmt.Println("New round started")
		}

		// end results
		if ga.Matches(frame, EndResults) {
			newState = endResults
		}

	case endResults:
		// TODO
		fmt.Println("Game ended")
		ga.observer.Finish()
		newState = idle
	}

	if newState != ga.state {
		gocv.IMWrite(fmt.Sprintf("stateChanges/%v_%v-%v.png", time.Now().Format("20060102150405"), ga.state, newState), frame)
		ga.state = newState
		ga.stateUpdatedAt = time.Now()

	}
}

func (ga *GameAnalyzer) gameStarted(frame gocv.Mat, state int) (bool, int) {
	player := 0
	if ga.Matches(frame, OnePlayerPlaying) {
		player = 1
	} else if ga.Matches(frame, TwoPlayerPlaying) {
		player = 2
	} else if ga.Matches(frame, FourPlayerPlaying) {
		player = 4
	} else if ga.Matches(frame, ThreePlayerPlaying) {
		player = 3
	}

	return player > 0, player
}

func (ga *GameAnalyzer) isRacing(frame gocv.Mat, state int) bool {

	if ga.playerCount == 1 && ga.Matches(frame, OnePlayerPlaying) {
		return true
	} else if ga.playerCount == 2 && ga.Matches(frame, TwoPlayerPlaying) {
		return true
	} else if ga.playerCount == 3 && ga.Matches(frame, ThreePlayerPlaying) {
		return true
	} else if ga.playerCount == 1 && ga.Matches(frame, FourPlayerPlaying) {
		return true
	}

	return false
}

func (ga *GameAnalyzer) getRoundResult(frame gocv.Mat) ([4]int, bool) {

	if ga.playerCount == 1 {
		position, ok := ga.getRoundResultOnePlayer(frame)
		return [4]int{position, 0, 0, 0}, ok
	}

	colors := [4]color.RGBA{
		color.RGBA{250, 229, 38, 255},  //P1
		color.RGBA{33, 229, 251, 255},  //P2
		color.RGBA{253, 116, 116, 255}, //P3
		color.RGBA{115, 242, 40, 255},  //P4
	}

	placements := [4]int{0, 0, 0, 0}
	foundPlayer := 0

	rowDistance := 75
	row := [6]Pixel{
		Pixel{x: 485, y: 100},
		Pixel{x: 485, y: 155},
		Pixel{x: 1435, y: 155},
		Pixel{x: 1435, y: 100},
		Pixel{x: 955, y: 100},
		Pixel{x: 955, y: 155},
	}

	for i := 0; i < 12; i++ {

		if foundPlayer == ga.playerCount {
			break
		}

		for p := 0; p < ga.playerCount; p++ {

			row[0].c = colors[p]
			row[1].c = colors[p]
			row[2].c = colors[p]
			row[3].c = colors[p]
			row[4].c = colors[p]
			row[5].c = colors[p]

			if ga.Matches(frame, row[:]) {
				placements[p] = i + 1
				foundPlayer++
			}

			if foundPlayer == ga.playerCount {
				break
			}
		}

		row[0].y += rowDistance
		row[1].y += rowDistance
		row[2].y += rowDistance
		row[3].y += rowDistance
		row[4].y += rowDistance
		row[5].y += rowDistance

	}

	ok := foundPlayer == ga.playerCount

	return placements, ok
}

func (ga *GameAnalyzer) getRoundResultOnePlayer(frame gocv.Mat) (int, bool) {

	color := color.RGBA{250, 229, 38, 255} //P1

	placement := 0

	rowDistance := 75
	row := [6]Pixel{
		Pixel{x: 840, y: 100, c: color},
		Pixel{x: 840, y: 155, c: color},
		Pixel{x: 1790, y: 155, c: color},
		Pixel{x: 1790, y: 100, c: color},
		Pixel{x: 1310, y: 100, c: color},
		Pixel{x: 1310, y: 155, c: color},
	}

	for i := 0; i < 12; i++ {

		if ga.Matches(frame, row[:]) {
			placement = i + 1
			break
		}

		row[0].y += rowDistance
		row[1].y += rowDistance
		row[2].y += rowDistance
		row[3].y += rowDistance
		row[4].y += rowDistance
		row[5].y += rowDistance

	}

	ok := placement > 0

	return placement, ok
}

func (ga *GameAnalyzer) getInterimResults(frame gocv.Mat) ([4]int, bool) {

	if ga.playerCount == 1 {
		position, ok := ga.getInterimResultOnePlayer(frame)
		return [4]int{position, 0, 0, 0}, ok
	}

	results, ok := ga.getRoundResult(frame)

	if ok {
		ok = ga.Matches(frame, NeutralResultP1) || ga.Matches(frame, PositiveResultP1) || ga.Matches(frame, NegativeResultP1)
	}

	return results, ok
}

func (ga *GameAnalyzer) getInterimResultOnePlayer(frame gocv.Mat) (int, bool) {
	result, ok := ga.getRoundResultOnePlayer(frame)

	if ok {
		xOffset := 360

		neutral := ga.addOffset(NeutralResultP1, xOffset, 0)
		positive := ga.addOffset(PositiveResultP1, xOffset, 0)
		negative := ga.addOffset(NegativeResultP1, xOffset, 0)

		ok = ga.Matches(frame, neutral) || ga.Matches(frame, positive) || ga.Matches(frame, negative)
	}

	return result, ok
}

func (ga *GameAnalyzer) addOffset(pixels []Pixel, x int, y int) []Pixel {
	for i := range pixels {
		pixels[i].x += x
		pixels[i].y += y
	}

	return pixels
}
