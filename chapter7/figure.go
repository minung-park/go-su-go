package main

import (
	f "fmt"
	"math"
)

// interface
type geometry interface {
	area() float64  // method
	perim() float64 // method
}

// struct
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	f.Println("========== method value ==========")
	rf := r.scaleup
	f.Println("rectangle scale up area: ", rf(2.0))

	f.Println("========== method expression ==========")
	cf := circle.scaleup
	f.Println("circle scale up area: ", cf(c, 2.0))

	f.Println("========== interface ==========")
	measure(r)
	measure(c)

}

// method
func (r rect) scaleup(val float64) float64 {
	return r.area() * val * val
}
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) scaleup(val float64) float64 {
	return c.area() * val * val
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// function
func measure(g geometry) {
	f.Println("info::::", g)
	f.Println(g.area())
	f.Println(g.perim())
}
