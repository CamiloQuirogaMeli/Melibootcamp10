package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func newGeometry(geoType string, values ...float64) geometry {
	switch geoType {
	case "rect":
		return rect{width: values[0], height: values[1]}
	case "circle":
		return circle{radius: values[0]}
	}
	return nil
}
func main() {
	const (
		rect   = "rect"
		circle = "circle"
	)
	r := newGeometry(rect, 2, 3)
	fmt.Println("rect")
	fmt.Println(r.area())
	fmt.Println(r.perim())
	c := newGeometry(circle, 2)
	fmt.Println("circle")
	fmt.Println(c.area())
	fmt.Println(c.perim())
}
