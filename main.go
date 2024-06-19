package main

import (
	"jkli/mariogo/mariogo"
	"jkli/mariogo/mariogo/analyzer"
	"jkli/mariogo/mariogo/observer"

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

	ga := analyzer.NewGameAnalyzer()

	ga.AddObserver(&observer.Logger{})
	ga.AddObserver(&observer.Database{})

	// go mariogo.RunWebServer()

	ga.Run()

	// img := ga.GetCurrentFrame()

	// Bild einlesen
	// img := gocv.IMRead("images/testdata/race_results_4p.png", gocv.IMReadColor)
	// if img.Empty() {
	// 	fmt.Println("Fehler beim Einlesen des Bildes")
	// 	return
	// }

	// placements, ok := ga.GetRoundResult(img)
	// fmt.Println("Placements:", placements, "OK:", ok)

	// start := time.Now()
	// gocv.Resize(img, &img, image.Point{X: 1280, Y: 720}, 0, 0, gocv.InterpolationLinear)
	// fmt.Println("Elapsed time:", time.Since(start))
	// croppedMat := img.Region(image.Rect(710, 105, 1180, 155))
	// resultMat := croppedMat.Clone()

	// placements, ok := ga.GetInterimResults(img)

	// fmt.Println("Placements:", placements, "OK:", ok)

	// gocv.IMWrite("res.png", img)

}
