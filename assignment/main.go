package main

import "fmt"

type shape interface {
	getArea() float64
	printArea()
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	side float64
}

func main() {

	t := triangle{4.0, 6.0}
	s := square{5.0}

	t.printArea()
	s.printArea()

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func (t triangle) printArea() {
	fmt.Println(t.getArea())
}

func (s square) printArea() {
	fmt.Println(s.getArea())
}
