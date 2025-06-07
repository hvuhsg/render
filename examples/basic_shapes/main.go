package main

import (
	"image/png"
	"os"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

func main() {
	// Create a new canvas
	canvas := cv.NewCanvas(types.Size{Width: 800, Height: 600}, false)

	// Draw various shapes
	canvas.Circle(400, 300, 100, cv.Red, true)           // Filled red circle
	canvas.Circle(400, 300, 120, cv.Blue, false)         // Blue circle outline
	canvas.Rectangle(200, 200, 100, 100, cv.Green, true) // Filled green square
	canvas.Line(3, 3, 796, 596, cv.Yellow, 5)            // Yellow diagonal line (inset to avoid out-of-bounds)

	// Save the result
	file, _ := os.Create("basic_shapes.png")
	defer file.Close()
	png.Encode(file, canvas.Img)
}
