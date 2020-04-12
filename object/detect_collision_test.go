package object

import (
	"ganymede/vector"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCircleCollisions(t *testing.T) {
	Convey("Should detect circle collision", t, func() {
		c1 := NewCircleObject(10, 1, vector.NewVector(100, 100))
		c2 := NewCircleObject(10, 1, vector.NewVector(100, 119))
		res, _ := DetectCollision(&c1, &c2)
		So(res, ShouldBeTrue)
	})

	Convey("Circles not colliding", t, func() {
		c1 := NewCircleObject(10, 1, vector.NewVector(100, 100))
		c2 := NewCircleObject(10, 1, vector.NewVector(100, 121))
		res, _ := DetectCollision(&c1, &c2)
		So(res, ShouldBeFalse)
	})
}

func TestBBCollisions(t *testing.T) {
	Convey("Should detect boxes colliding", t, func() {
		r1 := NewRectangleObject(10, 20, 1, vector.NewVector(10, 10))
		r2 := NewRectangleObject(15, 20, 1, vector.NewVector(15, 20))
		res, _ := DetectCollision(&r1, &r2)
		So(res, ShouldBeTrue)
	})

	Convey("Boxes not colliding", t, func() {
		r1 := NewRectangleObject(10, 20, 1, vector.NewVector(10, 10))
		r2 := NewRectangleObject(15, 20, 1, vector.NewVector(10, 31))
		res, _ := DetectCollision(&r1, &r2)
		So(res, ShouldBeFalse)
	})
}

func TestCircleBBCollisions(t *testing.T) {
	Convey("Should calc that point is in box", t, func() {
		b := NewRectangleObject(10, 20, 1, vector.NewVector(10, 10))
		v := vector.NewVector(10, 20)
		res := isPointInsideBox(v, &b)
		So(res, ShouldBeTrue)

		v = vector.NewVector(20, 20)
		res = isPointInsideBox(v, &b)
		So(res, ShouldBeTrue)
	})

	Convey("Should calc that point is NOT in box", t, func() {
		b := NewRectangleObject(10, 20, 1, vector.NewVector(10, 10))
		v := vector.NewVector(9, 9)
		res := isPointInsideBox(v, &b)
		So(res, ShouldBeFalse)

		v = vector.NewVector(31, 9)
		res = isPointInsideBox(v, &b)
		So(res, ShouldBeFalse)
	})

	Convey("Should find nearest bbox point", t, func() {
		b := NewRectangleObject(10, 10, 1, vector.NewVector(10, 10))
		v := vector.NewVector(15, 30)
		res := nearestBoundingBoxEdge(v, &b)
		So(res.GetVals()[0], ShouldEqual, 15)
		So(res.GetVals()[1], ShouldEqual, 20)

		v = vector.NewVector(25, 25)
		res = nearestBoundingBoxEdge(v, &b)
		So(res.GetVals()[0], ShouldEqual, 20)
		So(res.GetVals()[1], ShouldEqual, 20)

		v = vector.NewVector(25, 15)
		res = nearestBoundingBoxEdge(v, &b)
		So(res.GetVals()[0], ShouldEqual, 20)
		So(res.GetVals()[1], ShouldEqual, 15)
	})

	Convey("Should detect circle is inside bounding box", t, func() {
		c := NewCircleObject(5, 1, vector.NewVector(5, 5))
		b1 := NewRectangleObject(10, 10, 1, vector.NewVector(0, 8))
		b2 := NewRectangleObject(10, 10, 1, vector.NewVector(8, 2))

		res, _ := DetectCollision(&c, &b1)
		So(res, ShouldBeTrue)
		res, _ = DetectCollision(&c, &b2)
		So(res, ShouldBeTrue)
	})

	Convey("Should detect circle is NOT inside bounding box", t, func() {
		c := NewCircleObject(5, 1, vector.NewVector(5, 5))
		b1 := NewRectangleObject(10, 10, 1, vector.NewVector(10, 10))
		b2 := NewRectangleObject(10, 10, 1, vector.NewVector(12, 10))

		res, _ := DetectCollision(&c, &b1)
		So(res, ShouldBeFalse)
		res, _ = DetectCollision(&c, &b2)
		So(res, ShouldBeFalse)
	})
}
