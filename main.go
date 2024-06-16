package main

import (
	"jkli/mariogo/mariogo"

	"github.com/joho/godotenv"
)

func main() {

	// startTime := time.Now()
	// client := gosseract.NewClient()
	// client.SetLanguage("deu")
	// defer client.Close()
	// client.SetImage("test.png")
	// text, _ := client.Text()
	// fmt.Println(text)
	// fmt.Println("Elapsed time:", time.Since(startTime))

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	mariogo.DatabaseConnect()

	// mariogo.DB.AutoMigrate(&mariogo.Game{}, &mariogo.Round{}, &mariogo.Character{}, &mariogo.Person{}, &mariogo.Player{}, &mariogo.RoundPlacement{}, &mariogo.Placement{})

	// mariogo.SeedCharacter()

	ga := mariogo.NewGameAnalyzer()

	go mariogo.RunWebServer()

	ga.Run()

	// img := ga.GetCurrentFrame()

	// Bild einlesen
	// img := gocv.IMRead("frame_101.png", gocv.IMReadColor)
	// if img.Empty() {
	// 	fmt.Println("Fehler beim Einlesen des Bildes")
	// 	return
	// }

	// croppedMat := img.Region(image.Rect(710, 105, 1180, 155))
	// resultMat := croppedMat.Clone()

	// placements, ok := ga.GetInterimResults(img)

	// fmt.Println("Placements:", placements, "OK:", ok)

	// gocv.IMWrite("testtt.png", resultMat)

}
