package analyzer

import (
	"fmt"
	"jkli/mariogo/mariogo/pixel"
)

func (ga *GameAnalyzer) gameStarted() (bool, int) {
	player := 0
	if ga.capture.Matches(pixel.OnePlayerPlaying) {
		player = 1
	} else if ga.capture.Matches(pixel.TwoPlayerPlaying) {
		player = 2
	} else if ga.capture.Matches(pixel.FourPlayerPlaying) {
		player = 4
	} else if ga.capture.Matches(pixel.ThreePlayerPlaying) {
		player = 3
	}

	return player > 0, player
}

func (ga *GameAnalyzer) isRacing() bool {

	if ga.playerCount == 1 && ga.capture.Matches(pixel.OnePlayerPlaying) {
		return true
	} else if ga.playerCount == 2 && ga.capture.Matches(pixel.TwoPlayerPlaying) {
		return true
	} else if ga.playerCount == 3 && ga.capture.Matches(pixel.ThreePlayerPlaying) {
		return true
	} else if ga.playerCount == 4 && ga.capture.Matches(pixel.FourPlayerPlaying) {
		return true
	}

	return false
}

func (ga *GameAnalyzer) getRoundName() {

	round := ga.currentRound + 1
	text, err := ga.capture.OCR(340, 610, 945, 670)

	if err != nil || text == "" {
		fmt.Println("Error reading round name")
		text = fmt.Sprintf("Round %v", round)
	}

	fmt.Println("Round ", round, ": ", text)
}
