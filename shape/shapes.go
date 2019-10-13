package shape

import "math"

type Shape interface {
	Volume() float64
	Area() float64
}

type Cube struct {
	Side float64
}

type RectangularBox struct {
	Length float64
	Width  float64
	Height float64
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Side, 3)
}

func (b RectangularBox) Volume() float64 {
	return b.Height * b.Length * b.Width
}
func (c Cube) Area() float64 {
	return 6 * math.Pow(c.Side, 2)
}

func (b RectangularBox) Area() float64 {
	return 2 * (b.Height*b.Length + b.Height*b.Width + b.Length*b.Width)
}
