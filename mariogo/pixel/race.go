package pixel

import "image/color"

var OnePlayerPlaying = []Pixel{
	{X: 171, Y: 657, C: color.RGBA{8, 8, 8, 255}},
	{X: 176, Y: 660, C: color.RGBA{4, 5, 4, 255}},
	{X: 183, Y: 660, C: color.RGBA{239, 239, 239, 255}},
	{X: 191, Y: 659, C: color.RGBA{8, 10, 9, 255}},
	{X: 190, Y: 667, C: color.RGBA{230, 231, 230, 255}},
	{X: 181, Y: 667, C: color.RGBA{10, 10, 10, 255}},
	{X: 174, Y: 667, C: color.RGBA{227, 227, 227, 255}},
	{X: 173, Y: 675, C: color.RGBA{8, 8, 8, 255}},
	{X: 181, Y: 674, C: color.RGBA{224, 224, 224, 255}},
	{X: 189, Y: 675, C: color.RGBA{8, 8, 8, 255}},
	{X: 244, Y: 666, C: color.RGBA{249, 249, 248, 255}},
	{X: 237, Y: 678, C: color.RGBA{221, 223, 219, 255}},
	{X: 167, Y: 681, C: color.RGBA{4, 4, 4, 255}},
}

var TwoPlayerPlaying = []Pixel{
	{X: 132, Y: 670, C: color.RGBA{14, 14, 14, 255}},
	{X: 135, Y: 672, C: color.RGBA{9, 10, 9, 255}},
	{X: 140, Y: 672, C: color.RGBA{242, 242, 242, 255}},
	{X: 145, Y: 671, C: color.RGBA{9, 10, 9, 255}},
	{X: 144, Y: 677, C: color.RGBA{231, 231, 231, 255}},
	{X: 139, Y: 677, C: color.RGBA{15, 15, 15, 255}},
	{X: 134, Y: 677, C: color.RGBA{223, 223, 223, 255}},
	{X: 133, Y: 682, C: color.RGBA{11, 11, 11, 255}},
	{X: 138, Y: 682, C: color.RGBA{220, 220, 219, 255}},
	{X: 144, Y: 682, C: color.RGBA{12, 12, 12, 255}},
	{X: 772, Y: 670, C: color.RGBA{14, 14, 14, 255}},
	{X: 774, Y: 672, C: color.RGBA{4, 4, 4, 255}},
	{X: 780, Y: 672, C: color.RGBA{242, 242, 242, 255}},
	{X: 786, Y: 671, C: color.RGBA{10, 10, 10, 255}},
	{X: 784, Y: 677, C: color.RGBA{231, 231, 231, 255}},
	{X: 779, Y: 677, C: color.RGBA{15, 15, 15, 255}},
	{X: 774, Y: 677, C: color.RGBA{223, 223, 223, 255}},
	{X: 773, Y: 682, C: color.RGBA{11, 11, 11, 255}},
	{X: 779, Y: 682, C: color.RGBA{210, 211, 210, 255}},
	{X: 784, Y: 682, C: color.RGBA{12, 12, 12, 255}},
}

var ThreePlayerPlaying = []Pixel{
	{X: 116, Y: 318, C: color.RGBA{9, 5, 5, 255}},
	{X: 120, Y: 318, C: color.RGBA{236, 236, 236, 255}},
	{X: 125, Y: 318, C: color.RGBA{11, 12, 11, 255}},
	{X: 124, Y: 322, C: color.RGBA{224, 224, 224, 255}},
	{X: 120, Y: 322, C: color.RGBA{16, 17, 16, 255}},
	{X: 116, Y: 322, C: color.RGBA{224, 224, 220, 255}},
	{X: 114, Y: 326, C: color.RGBA{5, 6, 5, 255}},
	{X: 119, Y: 326, C: color.RGBA{224, 225, 224, 255}},
	{X: 123, Y: 327, C: color.RGBA{13, 13, 13, 255}},
	{X: 1179, Y: 317, C: color.RGBA{9, 8, 8, 255}},
	{X: 1184, Y: 318, C: color.RGBA{239, 236, 239, 255}},
	{X: 1189, Y: 318, C: color.RGBA{4, 7, 8, 255}},
	{X: 1188, Y: 322, C: color.RGBA{227, 226, 226, 255}},
	{X: 1184, Y: 322, C: color.RGBA{10, 13, 13, 255}},
	{X: 1179, Y: 322, C: color.RGBA{232, 232, 231, 255}},
	{X: 1178, Y: 326, C: color.RGBA{5, 6, 5, 255}},
	{X: 1183, Y: 326, C: color.RGBA{222, 224, 223, 255}},
	{X: 1187, Y: 327, C: color.RGBA{7, 7, 6, 255}},
	{X: 115, Y: 678, C: color.RGBA{13, 14, 13, 255}},
	{X: 116, Y: 682, C: color.RGBA{234, 233, 232, 255}},
	{X: 120, Y: 682, C: color.RGBA{12, 12, 12, 255}},
	{X: 124, Y: 682, C: color.RGBA{235, 235, 235, 255}},
	{X: 124, Y: 686, C: color.RGBA{16, 16, 16, 255}},
	{X: 119, Y: 686, C: color.RGBA{227, 228, 227, 255}},
	{X: 114, Y: 687, C: color.RGBA{7, 8, 9, 255}},
}

