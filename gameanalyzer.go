package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/lucasb-eyer/go-colorful"
	"gocv.io/x/gocv"
)

const (
	idle       = iota
	loading    = iota
	racing     = iota
	results    = iota
	endResults = iota
)

type GameAnalyzer struct {
	state          int
	stateUpdatedAt time.Time
	players        []Player
	rounds         []Round
	webcam         *gocv.VideoCapture
	frame          int
	currentRound   int
	playerCount    int
}

type Player struct {
	position  int
	character string
}

type Round struct {
	number     int
	trackName  string
	placements []struct {
		player   *Player
		position int
	}
}

type Pixel struct {
	x int
	y int
	c color.RGBA
}

type Point struct {
	x int
	y int
}

func NewGameAnalyzer() *GameAnalyzer {
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		panic("Error opening video capture device: 0")
	}

	// defer webcam.Close()

	return &GameAnalyzer{
		state:        idle,
		webcam:       webcam,
		frame:        0,
		playerCount:  0,
		currentRound: 0,
	}
}

func (ga *GameAnalyzer) Run() {
	for {
		// startTime := time.Now()
		frame := ga.GetCurrentFrame()

		// gocv.IMWrite(fmt.Sprintf("images/frame_%v.png", ga.frame), frame)
		// ga.frame++

		ga.updateState(frame)
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
	case idle:
		player := 0
		if ga.Matches(frame, NewRound1Player) {
			player = 1
		} else if ga.Matches(frame, NewRound2Player) {
			player = 2
		} else if ga.Matches(frame, NewRound4Player) {
			player = 4
		} else if ga.Matches(frame, NewRound3Or4Player) {
			player = 3
		}

		if player > 0 {
			newState = racing
			ga.playerCount = player
			ga.currentRound = 1
			fmt.Println("Player ", player, " detected")
		}
	case loading:
		// TODO
	case racing: // -> results || endResults
		placements, ok := ga.GetRoundResult(frame)
		if ok {
			fmt.Println("Round results:", placements)
			newState = results
		}

		if ga.Matches(frame, EndResults) {
			newState = endResults
		}
	case results:
		// restart round
		newRound := false
		if ga.playerCount == 1 && ga.Matches(frame, NewRound1Player) {
			newRound = true
		} else if ga.playerCount == 2 && ga.Matches(frame, NewRound2Player) {
			newRound = true
		} else if (ga.playerCount == 3 || ga.playerCount == 4) && ga.Matches(frame, NewRound3Or4Player) {
			newRound = true
		}

		if newRound {
			newState = racing
			ga.currentRound++
			fmt.Println("New round started")
		}
	case endResults:
		// newState = idle
	}

	if newState != ga.state {
		gocv.IMWrite(fmt.Sprintf("stateChanges/%v_%v-%v.png", time.Now().Format("20060102150405"), ga.state, newState), frame)
		ga.state = newState
		ga.stateUpdatedAt = time.Now()

	}
}

func (ga *GameAnalyzer) GetRoundResult(frame gocv.Mat) ([4]int, bool) {

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

	// if ok && ga.currentRound > 1 {
	// 	ok = ga.Matches(frame, NeutralResultP1) || ga.Matches(frame, PositiveResultP1) || ga.Matches(frame, NegativeResultP1)
	// }

	return placements, ok
}
