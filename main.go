package main

import (
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	DatabaseConnect()

	// DB.AutoMigrate(&Game{}, &Round{}, &Character{}, &Person{}, &Player{}, &RoundPlacement{}, &Placement{})

	// SeedCharacter()

	ga := NewGameAnalyzer()

	go runWebServer()

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
