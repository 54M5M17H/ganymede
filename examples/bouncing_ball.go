package main

import (
	"fmt"
	"ganymede/force"
	"ganymede/object"
	"ganymede/vector"
	"image/color"

	"github.com/faiface/pixel/pixelgl"
	"github.com/h8gi/canvas"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

const (
	frameRate = 30

	canvasWidth  = 800
	canvasHeight = 600
)

var (
	canvasColour   = colornames.Amber50
	ballColour     = colornames.LightGreen200
	platformColour = colornames.Purple100
	wallColour     = colornames.Purple100
)

func main() {
	canvasInstance := canvas.NewCanvas(&canvas.CanvasConfig{
		Width:     canvasWidth,
		Height:    canvasHeight,
		FrameRate: 30,
		Title:     "Bouncing Ball",
	})

	background := rectangle{
		canvasColour,
		object.NewRectangleObject(canvasWidth, canvasHeight, 0, vector.NewVector(0, 0)),
	}

	ball := circle{
		ballColour,
		object.NewCircleObject(20, 1, vector.NewVector(400, 400)),
	}

	platform := rectangle{
		platformColour,
		object.NewRectangleObject(800, 100, 0, vector.NewVector(0, 0)),
	}

	walls := []rectangle{
		rectangle{
			wallColour,
			object.NewRectangleObject(50, 400, 0, vector.NewVector(750, 0)),
		},
		rectangle{
			wallColour,
			object.NewRectangleObject(50, 400, 0, vector.NewVector(0, 0)),
		},
	}

	gravity := vector.NewVector(0, -1)
	drag := vector.NewVector(0, 0.2)
	windL := vector.NewVector(10, 0)
	windR := vector.NewVector(-10, 0)

	canvasInstance.Draw(func(ctx *canvas.Context) {

		sumAccel := gravity.Add(drag)

		if ctx.IsKeyPressed(pixelgl.KeyR) {
			fmt.Println("Wind right")
			sumAccel = sumAccel.Add(windL)
		} else if ctx.IsKeyPressed(pixelgl.KeyL) {
			fmt.Println("Wind left")
			sumAccel = sumAccel.Add(windR)
		}

		if collided, collisionNormal := object.DetectCollision(&ball, &platform); collided {
			dimensionV := vector.NewVector(ball.Radius, ball.Radius)
			ball.CollisionOverlapCorrection(collisionNormal, dimensionV)

			springForce := force.Spring(&ball, collisionNormal.AsUnitVector())
			sumAccel = sumAccel.Add(springForce)
		}

		for _, wall := range walls {
			if collided, collisionNormal := object.DetectCollision(&ball, &wall); collided {
				dimensionV := vector.NewVector(ball.Radius, ball.Radius)
				ball.CollisionOverlapCorrection(collisionNormal, dimensionV)

				springForce := force.Spring(&ball, collisionNormal.AsUnitVector())
				sumAccel = sumAccel.Add(springForce)
			}
		}

		ball.ApplyAcceleration(sumAccel)

		ctx.Clear()
		background.Draw(ctx)
		ball.Draw(ctx)
		for _, wall := range walls {
			wall.Draw(ctx)
		}
		platform.Draw(ctx)

	})
}

type circle struct {
	colour color.Color
	object.Circle
}

func (c circle) Draw(ctx *canvas.Context) {
	p := c.GetPosition().GetVals()
	ctx.Push()
	ctx.SetColor(c.colour)
	ctx.DrawCircle(p[0], p[1], c.Radius)
	ctx.Fill()
	ctx.Stroke()
	ctx.Pop()
}

type rectangle struct {
	colour color.Color
	object.Rectangle
}

func (r rectangle) Draw(ctx *canvas.Context) {
	p := r.GetPosition().GetVals()
	d := r.GetDimensions().GetVals()
	ctx.Push()
	ctx.SetColor(r.colour)
	ctx.DrawRoundedRectangle(p[0], p[1], d[0], d[1], 5)
	ctx.Fill()
	ctx.Stroke()
	ctx.Pop()
}
