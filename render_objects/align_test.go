package render_objects

import (
	"image/color"
	"testing"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

func TestAlign(t *testing.T) {
	canvas := cv.NewCanvas(types.Size{Width: 200, Height: 200}, false)
	red := color.RGBA{255, 0, 0, 255}

	// Create a colored box
	box := &ColoredBox{Width: 50, Height: 50, Color: red}

	// Test different alignments
	alignments := []AlignType{
		AlignTopLeft,
		AlignTopCenter,
		AlignTopRight,
		AlignLeftCenter,
		AlignCenter,
		AlignRightCenter,
		AlignBottomLeft,
		AlignBottomCenter,
		AlignBottomRight,
	}

	for _, alignment := range alignments {
		align := &Align{
			Child: box,
			Align: alignment,
		}

		// Test Size method
		size := align.Size(canvas.Size)
		if size.Width != 50 || size.Height != 50 {
			t.Errorf("Expected aligned box size 50x50, got %v", size)
		}

		// Test Paint method
		align.Paint(canvas)

		// Verify that the box was drawn
		hasRedPixels := false
		for x := 0; x < 200; x++ {
			for y := 0; y < 200; y++ {
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
			t.Errorf("Expected red pixels for alignment %v", alignment)
		}
	}
}
