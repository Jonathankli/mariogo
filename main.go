package main

import (
	"jkli/mariogo/mariogo"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	mariogo.DatabaseConnect()

	// mariogo.DB.AutoMigrate(&Game{}, &Round{}, &Character{}, &Person{}, &Player{}, &RoundPlacement{}, &Placement{})

	// mariogo.SeedCharacter()

	ga := mariogo.NewGameAnalyzer()

	go mariogo.RunWebServer()

	ga.Run()

	// img := ga.GetCurrentFrame()

	// Bild einlesen
	// img := gocv.IMRead("images/frame_1.png", gocv.IMReadColor)
	// if img.Empty() {
	// 	fmt.Println("Fehler beim Einlesen des Bildes")
	// 	return
	// }

	// placements, ok := ga.GetInterimResults(img)

	// fmt.Println("Placements:", placements, "OK:", ok)

	// gocv.IMWrite("testtt.png", img)

}
