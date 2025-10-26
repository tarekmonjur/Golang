package main

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

type shape interface {
	getArea() float64
}

func main() {
	sq := square{sideLength: 5}
	tr := triangle{base: 10, height: 5}

	printArea(sq)
	printArea(tr)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(sh shape) {
	println("Area: ", sh.getArea())
}
