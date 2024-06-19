package mariogo

import (
	"fmt"
	"image"
	"image/color"
	"jkli/mariogo/mariogo/pixel"
	"log"
	"time"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/otiai10/gosseract/v2"
	"gocv.io/x/gocv"
)

type Capture struct {
	Frame            *gocv.Mat
	webcam           *gocv.VideoCapture
	webcamRetryTime  time.Duration
	maxPixelDistance float64
	frameDeviations  float64
}

func NewCapture() *Capture {
	c := Capture{
		webcamRetryTime: 1 * time.Second,
	}

	c.DefaultMatchSetting()

	return &c
}

func (c *Capture) ConnectWebcam() error {

	if c.webcam != nil {
		c.webcam.Close()
	}

	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Println("Failed to connect webcam stream")
	}
	c.webcam = webcam

	return err
}

func (c *Capture) Stop() {
	defer c.webcam.Close()
}

func (c *Capture) readWebcamWithRetry() *gocv.Mat {
	img := gocv.NewMat()

	if c.webcam == nil {
		for err := c.ConnectWebcam(); err != nil; {
			time.Sleep(c.webcamRetryTime)
		}
	}

	for !c.webcam.IsOpened() {
		// TODO: maybe reconnect?
		time.Sleep(c.webcamRetryTime)
	}

	for ok := c.webcam.Read(&img); !ok; { // TODO FIx reconnecting
		fmt.Println("Retrying webcam connection...")
		c.ConnectWebcam()
		time.Sleep(c.webcamRetryTime)
	}

	return &img

}

func (c *Capture) NextFrame() {

	if c.Frame != nil {
		c.Frame.Close()
	}

	img := c.readWebcamWithRetry()

	gocv.Resize(*img, img, image.Point{X: 1280, Y: 720}, 0, 0, gocv.InterpolationLinear)

	c.Frame = img

}

func (c *Capture) MatchSetting(frameDeviations float64, maxPixelDistance float64) {
	c.frameDeviations = frameDeviations
	c.maxPixelDistance = maxPixelDistance
}

func (c *Capture) DefaultMatchSetting() {
	c.frameDeviations = 0.8
	c.maxPixelDistance = 0.1
}

func (c *Capture) Matches(pixels []pixel.Pixel) bool {
	matchCount := 0

	for _, pixel := range pixels {
		vec := c.Frame.GetVecbAt(pixel.Y, pixel.X)

		color := color.RGBA{vec[2], vec[1], vec[0], 255}
		color1, ok1 := colorful.MakeColor(pixel.C)
		color2, ok2 := colorful.MakeColor(color)

		if !ok1 || !ok2 {
			fmt.Printf("Error converting color: %v\n", vec)
			continue
		}

		if color1.DistanceCIEDE2000(color2) > c.maxPixelDistance {
			continue
		}

		matchCount++
	}

	matchPercentage := float64(matchCount) / float64(len(pixels))

	return matchPercentage > c.frameDeviations
}

func (c *Capture) OCR(x0, y0, x1, y1 int) (out string, err error) {

	crop := c.Crop(x0, y0, x1, y1)

	imageName := fmt.Sprintf("images/temp/%v.png", time.Now().Unix())
	log.Println(imageName)
	gocv.IMWrite(imageName, *crop)

	client := gosseract.NewClient()
	defer client.Close()

	client.SetLanguage("deu")
	client.SetImage(imageName)

	return client.Text()
}

func (c *Capture) Crop(x0, y0, x1, y1 int) *gocv.Mat {
	croppedMat := c.Frame.Region(image.Rect(x0, y0, x1, y1))
	playerImg := croppedMat.Clone()
	return &playerImg
}
