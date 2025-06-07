package render_objects

import (
	"image/color"
	"testing"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

func TestPainter(t *testing.T) {
	canvas := cv.NewCanvas(types.Size{Width: 100, Height: 100})
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	// Create a painter that draws a red circle and a blue line
	painter := &Painter{
		Painter: func(canvas *cv.Canvas) {
			// Draw a circle in the center
			canvas.Circle(50, 50, 20, red, true)
			// Draw a line from top-left to bottom-right
			canvas.Line(10, 10, 90, 90, blue, 2)
		},
		Width:  100,
		Height: 100,
	}

	// Test Size method
	size := painter.Size(canvas.Size)
	if size.Width != 100 || size.Height != 100 {
		t.Errorf("Expected painter size 100x100, got %v", size)
	}

	// Test Paint method
	painter.Paint(canvas)

	// Verify that both shapes were drawn
	hasRedPixels := false
	hasBluePixels := false

	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			if canvas.Img.At(x, y) == red {
				hasRedPixels = true
			}
			if canvas.Img.At(x, y) == blue {
				hasBluePixels = true
			}
		}
	}

	if !hasRedPixels {
		t.Error("Expected red pixels to be drawn for the circle")
	}
	if !hasBluePixels {
		t.Error("Expected blue pixels to be drawn for the line")
	}
}

func TestPainterWithCustomSize(t *testing.T) {
	canvas := cv.NewCanvas(types.Size{Width: 200, Height: 200})
	red := color.RGBA{255, 0, 0, 255}

	// Create a painter with custom size
	painter := &Painter{
		Painter: func(canvas *cv.Canvas) {
			// Draw a rectangle that fits within the canvas
			canvas.Rectangle(0, 0, 50, 50, red, true)
		},
		Width:  50,
		Height: 50,
	}

	// Test Size method with different parent sizes
	testSizes := []types.Size{
		{Width: 100, Height: 100},
		{Width: 200, Height: 200},
		{Width: 50, Height: 50},
	}

	for _, parentSize := range testSizes {
		size := painter.Size(parentSize)
		if size.Width != 50 || size.Height != 50 {
			t.Errorf("Expected painter size 50x50 for parent size %v, got %v", parentSize, size)
		}
	}

	// Test Paint method
	painter.Paint(canvas)

	// Verify that the rectangle was drawn
	hasRedPixels := false
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			if canvas.Img.At(x, y) == red {
				hasRedPixels = true
				break
			}
		}
		if hasRedPixels {
			break
		}
	}

	if !hasRedPixels {
		t.Error("Expected red pixels to be drawn for the rectangle")
	}
}
