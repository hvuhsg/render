package render_objects

import (
	"image/color"
	"testing"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

func TestColoredBox(t *testing.T) {
	canvas := cv.NewCanvas(types.Size{Width: 200, Height: 100})
	blue := color.RGBA{0, 0, 255, 255}

	box := &ColoredBox{
		Width:  50,
		Height: 30,
		Color:  blue,
	}

	// Test Size method
	size := box.Size(canvas.Size)
	if size.Width != 50 || size.Height != 30 {
		t.Errorf("Expected box size 50x30, got %v", size)
	}

	// Test Paint method
	box.Paint(canvas)

	// Verify that the box was drawn correctly
	for x := 0; x < 50; x++ {
		for y := 0; y < 30; y++ {
			if canvas.Img.At(x, y) != blue {
				t.Errorf("Expected blue pixel at (%d,%d)", x, y)
			}
		}
	}
}
