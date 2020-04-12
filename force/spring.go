package force

import (
	"ganymede/vector"
	"math"
)

type object interface {
	GetAcceleration() vector.Vector
}

// Spring calculates the spring rebound force
func Spring(o object, collisionNormalUnit vector.Vector) vector.Vector {
	currentAccel := o.GetAcceleration()
	collisionAcceleration := currentAccel.Multiply(collisionNormalUnit.Abs())
	var bounceEnergyReturnCoefficient float64 // percentage of energy retained after bounce
	absDotProdAccel := math.Abs(collisionAcceleration.DotProduct(vector.NewVector(1, 1)))
	if absDotProdAccel > 5 {
		bounceEnergyReturnCoefficient = 0.7
	} else if absDotProdAccel > 3 {
		bounceEnergyReturnCoefficient = 0.5
	} else if absDotProdAccel > 1 {
		bounceEnergyReturnCoefficient = 0.1
	}
	impulseVector := collisionAcceleration.Scale(-1 - bounceEnergyReturnCoefficient)
	return impulseVector
}
