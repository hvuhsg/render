package main

import (
	"image/png"
	"os"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/render_objects"
	"github.com/hvuhsg/render/types"
)

// CustomPattern is a custom render object that creates a pattern
type CustomPattern struct{}

func (p *CustomPattern) Paint(canvas *cv.Canvas) {
	// Draw a pattern of circles
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			x := i*100 + 100
			y := j*100 + 100
			radius := 40
			color := cv.Red
			if (i+j)%2 == 0 {
				color = cv.Blue
			}
			canvas.Circle(x, y, radius, color, true)
		}
	}
}

func main() {
	// Create a new canvas
	canvas := cv.NewCanvas(types.Size{Width: 800, Height: 600})

	// Create a custom pattern
	pattern := &CustomPattern{}

	// Create a painter that uses our custom pattern
	painter := &render_objects.Painter{
		Painter: pattern.Paint,
		Width:   800,
		Height:  600,
	}

	// Render the pattern
	painter.Paint(canvas)

	// Save the result
	file, _ := os.Create("custom_rendering.png")
	defer file.Close()
	png.Encode(file, canvas.Img)
}
