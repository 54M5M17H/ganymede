package object

import (
	"ganymede/vector"
)

// NewCircleObject creates a new circle
func NewCircleObject(r float64, mass float64, position vector.Vector) Circle {
	return Circle{
		r,
		NewGenericObject(mass, position, collisionCircle),
	}
}

// Circle is an object with physical implementation for a 2D circle
type Circle struct {
	Radius float64
	GenericObject
}

// GetRadius returns circle radius. Implements collision.circleCollider
func (c Circle) GetRadius() float64 {
	return c.Radius
}
