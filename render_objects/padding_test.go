package render_objects

import (
	"image/color"
	"testing"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

func TestPadding(t *testing.T) {
	canvas := cv.NewCanvas(types.Size{Width: 200, Height: 200}, false)
	red := color.RGBA{255, 0, 0, 255}

	// Create a colored box
	box := &ColoredBox{Width: 50, Height: 50, Color: red}

	// Test equal padding
	padding := NewPadding(box, 20)
	size := padding.Size(canvas.Size)
	if size.Width != 90 || size.Height != 90 {
		t.Errorf("Expected padded box size 90x90, got %v", size)
	}

	// Test Paint method
	padding.Paint(canvas)

	// Verify that the box was drawn with padding
	hasRedPixels := false
	for x := 20; x < 70; x++ {
		for y := 20; y < 70; y++ {
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
		t.Error("Expected red pixels in the padded area")
	}

	// Verify that padding area is transparent
	transparent := color.RGBA{0, 0, 0, 0}
	for x := 0; x < 200; x++ {
		for y := 0; y < 200; y++ {
			if (x < 20 || x >= 70 || y < 20 || y >= 70) && canvas.Img.At(x, y) != transparent {
				t.Errorf("Expected transparent pixel at (%d,%d)", x, y)
			}
		}
	}
}

func TestPaddingWithDifferentSides(t *testing.T) {
	canvas := cv.NewCanvas(types.Size{Width: 200, Height: 200}, false)
	red := color.RGBA{255, 0, 0, 255}

	// Create a colored box
	box := &ColoredBox{Width: 50, Height: 50, Color: red}

	// Test different padding for each side
	padding := NewPaddingWithSides(box, 10, 20, 30, 40)
	size := padding.Size(canvas.Size)
	if size.Width != 110 || size.Height != 90 {
		t.Errorf("Expected padded box size 110x90, got %v", size)
	}

	// Test Paint method
	padding.Paint(canvas)

	// Verify that the box was drawn with correct padding
	hasRedPixels := false
	for x := 40; x < 90; x++ {
		for y := 10; y < 60; y++ {
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
		t.Error("Expected red pixels in the padded area")
	}

	// Verify that padding area is transparent
	transparent := color.RGBA{0, 0, 0, 0}
	for x := range 200 {
		for y := range 200 {
			if (x < 40 || x >= 90 || y < 10 || y >= 60) && canvas.Img.At(x, y) != transparent {
				t.Errorf("Expected transparent pixel at (%d,%d)", x, y)
			}
		}
	}
}
