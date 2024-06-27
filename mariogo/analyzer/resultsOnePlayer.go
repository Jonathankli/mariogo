package analyzer

import (
	"image/color"
	"jkli/mariogo/mariogo/pixel"
)

func (ga *GameAnalyzer) GetRoundResultOnePlayer() (int, bool) {

	color := color.RGBA{250, 229, 38, 255} //P1

	placement := 0

	rowDistance := 52
	row := [6]pixel.Pixel{
		{X: 555, Y: 58, C: color},
		{X: 555, Y: 95, C: color},
		{X: 890, Y: 95, C: color},
		{X: 890, Y: 58, C: color},
		{X: 1223, Y: 58, C: color},
		{X: 1223, Y: 95, C: color},
	}

	for i := 0; i < 12; i++ {

		if ga.capture.Matches(row[:]) {
			placement = i + 1
			break
		}

		row[0].Y += rowDistance
		row[1].Y += rowDistance
		row[2].Y += rowDistance
		row[3].Y += rowDistance
		row[4].Y += rowDistance
		row[5].Y += rowDistance

	}

	ok := placement > 0

	return placement, ok
}
