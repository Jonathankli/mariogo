package main

import "github.com/joho/godotenv"

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	DatabaseConnect()

	// DB.AutoMigrate(&Game{}, &Round{}, &Character{}, &Person{}, &Player{}, &RoundPlacement{}, &Placement{})

	ga := NewGameAnalyzer()

	go runWebServer()

	ga.Run()

	// img := ga.GetCurrentFrame()

	// Bild einlesen
	// img := gocv.IMRead("images/2plong/frame_278.png", gocv.IMReadColor)
	// if img.Empty() {
	// 	fmt.Println("Fehler beim Einlesen des Bildes")
	// 	return
	// }

	// placements, ok := ga.GetRoundResult(img)

	// fmt.Println("Placements:", placements, "OK:", ok)

	// gocv.IMWrite("testtt.png", img)

}
