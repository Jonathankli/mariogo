package analyzer

import (
	"fmt"
	"jkli/mariogo/mariogo"
	"jkli/mariogo/mariogo/pixel"
	"time"

	"github.com/corona10/goimagehash"
	"gocv.io/x/gocv"
)

type OffsetPoints struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

var PlacementHashes [12][2]goimagehash.ImageHash
var PlacementHashes2P [12][2]goimagehash.ImageHash

var PlacementReference = [4]pixel.Pixel{
	{X: 75, Y: 250},
	{X: 1185, Y: 250},
	{X: 75, Y: 610},
	{X: 1185, Y: 610},
}

var Offset = OffsetPoints{X1: -12, Y1: -10, X2: 20, Y2: 40}

var PlacementReference2Player = [4]pixel.Pixel{
	{X: 531, Y: 636},
	{X: 1171, Y: 636},
}

var Offset2P = OffsetPoints{X1: -15, Y1: -16, X2: 26, Y2: 49}

func GeneratePlacmentHashes() {

	for i := 1; i <= 12; i++ {
		img := gocv.IMRead(fmt.Sprintf("mariogo/samples/placements/%v.png", i), gocv.IMReadColor)
		imgDark := gocv.IMRead(fmt.Sprintf("mariogo/samples/placements/%v_dark.png", i), gocv.IMReadColor)
		img2P := gocv.IMRead(fmt.Sprintf("mariogo/samples/placements2player/%v.png", i), gocv.IMReadColor)
		img2PDark := gocv.IMRead(fmt.Sprintf("mariogo/samples/placements2player/%v_dark.png", i), gocv.IMReadColor)
		defer img.Close()
		defer imgDark.Close()
		defer img2P.Close()
		defer img2PDark.Close()

		if img.Empty() || imgDark.Empty() || img2P.Empty() || img2PDark.Empty() {
			fmt.Println("Fehler beim Einlesen des Bildes")
			return
		}

		image, _ := img.ToImage()
		refImgHash, _ := goimagehash.PerceptionHash(image)

		imageDark, _ := imgDark.ToImage()
		refImgHashDark, _ := goimagehash.PerceptionHash(imageDark)

		PlacementHashes[i-1] = [2]goimagehash.ImageHash{*refImgHash, *refImgHashDark}

		image2P, _ := img2P.ToImage()
		refImgHash2P, _ := goimagehash.PerceptionHash(image2P)

		image2PDark, _ := img2PDark.ToImage()
		refImgHash2PDark, _ := goimagehash.PerceptionHash(image2PDark)

		PlacementHashes2P[i-1] = [2]goimagehash.ImageHash{*refImgHash2P, *refImgHash2PDark}
	}

}

func (ga *GameAnalyzer) AnalyzeCurrentPlacements() {

	//Only 2, 3 and 4 player games are supported
	if ga.playerCount < 2 {
		return
	}

	offset := Offset
	placementReference := PlacementReference
	placementHashes := PlacementHashes

	if ga.playerCount == 2 {
		offset = Offset2P
		placementReference = PlacementReference2Player
		placementHashes = PlacementHashes2P
	}

	newPlacements := ga.placements
	change := false

	for i := 0; i < ga.playerCount; i++ {
		refPx := placementReference[i]

		cropped := ga.capture.Crop(refPx.X+offset.X1, refPx.Y+offset.Y1, refPx.X+offset.X2, refPx.Y+offset.Y2)
		placeMat := gocv.NewMat()

		gocv.CvtColor(*cropped, &placeMat, gocv.ColorBGRToGray)
		image, _ := placeMat.ToImage()

		cropped.Close()
		placeMat.Close()

		place := 0

		for i, hash := range placementHashes {
			comp, _ := goimagehash.PerceptionHash(image)

			dist, _ := hash[0].Distance(comp)
			distDark, _ := hash[1].Distance(comp)

			if dist < 8 || distDark < 8 {
				place = i + 1
				break
			}
		}

		if place != newPlacements[i] && place != 0 {
			newPlacements[i] = place
			change = true
		}
	}

	if change {
		oldPlacements := ga.placements

		//make sure every placement is unique
		for i := 0; i < ga.playerCount; i++ { //loop over every player
			for j := 0; j < ga.playerCount; j++ { //loop over every player to compare
				isSamePlayer := i == j
				if !isSamePlayer && newPlacements[i] == newPlacements[j] {

					// if tow players have the same placement, we assume they switched places and increment or decrement the one who has not changed
					if oldPlacements[i] == newPlacements[i] {
						newPlacements[j]++
					} else {
						newPlacements[j]--
					}

				}
			}
		}

		since := time.Since(ga.roundStartedAt)
		ga.NotifyObservers(func(o mariogo.Observer) {
			o.PlacementsChanged(oldPlacements, newPlacements, since)
		})

		ga.placements = newPlacements
	}

}
