package object

import (
	"ganymede/vector"
	"math"
)

// Type defines an objects collision type
type collisionType = int

const (
	collisionCircle = iota
	collisionBoundingBox
)

type collider interface {
	GetCollisionType() collisionType
	GetPosition() vector.Vector
}

type boundingBoxCollider interface {
	GetDimensions() vector.Vector
	collider
}

type circleCollider interface {
	GetRadius() float64
	collider
}

// DetectCollision returns true if the two objects have collided
func DetectCollision(o1 collider, o2 collider) (bool, vector.Vector) {
	o1Type := o1.GetCollisionType()

	switch o1Type {
	case collisionCircle:
		c1 := o1.(circleCollider)
		return circleAnd(c1, o2)
	case collisionBoundingBox:
		b1 := o1.(boundingBoxCollider)
		return boundingBoxAnd(b1, o2)
	default:
		panic("Unknown collision type")
	}
}

func circleAnd(c1 circleCollider, o2 collider) (bool, vector.Vector) {
	o2Type := o2.GetCollisionType()
	switch o2Type {
	case collisionCircle:
		c2 := o2.(circleCollider)
		return circleAndCircle(c1, c2)
	case collisionBoundingBox:
		b2 := o2.(boundingBoxCollider)
		return circleAndBB(c1, b2)
	default:
		panic("Unknown collision type")
	}
}

func boundingBoxAnd(b1 boundingBoxCollider, o2 collider) (bool, vector.Vector) {
	o2Type := o2.GetCollisionType()
	switch o2Type {
	case collisionCircle:
		c2 := o2.(circleCollider)
		return circleAndBB(c2, b1)
	case collisionBoundingBox:
		b2 := o2.(boundingBoxCollider)
		return bBAndBB(b1, b2)
	default:
		panic("Unknown collision type")
	}
}

func bBAndBB(b1 boundingBoxCollider, b2 boundingBoxCollider) (bool, vector.Vector) {
	topLeft1 := b1.GetPosition()
	topLeft2 := b2.GetPosition()

	bottomRight1 := boundingBoxBottomRight(b1)
	bottomRight2 := boundingBoxBottomRight(b2)

	// if top left of one is lower than bottom right of other, or vice versa, no overlap
	// if top left of one is to right of bottom right of other, or vice versa, no overlap

	oneAboveTwo := topLeft1.Subtract(bottomRight2)
	twoAboveOne := topLeft2.Subtract(bottomRight1)

	for _, val := range append(oneAboveTwo.GetVals(), twoAboveOne.GetVals()...) {
		if val > 0 {
			return false, vector.Vector{}
		}
	}

	return true, vector.Vector{}
}

func boundingBoxBottomRight(b boundingBoxCollider) vector.Vector {
	return b.GetPosition().Add(b.GetDimensions())
}

func circleAndCircle(c1 circleCollider, c2 circleCollider) (bool, vector.Vector) {
	maxDistance := c1.GetRadius() + c2.GetRadius()
	collided := !distanceBetweenPointsIsGreaterThan(c1.GetPosition(), c2.GetPosition(), maxDistance)
	return collided, c1.GetPosition().Subtract(c2.GetPosition())
}

func distanceBetweenPointsIsGreaterThan(p1, p2 vector.Vector, distance float64) bool {
	dSq := math.Pow(distance, 2)
	diff := p1.Subtract(p2)
	distanceSq := diff.DotProduct(diff)
	return distanceSq > dSq
}

func circleAndBB(c circleCollider, b boundingBoxCollider) (bool, vector.Vector) {
	// is circle centre inside box?
	pointNearestToCentre := nearestBoundingBoxEdge(c.GetPosition(), b)
	diffVector := c.GetPosition().Subtract(pointNearestToCentre)
	if isPointInsideBox(c.GetPosition(), b) {
		return true, diffVector.Scale(-1)
	}

	collided := !distanceBetweenPointsIsGreaterThan(c.GetPosition(), pointNearestToCentre, c.GetRadius())
	return collided, diffVector
}

func isPointInsideBox(point vector.Vector, b boundingBoxCollider) bool {
	bbTopLeft := b.GetPosition()
	bbBottomRight := boundingBoxBottomRight(b)

	// is it above or left of top left?
	// is it below or right of bottom right?
	// if yes to any, not inside the box

	vToTopLeft := point.Subtract(bbTopLeft)
	vToBottomRight := bbBottomRight.Subtract(point)

	for _, val := range append(vToTopLeft.GetVals(), vToBottomRight.GetVals()...) {
		if val < 0 {
			return false
		}
	}

	return true
}

func nearestBoundingBoxEdge(point vector.Vector, b boundingBoxCollider) vector.Vector {
	topLeft := b.GetPosition()
	bottomRight := boundingBoxBottomRight(b)

	p := point.GetVals()
	vToBottomRight := bottomRight.Subtract(point).GetVals() // if negative, point was bigger
	vToTopLeft := point.Subtract(topLeft).GetVals()         // if negative, point was smaller
	nearestPointCoords := []float64{}

	// if outside of box edge, fix coord to that edge
	for i := range vToBottomRight {
		var newCoord float64
		if vToBottomRight[i] < 0 {
			newCoord = bottomRight.GetVals()[i]
		} else if vToTopLeft[i] < 0 {
			newCoord = topLeft.GetVals()[i]
		} else {
			newCoord = p[i]
		}
		nearestPointCoords = append(nearestPointCoords, newCoord)
	}

	return vector.NewVector(nearestPointCoords...)
}
