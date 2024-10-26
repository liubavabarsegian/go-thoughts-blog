package logo_generator

import "image/color"

var WHITE = color.RGBA{255, 255, 255, 0xff}
var DARK_GREEN = color.RGBA{0, 51, 0, 0xff}
var GREEN = color.RGBA{51, 153, 51, 0xff}
var PINK = color.RGBA{255, 102, 102, 0xff}

var PixelSize = 20

var width = PixelSize * 14
var height = PixelSize * 14

var FrogImagePixels = []Pixel{
	{0, 0, WHITE},
	{0, 1, WHITE},
	{0, 2, WHITE},
	{0, 3, WHITE},
	{0, 4, WHITE},
	{0, 5, WHITE},
	{0, 6, WHITE},
	{0, 7, WHITE},
	{0, 8, WHITE},
	{0, 9, WHITE},
	{0, 10, WHITE},
	{0, 11, WHITE},
	{0, 12, WHITE},
	{0, 13, WHITE},

	{1, 0, WHITE},
	{1, 1, WHITE},
	{1, 2, WHITE},
	{1, 3, WHITE},
	{1, 4, WHITE},
	{1, 5, WHITE},
	{1, 6, WHITE},
	{1, 7, WHITE},
	{1, 8, WHITE},
	{1, 9, WHITE},
	{1, 10, WHITE},
	{1, 11, WHITE},
	{1, 12, WHITE},
	{1, 13, WHITE},

	{2, 0, WHITE},
	{2, 1, WHITE},
	{2, 2, WHITE},
	{2, 3, DARK_GREEN},
	{2, 4, DARK_GREEN},
	{2, 5, WHITE},
	{2, 6, WHITE},
	{2, 7, WHITE},
	{2, 8, WHITE},
	{2, 9, DARK_GREEN},
	{2, 10, DARK_GREEN},
	{2, 11, WHITE},
	{2, 12, WHITE},
	{2, 13, WHITE},

	{3, 0, WHITE},
	{3, 1, WHITE},
	{3, 2, DARK_GREEN},
	{3, 3, GREEN},
	{3, 4, GREEN},
	{3, 5, DARK_GREEN},
	{3, 6, WHITE},
	{3, 7, WHITE},
	{3, 8, DARK_GREEN},
	{3, 9, GREEN},
	{3, 10, GREEN},
	{3, 11, DARK_GREEN},
	{3, 12, WHITE},
	{3, 13, WHITE},

	{4, 0, WHITE},
	{4, 1, WHITE},
	{4, 2, DARK_GREEN},
	{4, 3, GREEN},
	{4, 4, DARK_GREEN},
	{4, 5, GREEN},
	{4, 6, DARK_GREEN},
	{4, 7, DARK_GREEN},
	{4, 8, GREEN},
	{4, 9, DARK_GREEN},
	{4, 10, GREEN},
	{4, 11, DARK_GREEN},
	{4, 12, WHITE},
	{4, 13, WHITE},

	{5, 0, WHITE},
	{5, 1, DARK_GREEN},
	{5, 2, GREEN},
	{5, 3, GREEN},
	{5, 4, DARK_GREEN},
	{5, 5, GREEN},
	{5, 6, GREEN},
	{5, 7, GREEN},
	{5, 8, GREEN},
	{5, 9, DARK_GREEN},
	{5, 10, GREEN},
	{5, 11, GREEN},
	{5, 12, DARK_GREEN},
	{5, 13, WHITE},

	{6, 0, WHITE},
	{6, 1, DARK_GREEN},
	{6, 2, GREEN},
	{6, 3, PINK},
	{6, 4, PINK},
	{6, 5, GREEN},
	{6, 6, GREEN},
	{6, 7, GREEN},
	{6, 8, GREEN},
	{6, 9, PINK},
	{6, 10, PINK},
	{6, 11, GREEN},
	{6, 12, DARK_GREEN},
	{6, 13, WHITE},

	{7, 0, WHITE},
	{7, 1, DARK_GREEN},
	{7, 2, GREEN},
	{7, 3, GREEN},
	{7, 4, GREEN},
	{7, 5, DARK_GREEN},
	{7, 6, DARK_GREEN},
	{7, 7, DARK_GREEN},
	{7, 8, DARK_GREEN},
	{7, 9, GREEN},
	{7, 10, GREEN},
	{7, 11, GREEN},
	{7, 12, DARK_GREEN},
	{7, 13, WHITE},

	{8, 0, WHITE},
	{8, 1, DARK_GREEN},
	{8, 2, GREEN},
	{8, 3, GREEN},
	{8, 4, GREEN},
	{8, 5, GREEN},
	{8, 6, GREEN},
	{8, 7, GREEN},
	{8, 8, GREEN},
	{8, 9, GREEN},
	{8, 10, GREEN},
	{8, 11, GREEN},
	{8, 12, DARK_GREEN},
	{8, 13, WHITE},

	{9, 0, WHITE},
	{9, 1, DARK_GREEN},
	{9, 2, GREEN},
	{9, 3, GREEN},
	{9, 4, GREEN},
	{9, 5, GREEN},
	{9, 6, GREEN},
	{9, 7, GREEN},
	{9, 8, GREEN},
	{9, 9, GREEN},
	{9, 10, GREEN},
	{9, 11, GREEN},
	{9, 12, DARK_GREEN},
	{9, 13, WHITE},

	{10, 0, WHITE},
	{10, 1, DARK_GREEN},
	{10, 2, GREEN},
	{10, 3, GREEN},
	{10, 4, GREEN},
	{10, 5, GREEN},
	{10, 6, GREEN},
	{10, 7, GREEN},
	{10, 8, GREEN},
	{10, 9, GREEN},
	{10, 10, GREEN},
	{10, 11, GREEN},
	{10, 12, DARK_GREEN},
	{10, 13, WHITE},

	{11, 0, WHITE},
	{11, 1, WHITE},
	{11, 2, DARK_GREEN},
	{11, 3, DARK_GREEN},
	{11, 4, DARK_GREEN},
	{11, 5, DARK_GREEN},
	{11, 6, DARK_GREEN},
	{11, 7, DARK_GREEN},
	{11, 8, DARK_GREEN},
	{11, 9, DARK_GREEN},
	{11, 10, DARK_GREEN},
	{11, 11, DARK_GREEN},
	{11, 12, WHITE},
	{11, 13, WHITE},

	{12, 0, WHITE},
	{12, 1, WHITE},
	{12, 2, WHITE},
	{12, 3, WHITE},
	{12, 4, WHITE},
	{12, 5, WHITE},
	{12, 6, WHITE},
	{12, 7, WHITE},
	{12, 8, WHITE},
	{12, 9, WHITE},
	{12, 10, WHITE},
	{12, 11, WHITE},
	{12, 12, WHITE},
	{12, 13, WHITE},

	{13, 0, WHITE},
	{13, 1, WHITE},
	{13, 2, WHITE},
	{13, 3, WHITE},
	{13, 4, WHITE},
	{13, 5, WHITE},
	{13, 6, WHITE},
	{13, 7, WHITE},
	{13, 8, WHITE},
	{13, 9, WHITE},
	{13, 10, WHITE},
	{13, 11, WHITE},
	{13, 12, WHITE},
	{13, 13, WHITE},
}