package random

// Random представляет абстракцию рандомайзера, чтобы было удобнее тестировать
type Random interface {
	Float64() float64
	Intn(n int) int
}
