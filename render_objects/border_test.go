package render_objects

import (
	"image/color"
	"testing"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

func TestBorder(t *testing.T) {
	// Test in-bounds border
	canvas := cv.NewCanvas(types.Size{Width: 100, Height: 100}, false)
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	// Create a colored box
	box := &ColoredBox{
		Width:  50,
		Height: 30,
		Color:  blue,
	}

	// Test in-bounds border
	border := &Border{
		Child: box,
		Width: 2,
		Color: red,
	}

	// Test Size method
	size := border.Size(canvas.Size)
	if size.Width != 50 || size.Height != 30 {
		t.Errorf("Expected border size 50x30 for in-bounds mode, got %v", size)
	}

	// Test Paint method
	border.Paint(canvas)

	// Verify that the box and border were drawn correctly
	hasBluePixels := false
	hasRedPixels := false
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			if canvas.Img.At(x, y) == blue {
				hasBluePixels = true
			}
			if canvas.Img.At(x, y) == red {
				hasRedPixels = true
			}
		}
	}

	if !hasBluePixels {
		t.Error("Expected blue pixels for the box")
	}
	if !hasRedPixels {
		t.Error("Expected red pixels for the border")
	}
}
