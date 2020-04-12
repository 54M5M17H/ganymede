package object

import (
	"ganymede/vector"
)

// Object interface is implemented by all objects
type Object interface {
	GetMass() float64
	GetPosition() vector.Vector
	GetAcceleration() vector.Vector
	ApplyAcceleration(vector.Vector)
	RotateAcceleration(float64)
}

// NewGenericObject creates a generic object
func NewGenericObject(mass float64, position vector.Vector, collisionType collisionType) GenericObject {
	acceleration := initAcceleration(position)
	return GenericObject{mass, position, collisionType, acceleration}
}

func initAcceleration(v vector.Vector) vector.Vector {
	a := []float64{}
	for range v.GetVals() {
		a = append(a, 0)
	}
	return vector.NewVector(a...)
}

// GenericObject is a generic object implementing the object interface
type GenericObject struct {
	mass          float64
	position      vector.Vector
	collisionType collisionType
	acceleration  vector.Vector
}

// GetMass returns the mass of the object
func (o *GenericObject) GetMass() float64 {
	return o.mass
}

// ApplyAcceleration allows you to apply acceleration without factoring in the mass.
// This might be useful for player interaction or impulse resolution.
func (o *GenericObject) ApplyAcceleration(acceleration vector.Vector) {
	o.acceleration = o.acceleration.Add(acceleration)
	o.move()
}

// RotateAcceleration allows you to change the direction of the acceleration
// without changing its magnitude.
func (o *GenericObject) RotateAcceleration(radians float64) {
	o.acceleration = o.acceleration.RotateAboutTail(radians)
	o.move()
}

// AdjustPosition changes an objects position without creating acceleration.
// This is useful for changes that shouldn't apply repetitively.
// It's also useful for making a position change that doesn't also apply existing acceleration.
// For example slightly moving the object following a collision.
func (o *GenericObject) AdjustPosition(v vector.Vector) {
	o.position = o.position.Add(v)
}

func (o *GenericObject) move() {
	o.position = o.position.Add(o.acceleration)
}

// GetPosition returns the position of the object as a vector
func (o *GenericObject) GetPosition() vector.Vector {
	return o.position
}

// DetectCollision returns true if the object
func (o *GenericObject) DetectCollision(v vector.Vector) bool {
	panic("Not implemented")
	// return false
}

// GetCollisionType returns the objects collision type.
// Implements collision.collider interface
func (o *GenericObject) GetCollisionType() collisionType {
	return o.collisionType
}

// GetAcceleration returns the acceleration vector
func (o *GenericObject) GetAcceleration() vector.Vector {
	return o.acceleration
}

// CollisionOverlapCorrection corrects overlap between colliding objects
func (o *GenericObject) CollisionOverlapCorrection(collisionNormal, objectDimensions vector.Vector) {
	collisionNormalUnit := collisionNormal.AsUnitVector()
	desiredResetDistance := objectDimensions.Multiply(collisionNormalUnit)
	adjustment := desiredResetDistance.Subtract(collisionNormal)
	o.AdjustPosition(adjustment)
}
