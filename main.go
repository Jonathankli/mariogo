package main

import (
	"jkli/mariogo/mariogo"
	"jkli/mariogo/mariogo/analyzer"
	"jkli/mariogo/mariogo/api"
	"jkli/mariogo/mariogo/observer"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	mariogo.DatabaseConnect()
	mariogo.DB.AutoMigrate(&mariogo.Game{}, &mariogo.Round{}, &mariogo.Character{}, &mariogo.Person{}, &mariogo.Player{}, &mariogo.Placement{}, &mariogo.PlacementChangeLog{}, &mariogo.RoundTime{})
	go api.RunWebServer()

	ga := analyzer.NewGameAnalyzer()
	ga.AddObserver(&observer.Logger{})
	ga.AddObserver(&observer.Database{})
	ga.Run()

}

// reverence := pixel.PlacementReference[0]

// fmt.Println("Colored:")
// pixel.PrintPx(pixel.GetRelativePixels(reverence, pixel.Place12Colored))

// fmt.Println("Border:")
// pixel.PrintPx(pixel.GetRelativePixels(reverence, pixel.Place12Border))

// startTime := time.Now()
// client := gosseract.NewClient()
// client.SetLanguage("deu")
// defer client.Close()
// client.SetImage("test.png")
// text, _ := client.Text()
// fmt.Println(text)
// fmt.Println("Elapsed time:", time.Since(startTime))

// err := godotenv.Load()

// if err != nil {
// 	panic("Error loading .env file")
// }

// mariogo.DatabaseConnect()

// mariogo.DB.AutoMigrate(&mariogo.Game{}, &mariogo.Round{}, &mariogo.Character{}, &mariogo.Person{}, &mariogo.Player{}, &mariogo.RoundPlacement{}, &mariogo.Placement{}, &mariogo.PlacementChangeLog{}, &mariogo.RoundTime{})

// mariogo.SeedCharacter()

// ga := analyzer.NewGameAnalyzer()

// ga.AddObserver(&observer.Logger{})

// ga.GetCurrentPlacements()

// ga.AnalyzeRounds()
// ga.AddObserver(&observer.Database{})

// go mariogo.RunWebServer()

// ga.Run()

// img := ga.GetCurrentFrame()

// Bild einlesen
// 	img := gocv.IMRead(fmt.Sprintf("images/%v.png", i), gocv.IMReadColor)
// 	if img.Empty() {
// 		fmt.Println("Fehler beim Einlesen des Bildes")
// 		return
// 	}

// for i := 1; i < 13; i++ {
// 	img := gocv.IMRead(fmt.Sprintf("images/%v.png", i), gocv.IMReadColor)
// 	if img.Empty() {
// 		fmt.Println("Fehler beim Einlesen des Bildes")
// 		return
// 	}

// 	gocv.CvtColor(img, &img, gocv.ColorBGRToGray)
// 	gocv.IMWrite(fmt.Sprintf("images/%v_gray.png", i), img)

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

// TEST IMAGE HASH
// img := gocv.IMRead("images/last_p3.png", gocv.IMReadColor)
// if img.Empty() {
// 	fmt.Println("Fehler beim Einlesen des Bildes")
// 	return
// }
// refPx := pixel.PlacementReference[2]
// croppedMat := img.Region(image.Rect(refPx.X-12, refPx.Y-10, refPx.X+20, refPx.Y+40))
// img = croppedMat.Clone()
// gocv.CvtColor(img, &img, gocv.ColorBGRToGray)
// gocv.IMWrite("res.png", img)

// ref := gocv.IMRead("images/12_gray.png", gocv.IMReadColor)
// if ref.Empty() {
// 	fmt.Println("Fehler beim Einlesen des Bildes")
// 	return
// }

// start := time.Now()
// refImg, _ := ref.ToImage()
// image, _ := img.ToImage()
// refImgHash, _ := goimagehash.AverageHash(refImg)
// iamgeHash, _ := goimagehash.AverageHash(image)

// distance, _ := refImgHash.Distance(iamgeHash)
// fmt.Println("Distance:", distance)
// fmt.Println("Elapsed time:", time.Since(start))

// for i := 1; i < 13; i++ {
// 	img := gocv.IMRead(fmt.Sprintf("images/%v_gray.png", i), gocv.IMReadColor)
// 	if img.Empty() {
// 		fmt.Println("Fehler beim Einlesen des Bildes")
// 		return
// 	}

// 	gocv.CvtColor(img, &img, gocv.ColorBGRToGray)

// 	image, _ := img.ToImage()
// 	refImgHash, _ := goimagehash.PerceptionHash(image)
// 	fmt.Printf("\"%v\"\n", refImgHash.GetHash())
// }
