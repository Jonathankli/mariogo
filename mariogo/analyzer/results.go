package analyzer

import (
	"fmt"
	"image/color"
	"jkli/mariogo/mariogo/pixel"
)

func (ga *GameAnalyzer) GetRoundResult() ([4]int, bool) {

	if ga.playerCount == 1 {
		position, ok := ga.GetRoundResultOnePlayer()
		return [4]int{position, 0, 0, 0}, ok
	}

	colors := [4]color.RGBA{
		{250, 229, 38, 255},  //P1
		{33, 229, 251, 255},  //P2
		{253, 116, 116, 255}, //P3
		{115, 242, 40, 255},  //P4
	}

	placements := [4]int{0, 0, 0, 0}
	foundPlayer := 0

	rowDistance := 52
	row := [6]pixel.Pixel{
		{X: 305, Y: 58},
		{X: 305, Y: 95},
		{X: 640, Y: 95},
		{X: 640, Y: 58},
		{X: 973, Y: 58},
		{X: 973, Y: 95},
	}

	for i := 0; i < 12; i++ {

		if foundPlayer == ga.playerCount {
			break
		}

		for p := 0; p < ga.playerCount; p++ {

			row[0].C = colors[p]
			row[1].C = colors[p]
			row[2].C = colors[p]
			row[3].C = colors[p]
			row[4].C = colors[p]
			row[5].C = colors[p]

			if ga.capture.Matches(row[:]) {
				placements[p] = i + 1
				foundPlayer++
				fmt.Println("Player ", p+1, " found at position ", i+1)
				if ga.currentRound == 1 && ga.observer.GetRegisteredPlayer() != ga.playerCount {
					ga.getPayerName(p+1, 0, rowDistance*i)
				}

			}

			if foundPlayer == ga.playerCount {
				break
			}
		}

		row[0].Y += rowDistance
		row[1].Y += rowDistance
		row[2].Y += rowDistance
		row[3].Y += rowDistance
		row[4].Y += rowDistance
		row[5].Y += rowDistance

	}

	ok := foundPlayer == ga.playerCount

	return placements, ok
}

func (ga *GameAnalyzer) getInterimResults() ([4]int, bool) {

	if ga.playerCount == 1 {
		position, ok := ga.getInterimResultOnePlayer()
		return [4]int{position, 0, 0, 0}, ok
	}

	results, ok := ga.GetRoundResult()

	if ok {
		ok = ga.capture.Matches(pixel.NeutralResultP1) || ga.capture.Matches(pixel.PositiveResultP1) || ga.capture.Matches(pixel.NegativeResultP1)
	}

	return results, ok
}

func (ga *GameAnalyzer) getPayerName(player int, xOffset int, yOffset int) {
	text, err := ga.capture.OCR(470+xOffset, 60+yOffset, 795+xOffset, 90+yOffset)

	if err != nil {
		fmt.Println("Error reading player name")
		text = fmt.Sprintf("Player %v", player)
	}

	fmt.Println("Player ", player, ": ", text)
	ga.observer.RegisterPlayer(player, text)
}
