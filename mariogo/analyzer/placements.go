package analyzer

import (
	"fmt"
	"jkli/mariogo/mariogo"
	"jkli/mariogo/mariogo/pixel"

	"github.com/corona10/goimagehash"
	"gocv.io/x/gocv"
)

var placementHashes [12][2]goimagehash.ImageHash

var PlacementReference = [4]pixel.Pixel{
	{X: 75, Y: 250},
	{X: 1185, Y: 250},
	{X: 75, Y: 610},
	{X: 1185, Y: 610},
}

func GeneratePlacmentHashes() {

	for i := 1; i <= 12; i++ {
		img := gocv.IMRead(fmt.Sprintf("mariogo/samples/placements/%v.png", i), gocv.IMReadColor)
		imgDark := gocv.IMRead(fmt.Sprintf("mariogo/samples/placements/%v_dark.png", i), gocv.IMReadColor)
		if img.Empty() || imgDark.Empty() {
			fmt.Println("Fehler beim Einlesen des Bildes")
			return
		}

		gocv.CvtColor(img, &img, gocv.ColorBGRToGray)
		gocv.CvtColor(imgDark, &imgDark, gocv.ColorBGRToGray)

		image, _ := img.ToImage()
		refImgHash, _ := goimagehash.PerceptionHash(image)

		imageDark, _ := imgDark.ToImage()
		refImgHashDark, _ := goimagehash.PerceptionHash(imageDark)

		placementHashes[i-1] = [2]goimagehash.ImageHash{*refImgHash, *refImgHashDark}
	}

}

func (ga *GameAnalyzer) AnalyzeCurrentPlacements() {

	//Only 3 and 4 player games are supported
	if ga.playerCount < 3 {
		return
	}

	newPlacements := ga.placements
	change := false

	for i := 0; i < ga.playerCount; i++ {
		refPx := PlacementReference[i]
		cropped := ga.capture.Crop(refPx.X-12, refPx.Y-10, refPx.X+20, refPx.Y+40)
		gocv.CvtColor(*cropped, cropped, gocv.ColorBGRToGray)
		image, _ := cropped.ToImage()

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
		ga.NotifyObservers(func(o mariogo.Observer) {
			o.PlacementsChanged(oldPlacements, newPlacements)
		})

		//TODO: Check if placements are valid an no double placements
		ga.placements = newPlacements
	}

}