var FourPlayerPlaying = append([]Pixel{
	{X: 1179, Y: 678, C: color.RGBA{2, 3, 3, 255}},
	{X: 1184, Y: 678, C: color.RGBA{239, 241, 242, 255}},
	{X: 1189, Y: 678, C: color.RGBA{2, 4, 4, 255}},
	{X: 1188, Y: 682, C: color.RGBA{225, 231, 230, 255}},
	{X: 1184, Y: 682, C: color.RGBA{6, 5, 7, 255}},
	{X: 1180, Y: 682, C: color.RGBA{229, 230, 229, 255}},
	{X: 1178, Y: 687, C: color.RGBA{5, 7, 6, 255}},
	{X: 1183, Y: 686, C: color.RGBA{221, 219, 219, 255}},
	{X: 1188, Y: 687, C: color.RGBA{6, 14, 11, 255}},
}, ThreePlayerPlaying...)

var IntroPage = []Pixel{
	{X: 109, Y: 596, C: color.RGBA{214, 16, 10, 255}},
	{X: 160, Y: 585, C: color.RGBA{234, 194, 45, 255}},
	{X: 145, Y: 609, C: color.RGBA{255, 254, 86, 255}},
	{X: 140, Y: 600, C: color.RGBA{255, 255, 87, 255}},
	{X: 125, Y: 601, C: color.RGBA{255, 255, 80, 255}},
	{X: 119, Y: 612, C: color.RGBA{255, 250, 95, 255}},
	{X: 133, Y: 595, C: color.RGBA{219, 13, 11, 255}},
	{X: 132, Y: 617, C: color.RGBA{194, 0, 0, 255}},
	{X: 144, Y: 644, C: color.RGBA{254, 212, 4, 255}},
	{X: 166, Y: 627, C: color.RGBA{255, 209, 0, 255}},
	{X: 90, Y: 638, C: color.RGBA{236, 164, 0, 255}},
	{X: 90, Y: 586, C: color.RGBA{245, 206, 70, 255}},
	{X: 159, Y: 569, C: color.RGBA{245, 210, 93, 255}},
	{X: 231, Y: 584, C: color.RGBA{241, 215, 60, 255}},
	{X: 229, Y: 638, C: color.RGBA{242, 178, 0, 255}},
	{X: 161, Y: 666, C: color.RGBA{232, 181, 0, 255}},
	{X: 161, Y: 645, C: color.RGBA{227, 184, 2, 255}},
	{X: 350, Y: 620, C: color.RGBA{82, 85, 72, 255}},
	{X: 936, Y: 662, C: color.RGBA{81, 82, 61, 255}},
	{X: 254, Y: 585, C: color.RGBA{243, 241, 242, 255}},
	{X: 261, Y: 601, C: color.RGBA{234, 234, 232, 255}},
	{X: 269, Y: 584, C: color.RGBA{249, 253, 252, 255}},
	{X: 282, Y: 587, C: color.RGBA{245, 246, 245, 255}},
	{X: 280, Y: 594, C: color.RGBA{243, 244, 241, 255}},
	{X: 278, Y: 602, C: color.RGBA{236, 235, 235, 255}},
	{X: 295, Y: 590, C: color.RGBA{238, 238, 238, 255}},
	{X: 295, Y: 598, C: color.RGBA{235, 236, 233, 255}},
	{X: 309, Y: 591, C: color.RGBA{243, 245, 239, 255}},
	{X: 320, Y: 598, C: color.RGBA{225, 225, 230, 255}},
}
