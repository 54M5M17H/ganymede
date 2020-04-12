package vector

import "math"

var (
	degreesToRadias = math.Pi / 180
)

// NewVector creates a new vector
func NewVector(vals ...float64) Vector {
	return Vector{vals}
}

// Vector is the vector object
type Vector struct {
	vals []float64
}

// GetVals returns the values of the vector
func (v1 Vector) GetVals() []float64 {
	return v1.vals
}

func convertDegreesToRadians(d float64) float64 {
	return d * degreesToRadias
}

// Add will add one vector to another and return the result
func (v1 Vector) Add(v2 Vector) Vector {
	if len(v1.vals) != len(v2.vals) {
		panic("Cannot add vectors that don't represent equal dimensions")
	}

	v := Vector{vals: make([]float64, len(v1.vals))}
	for i, val := range v1.vals {
		v.vals[i] = val + v2.vals[i]
	}
	return v
}

// Subtract will subtract the passed vector from the instance vector and return the result
func (v1 Vector) Subtract(v2 Vector) Vector {
	if len(v1.vals) != len(v2.vals) {
		panic("Cannot subtract vectors that don't represent equal dimensions")
	}

	v := Vector{vals: make([]float64, len(v1.vals))}
	for i, val := range v1.vals {
		v.vals[i] = val - v2.vals[i]
	}
	return v
}

// Multiply multiplies two vectors
func (v1 Vector) Multiply(v2 Vector) Vector {
	if len(v1.vals) != len(v2.vals) {
		panic("Cannot subtract vectors that don't represent equal dimensions")
	}
	v := Vector{vals: make([]float64, len(v1.vals))}
	for i, val := range v1.vals {
		v.vals[i] = val * v2.vals[i]
	}
	return v
}

// Scale multiplies the vector by the scalar provided and returns the result
func (v1 Vector) Scale(scalar float64) Vector {
	v := Vector{vals: make([]float64, len(v1.vals))}
	for i, val := range v1.vals {
		v.vals[i] = val * scalar
	}
	return v
}

// CrossProduct performs the cross product of 2 vectors and returns the result
func (v1 Vector) CrossProduct(v2 Vector) Vector {
	// magnitude of v1 * magnitude of v2 sin(angle between them)

	// don't know how to implement for 2D
	return Vector{}
}

// DotProduct performs the dot product of 2 vectors and returns the result
func (v1 Vector) DotProduct(v2 Vector) float64 {
	if len(v1.vals) != len(v2.vals) {
		panic("Cannot take dot product vectors that don't represent equal dimensions")
	}

	var res float64 = 0
	for i, val := range v1.vals {
		res += val * v2.vals[i]
	}
	return res
}

// RotateAboutTail rotates the vector about its tail
func (v1 Vector) RotateAboutTail(clockWiseAngleInRadians float64) Vector {
	if len(v1.vals) != 2 {
		panic("Rotate only implemented for 2D vectors")
	}

	// sin & cosin expect anti-clockwise so negate
	cos := math.Cos(-clockWiseAngleInRadians)
	sin := math.Sin(-clockWiseAngleInRadians)

	x1 := v1.vals[0]
	y1 := v1.vals[1]

	x2 := (cos * x1) - (sin * y1)
	y2 := (sin * x1) + (cos * y1)

	return Vector{vals: []float64{x2, y2}}
}

// AsUnitVector converts any vector to a unit vector (magnitudes of 1)
func (v1 Vector) AsUnitVector() (unitV Vector) {
	for _, val := range v1.vals {
		var v float64
		if val > 0 {
			v = 1
		} else if val < 0 {
			v = -1
		}

		// else remains as 0
		unitV.vals = append(unitV.vals, v)
	}
	return
}

// Abs returns a vector with absolute magnitudes
func (v1 Vector) Abs() Vector {
	v := Vector{vals: make([]float64, len(v1.vals))}
	for i, val := range v1.vals {
		v.vals[i] = math.Abs(val)
	}
	return v
}
