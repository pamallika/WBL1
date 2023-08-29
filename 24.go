package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) GetX() int {
	return p.x
}

func (p *Point) GetY() int {
	return p.y
}

func (p *Point) SetX(x int) {
	p.x = x
}

func (p *Point) SetY(y int) {
	p.x = y
}

func Distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.x-p1.x), 2.0) + math.Pow(float64(p2.y-p1.y), 2.0))
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(2, 1)
	fmt.Println("First Point:", *p1)
	fmt.Println("Second Point:", *p2)
	fmt.Println("Distance:", Distance(*p1, *p2))
}
