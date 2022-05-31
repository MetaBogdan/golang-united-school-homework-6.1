package golang_united_school_homework

import "errors"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity == len(b.shapes) {
		return errors.New("shape out of the box capacity")
	}

	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("index not found")
	}

	for k, val := range b.shapes {
		if k == i {
			return val, nil
		}
	}
	return nil, errors.New("index not found")

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape_i, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return shape_i, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	_, err := b.GetByIndex(i)
	if err != nil {
		return nil, err
	}
	copySh := make([]Shape, len(b.shapes))
	copy(copySh, b.shapes)
	b.shapes[i] = shape

	return copySh[i], nil

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {

	var sumPerimeter float64
	for _, shape := range b.shapes {
		sumPerimeter += shape.CalcPerimeter()
	}
	return sumPerimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {

	var sumArea float64
	for _, shape := range b.shapes {

		sumArea += shape.CalcArea()
	}
	return sumArea
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {

	var crShapes []Shape
	for _, shape := range b.shapes {
		switch shape.(type) {
		case *Circle:
			continue
		default:
			crShapes = append(crShapes, shape)
		}

	}
	if len(b.shapes) == len(crShapes) {
		return errors.New("circle not found")
	}
	b.shapes = crShapes

	return nil

}
