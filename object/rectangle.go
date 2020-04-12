package object

import (
	"ganymede/vector"
)

// NewRectangleObject creates a new rectangle
func NewRectangleObject(w float64, h float64, mass float64, position vector.Vector) Rectangle {
	return Rectangle{
		vector.NewVector(w, h),
		NewGenericObject(mass, position, collisionBoundingBox),
	}
}

// Rectangle is an object with physical implementation for a 2D rectangle
type Rectangle struct {
	dimensions vector.Vector
	GenericObject
}

// GetDimensions returns a dimensions tuple
func (r Rectangle) GetDimensions() vector.Vector {
	return r.dimensions
}
