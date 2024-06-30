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

func (ga *GameAnalyzer) GetRoundResult() ([]mariogo.PlayerPlacement, bool) {
	LoadResultHashes()
	if ga.playerCount == 1 {
		placement, ok := ga.GetRoundResultOnePlayer()
		return []mariogo.PlayerPlacement{placement}, ok
	}

	var placements []mariogo.PlayerPlacement
	foundPlayer := 0
	foundBots := 0
	botCount := 12 - ga.playerCount

	rowDistance := 52
	p1 := pixel.Pixel{X: 810, Y: 66}
	p2 := pixel.Pixel{X: 828, Y: 85}

	// loop over every row
	for i := 0; i < 12; i++ {

		// get crop of the plus and convert to grayscale
		croppedMat := ga.capture.Frame.Region(image.Rect(p1.X, p1.Y, p2.X, p2.Y))
		grayMat := gocv.NewMat()
		gocv.CvtColor(croppedMat, &grayMat, gocv.ColorBGRToGray)
		croppedImg, _ := grayMat.ToImage()
		grayMat.Close()
		hash, _ := goimagehash.PerceptionHash(croppedImg)

		// compare hash with the result plus hash
		dist, _ := hash.Distance(resultPlusHash)
		distDark, _ := hash.Distance(resultPlusHashDark)

		rowIAPlayer := false
		// if images are similar the plus we found a player
		if dist < 18 || distDark < 18 {
			for p := 0; p < ga.playerCount; p++ {
				vec := croppedMat.GetVecbAt(0, 0)
				color := color.RGBA{vec[2], vec[1], vec[0], 255}
				playerColor := pixel.PlayerColors[p]

				dist, err := ga.capture.ColorDistance(color, playerColor)

				if dist < 0.16 && err == nil {
					placements = append(placements, mariogo.PlayerPlacement{Position: i + 1, IsBot: false, PlayerNumber: p + 1})
					foundPlayer++
					rowIAPlayer = true
					break
				}
			}
		}
		p1.Y = p1.Y + rowDistance
		p2.Y = p2.Y + rowDistance

		if rowIAPlayer {
			continue
		}

		// find bots
		pixels := pixel.AddOffset(pixel.BotPlusP1, 0, i*rowDistance)
		if ga.capture.Matches(pixels) {
			icon := pixel.AddOffset(pixel.CharacterThumbnailP1, 0, i*rowDistance)
			croppedMat := ga.capture.Crop(icon[0].X, icon[0].Y, icon[1].X, icon[1].Y)
			croppedImg, _ := croppedMat.ToImage()
			croppedMat.Close()
			hash, _ := goimagehash.PerceptionHash(croppedImg)
			placements = append(placements, mariogo.PlayerPlacement{Position: i + 1, IsBot: true, IconHash: hash})
			foundBots++
		}

	}

	ok := foundPlayer == ga.playerCount && foundBots == botCount

	// Extract player names
	if ok {
		for _, p := range placements {
			if !p.IsBot && !ga.playerNamesRegistered[p.PlayerNumber-1] {
				ga.getPayerName(p.PlayerNumber, 0, (p.Position-1)*rowDistance)
				ga.playerNamesRegistered[p.PlayerNumber-1] = true
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
