// Object Oriented Programmin in Go

package main

import "errors"

type Point struct {
	x int
	y int
}

func newPoint(x int, y int) *Point {
	point := &Point{x, y}
	return point
}

func (p *Point) Move(dx int, dy int) {
	p.x += dx
	p.y += dy
}

type Square struct {
	Center Point
	Length int
}

func NewSquare(x int, y int, Length int) (*Square, error) {
	if Length <= 0 {
		err := errors.New("Invalid Argument - Length")
		return nil, err
	}
	// point := newPoint(x, y)
	// square := &Square{point, Length}
	// square := {Point(x,y), Length} -> this is another way
	square := &Square{Point{x, y}, Length}
	return square, nil
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func (s *Square) Move(dx int, dy int) {
	s.Center.Move(dx, dy)
}
