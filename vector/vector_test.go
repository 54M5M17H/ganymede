package vector

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVector(t *testing.T) {
	Convey("Should add two vectors", t, func() {
		v1 := Vector{vals: []float64{1, 2, 3}}
		v2 := Vector{vals: []float64{4, 5, 6}}
		res1 := v1.Add(v2)
		So(res1.vals[0], ShouldEqual, 5)
		So(res1.vals[1], ShouldEqual, 7)
		So(res1.vals[2], ShouldEqual, 9)

		v3 := Vector{vals: []float64{9, 8, 7}}
		v4 := Vector{vals: []float64{-11, 5, 0}}
		res2 := v3.Add(v4)
		So(res2.vals[0], ShouldEqual, -2)
		So(res2.vals[1], ShouldEqual, 13)
		So(res2.vals[2], ShouldEqual, 7)
	})

	Convey("Should substract one vector from another", t, func() {
		v1 := Vector{vals: []float64{7, 3, 9}}
		v2 := Vector{vals: []float64{4, 5, 6}}
		res1 := v1.Subtract(v2)
		So(res1.vals[0], ShouldEqual, 3)
		So(res1.vals[1], ShouldEqual, -2)
		So(res1.vals[2], ShouldEqual, 3)
	})

	Convey("Should scale a vector", t, func() {
		v1 := Vector{vals: []float64{7, 3, 9}}
		res1 := v1.Scale(3)
		So(res1.vals[0], ShouldEqual, 21)
		So(res1.vals[1], ShouldEqual, 9)
		So(res1.vals[2], ShouldEqual, 27)
	})

	Convey("Should get the dot product of vectors", t, func() {
		v1 := Vector{vals: []float64{-6, 8}}
		v2 := Vector{vals: []float64{5, 12}}
		res := v1.DotProduct(v2)
		So(res, ShouldEqual, 66)
	})

	Convey("Should rotate a vector", t, func() {
		v1 := Vector{vals: []float64{2, 2}}
		res1 := v1.RotateAboutTail(-math.Pi / 2)
		So(res1.vals[0], ShouldAlmostEqual, -2)
		So(res1.vals[1], ShouldAlmostEqual, 2)
	})
}
