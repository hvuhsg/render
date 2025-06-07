package canvas

import (
	"image/color"
	"testing"

	"github.com/hvuhsg/render/types"
)

func TestNewTextPainter(t *testing.T) {
	painter := NewTextPainter()

	if painter.Font == nil {
		t.Error("Expected non-nil font")
	}

	if painter.FontSize != 12 {
		t.Errorf("Expected default font size 12, got %f", painter.FontSize)
	}

	if painter.TextColor != color.Black {
		t.Error("Expected default text color to be black")
	}
}

func TestDrawText(t *testing.T) {
	canvas := NewCanvas(types.Size{Width: 200, Height: 100}, false)
	red := color.RGBA{255, 0, 0, 255}

	painter := NewTextPainter()
	painter.TextColor = red
	painter.FontSize = 24

	// Draw text
	canvas.DrawText("Hello", 10, 30, painter)

	// Since text rendering is complex and depends on the font,
	// we'll just verify that some pixels were drawn
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

func TestMeasureText(t *testing.T) {
	canvas := NewCanvas(types.Size{Width: 200, Height: 100}, false)
	painter := NewTextPainter()
	painter.FontSize = 24

	// Test measurement of a simple string
	size := canvas.MeasureText("Hello", painter)

	if size.Width <= 0 || size.Height <= 0 {
		t.Errorf("Expected positive text dimensions, got %v", size)
	}

	// Test measurement of a longer string
	longerSize := canvas.MeasureText("Hello, World!", painter)

	if longerSize.Width <= size.Width {
		t.Error("Expected longer text to have greater width")
	}

	// Test measurement with different font sizes
	painter.FontSize = 36
	largerSize := canvas.MeasureText("Hello", painter)

	if largerSize.Width <= size.Width || largerSize.Height <= size.Height {
		t.Error("Expected larger font size to result in larger dimensions")
	}
}

func TestDrawTextOutOfBounds(t *testing.T) {
	canvas := NewCanvas(types.Size{Width: 100, Height: 100}, false)
	painter := NewTextPainter()

	// Test drawing text that would go out of bounds
	// This should not panic as the canvas should clip the text
	canvas.DrawText("This is a very long text that should be clipped", 90, 50, painter)

	// Verify that no pixels were drawn outside the canvas bounds
	transparent := color.RGBA{0, 0, 0, 0}
	for x := 100; x < 200; x++ {
		for y := 0; y < 100; y++ {
			if canvas.Img.At(x, y) != transparent {
				t.Errorf("Expected transparent pixel at (%d,%d)", x, y)
			}
		}
	}
}
