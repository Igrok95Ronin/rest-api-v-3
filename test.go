package main

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) us() (int, int) {
	p.X = 44
	p.Y = 55
	return p.X, p.Y
}

func main() {
	p := Point{
		X: 1,
		Y: 3,
	}
	fmt.Println(p)
	x, y := p.us()
	fmt.Println(x, y)
}
