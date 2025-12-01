package domain

type Pixel struct {
	Color    Color
	HitCount uint32
	Normal   float64
}

// ColorPixel изменяет цвет пикселя на основе количества попаданий в него
func (pixel *Pixel) ColorPixel(color Color) {
	if pixel.HitCount == 0 {
		pixel.Color.R = color.R
		pixel.Color.G = color.G
		pixel.Color.B = color.B
	} else {
		pixel.Color.R = (pixel.Color.R + color.R) / 2
		pixel.Color.B = (pixel.Color.B + color.B) / 2
		pixel.Color.G = (pixel.Color.G + color.G) / 2
	}
	pixel.HitCount++
}
