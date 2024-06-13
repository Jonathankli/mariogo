package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {

	ga := NewGameAnalyzer()

	ga.Run()

	// img := ga.GetCurrentFrame()

	// Bild einlesen
	img := gocv.IMRead("images/4playerwithrfesults/frame_101.png", gocv.IMReadColor)
	if img.Empty() {
		fmt.Println("Fehler beim Einlesen des Bildes")
		return
	}

	placements, ok := ga.GetRoundResult(img)

	fmt.Println("Placements:", placements, "OK:", ok)

	gocv.IMWrite("testtt.png", img)

}
