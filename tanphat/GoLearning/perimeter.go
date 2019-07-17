package GoLearning

import "math"

type Rectangle struct{
	width float64
	height float64
}

type Circle struct{
	Radius float64
}

type Triangle struct{
	Base float64
	Height float64
}

func (t Triangle) Area() float64{
	return t.Base * t.Height *0.5
}

func (r Rectangle) Area() float64{
	area:= r.width * r.height
	return area
}

func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64{
	return 2 * (rectangle.width + rectangle.height)
}

type Shape interface {
	Area() float64
}

