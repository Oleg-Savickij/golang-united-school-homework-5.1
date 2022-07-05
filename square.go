package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (receiver Square) End() Point {
	var endPoint Point = receiver.start
	endPoint.x = endPoint.x + int(receiver.a)
	endPoint.y = endPoint.y + int(receiver.a)

	return endPoint
}

func (receiver Square) Area() uint {
	return receiver.a * receiver.a
}

func (receiver Square) Perimeter() uint {
	return receiver.a * 4
}
