package domain

// FractalImage представляет изображение фрактала
type FractalImage struct {
	Width  int
	Height int
	Pixels []Pixel
}

// NewFractalImage возвращает новый экземпляр FractalImage с указанными шириной и высотой
func NewFractalImage(width, height int) *FractalImage {
	pixels := make([]Pixel, width*height)
	return &FractalImage{Width: width, Height: height, Pixels: pixels}
}

// GetPixel возвращает пиксель, который находится в точке x,y, и флаг того найден он или нет
func (f *FractalImage) GetPixel(x, y int) (*Pixel, bool) {
	if !f.contains(x, y) {
		return nil, false
	}

	return &f.Pixels[y*f.Width+x], true
}

// contains проверяет находится ли пиксель в пределах изображения
func (f *FractalImage) contains(x, y int) bool {
	return x >= 0 && y >= 0 && x < f.Width && y < f.Height
}
