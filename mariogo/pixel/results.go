package pixel

import "image/color"

var PlayerColors = [4]color.RGBA{
	{250, 229, 38, 255},  //P1
	{33, 229, 251, 255},  //P2
	{253, 116, 116, 255}, //P3
	{115, 242, 40, 255},  //P4
}

var NeutralResultP1 = []Pixel{
	{X: 314, Y: 175, C: color.RGBA{9, 190, 2, 255}},
	{X: 326, Y: 181, C: color.RGBA{70, 221, 0, 255}},
	{X: 338, Y: 176, C: color.RGBA{20, 191, 0, 255}},
}

var PositiveResultP1 = []Pixel{
	{X: 324, Y: 167, C: color.RGBA{0, 171, 255, 255}},
	{X: 326, Y: 188, C: color.RGBA{0, 74, 255, 255}},
	{X: 316, Y: 177, C: color.RGBA{0, 125, 255, 255}},
	{X: 339, Y: 177, C: color.RGBA{0, 120, 255, 255}},
	{X: 326, Y: 177, C: color.RGBA{0, 121, 255, 255}},
}

var NegativeResultP1 = []Pixel{
	{X: 326, Y: 168, C: color.RGBA{253, 50, 14, 255}},
	{X: 315, Y: 178, C: color.RGBA{255, 106, 0, 255}},
	{X: 338, Y: 179, C: color.RGBA{255, 108, 0, 255}},
	{X: 327, Y: 189, C: color.RGBA{255, 170, 0, 255}},
	{X: 327, Y: 176, C: color.RGBA{255, 90, 0, 255}},
}

var EndResults = []Pixel{
	{X: 12, Y: 15, C: color.RGBA{246, 0, 0, 255}},
	{X: 23, Y: 81, C: color.RGBA{205, 0, 0, 255}},
	{X: 35, Y: 62, C: color.RGBA{255, 43, 41, 255}},
	{X: 128, Y: 21, C: color.RGBA{246, 0, 0, 255}},
	{X: 311, Y: 23, C: color.RGBA{245, 0, 0, 255}},
	{X: 411, Y: 5, C: color.RGBA{248, 1, 0, 255}},
	{X: 551, Y: 13, C: color.RGBA{247, 0, 0, 255}},
	{X: 234, Y: 122, C: color.RGBA{172, 0, 0, 255}},
	{X: 147, Y: 121, C: color.RGBA{174, 0, 0, 255}},
	{X: 601, Y: 121, C: color.RGBA{174, 0, 0, 255}},
	{X: 646, Y: 82, C: color.RGBA{203, 0, 0, 255}},
	{X: 668, Y: 66, C: color.RGBA{199, 0, 0, 255}},
	{X: 674, Y: 8, C: color.RGBA{242, 0, 0, 255}},
	{X: 716, Y: 49, C: color.RGBA{235, 0, 0, 255}},
	{X: 722, Y: 85, C: color.RGBA{204, 1, 0, 255}},
	{X: 694, Y: 120, C: color.RGBA{167, 1, 0, 255}},
}

var BotPlusP1 = []Pixel{
	{X: 818, Y: 68, C: color.RGBA{231, 224, 223, 255}},
	{X: 819, Y: 71, C: color.RGBA{231, 226, 220, 255}},
	{X: 815, Y: 75, C: color.RGBA{238, 231, 229, 255}},
	{X: 813, Y: 75, C: color.RGBA{238, 231, 241, 255}},
	{X: 818, Y: 78, C: color.RGBA{233, 226, 222, 255}},
	{X: 819, Y: 81, C: color.RGBA{236, 234, 235, 255}},
	{X: 822, Y: 75, C: color.RGBA{230, 232, 239, 255}},
	{X: 824, Y: 75, C: color.RGBA{239, 233, 230, 255}},
}

var CharacterThumbnailP1 = []Pixel{
	{X: 430, Y: 65},
	{X: 445, Y: 90},
}
