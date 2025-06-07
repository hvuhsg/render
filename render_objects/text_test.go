package render_objects

import (
	"image/color"
	"testing"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

func TestText(t *testing.T) {
	canvas := cv.NewCanvas(types.Size{Width: 200, Height: 100}, false)
	red := color.RGBA{255, 0, 0, 255}

	text := NewText("Hello", red, 24, "default")

	// Test Size method
	size := text.Size(canvas.Size)
	if size.Width <= 0 || size.Height <= 0 {
		t.Errorf("Expected positive text dimensions, got %v", size)
	}

	// Test Paint method
	text.Paint(canvas)

	// Verify that some pixels were drawn
	hasRedPixels := false
	for x := 0; x < 200; x++ {
		for y := 0; y < 100; y++ {
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
		t.Error("Expected some red pixels to be drawn for the text")
	}
}
