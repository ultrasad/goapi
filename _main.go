package main

import "fmt"

// Shaper is a representation of a shaper
type Shaper interface {
	Area() float64
}

//Rectangle is a representation of a rectangle
type Rectangle struct {
	width  float64
	height float64
}

//Area is function of rectangle
func (rectangle Rectangle) Area() float64 {
	return rectangle.width * rectangle.height
}

//Square is square
type Square struct {
	width float64
}

//Area is function of rectangle
func (square Square) Area() float64 {
	return square.width * square.width
}

//Circle is circle
type Circle struct {
	redius float64
}

//Area is function of rectangle
func (circle Circle) Area() float64 {
	return circle.redius * circle.redius * 3.1415
}

//Unknow is unknow type
type Unknow struct {
	string string
}

//Area unknow
func (unknow Unknow) Area() float64 {
	return 3.1415
}

//ComputeArea is function of compute area
func ComputeArea(shape Shaper) {
	fmt.Printf("Area of %+v=%.3f\n", shape, shape.Area())
}

func main() {
	rectangle := Rectangle{5, 10}
	square := Square{5}
	circle := Circle{5}
	unknow := Unknow{"test unknow"}

	//fmt.Printf("width %+v \n", rectangle)

	shapes := [...]Shaper{rectangle, square, circle, unknow}
	for _, shape := range shapes {
		ComputeArea(shape)
	}

}
