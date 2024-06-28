package analyzer

import (
	"fmt"
	"image"
	"image/color"
	"jkli/mariogo/mariogo"
	"jkli/mariogo/mariogo/pixel"

	"github.com/corona10/goimagehash"
	"gocv.io/x/gocv"
)

var resultPlusHash *goimagehash.ImageHash
var resultPlusHashDark *goimagehash.ImageHash

func LoadResultHashes() {
	plus := gocv.IMRead("mariogo/samples/results/plus.png", gocv.IMReadColor)
	dark := gocv.IMRead("mariogo/samples/results/plus_dark.png", gocv.IMReadColor)
	defer plus.Close()
	defer dark.Close()

	if dark.Empty() || plus.Empty() {
		fmt.Println("Fehler beim Einlesen des Bildes")
		return
	}

	plusimg, _ := plus.ToImage()
	darkimg, _ := dark.ToImage()

	plushash, _ := goimagehash.PerceptionHash(plusimg)
	darkhash, _ := goimagehash.PerceptionHash(darkimg)

	resultPlusHash = plushash
	resultPlusHashDark = darkhash
}

func (ga *GameAnalyzer) GetRoundResult() ([4]int, bool) {
	LoadResultHashes()
	if ga.playerCount == 1 {
		position, ok := ga.GetRoundResultOnePlayer()
		return [4]int{position, 0, 0, 0}, ok
	}

	placements := [4]int{0, 0, 0, 0}
	foundPlayer := 0

	rowDistance := 52
	p1 := pixel.Pixel{X: 810, Y: 66}
	p2 := pixel.Pixel{X: 828, Y: 85}

	for i := 0; i < 12; i++ {

		if foundPlayer == ga.playerCount {
			break
		}

		croppedMat := ga.capture.Frame.Region(image.Rect(p1.X, p1.Y, p2.X, p2.Y))
		grayMat := gocv.NewMat()
		gocv.CvtColor(croppedMat, &grayMat, gocv.ColorBGRToGray)
		croppedImg, _ := grayMat.ToImage()
		grayMat.Close()
		hash, _ := goimagehash.PerceptionHash(croppedImg)

		dist, _ := hash.Distance(resultPlusHash)
		distDark, _ := hash.Distance(resultPlusHashDark)

		if dist < 18 || distDark < 18 {
			for p := 0; p < ga.playerCount; p++ {
				vec := croppedMat.GetVecbAt(0, 0)
				color := color.RGBA{vec[2], vec[1], vec[0], 255}
				playerColor := pixel.PlayerColors[p]

				dist, err := ga.capture.ColorDistance(color, playerColor)

				if dist < 0.16 && err == nil {
					placements[p] = i + 1
					foundPlayer++
					break
				}
			}

		}
		p1.Y = p1.Y + rowDistance
		p2.Y = p2.Y + rowDistance

	}

	ok := foundPlayer == ga.playerCount

	// Extract player names
	if ok {
		for p := 0; p < ga.playerCount; p++ {
			if !ga.playerNamesRegistered[p] {
				ga.getPayerName(p+1, 0, (placements[p]-1)*rowDistance)
				ga.playerNamesRegistered[p] = true
			}
		}
	}

	return placements, ok
}

func (ga *GameAnalyzer) getPayerName(player int, xOffset int, yOffset int) {
	text, err := ga.capture.OCR(470+xOffset, 60+yOffset, 795+xOffset, 90+yOffset)
	fmt.Println("Player name:", text)
	if err != nil {
		fmt.Println("Error reading player name")
		text = fmt.Sprintf("Player %v", player)
	}

	ga.NotifyObservers(func(o mariogo.Observer) {
		o.PlayerName(player, text)
	})
}
