package canvas

import (
	"image/color"
	"testing"

	"github.com/hvuhsg/render/types"
)

func TestRectangle(t *testing.T) {
	canvas := NewCanvas(types.Size{Width: 100, Height: 100})
	red := color.RGBA{255, 0, 0, 255}

	// Test filled rectangle
	canvas.Rectangle(10, 10, 30, 30, red, true)

	// Check if rectangle was drawn correctly
	for x := 10; x < 40; x++ {
		for y := 10; y < 40; y++ {
			if canvas.Img.At(x, y) != red {
				t.Errorf("Expected red pixel at (%d,%d)", x, y)
			}
		}
	}

	// Test unfilled rectangle
	blue := color.RGBA{0, 0, 255, 255}
	canvas.Rectangle(50, 50, 30, 30, blue, false)

	// Check if only the outline was drawn
	for x := 50; x < 80; x++ {
		for y := 50; y < 80; y++ {
			if x == 50 || x == 79 || y == 50 || y == 79 {
				if canvas.Img.At(x, y) != blue {
					t.Errorf("Expected blue pixel at outline (%d,%d)", x, y)
				}
			} else {
				if canvas.Img.At(x, y) == blue {
					t.Errorf("Unexpected blue pixel inside rectangle at (%d,%d)", x, y)
				}
			}
		}
	}
}

func TestCircle(t *testing.T) {
	canvas := NewCanvas(types.Size{Width: 100, Height: 100})
	green := color.RGBA{0, 255, 0, 255}

	// Test filled circle
	canvas.Circle(50, 50, 20, green, true)

	// Check if circle was drawn correctly
	// We'll check a few key points that should be inside the circle
	points := []struct{ x, y int }{
		{50, 50}, // center
		{50, 30}, // top
		{50, 70}, // bottom
		{30, 50}, // left
		{70, 50}, // right
	}

	for _, p := range points {
		if canvas.Img.At(p.x, p.y) != green {
			t.Errorf("Expected green pixel at (%d,%d)", p.x, p.y)
		}
	}

	// Test unfilled circle
	yellow := color.RGBA{255, 255, 0, 255}
	canvas.Circle(50, 50, 20, yellow, false)

	// Check if only the outline was drawn
	// We'll check points that should be on the outline
	outlinePoints := []struct{ x, y int }{
		{50, 30}, // top
		{50, 70}, // bottom
		{30, 50}, // left
		{70, 50}, // right
	}

	for _, p := range outlinePoints {
		if canvas.Img.At(p.x, p.y) != yellow {
			t.Errorf("Expected yellow pixel at outline (%d,%d)", p.x, p.y)
		}
	}
}

func TestLine(t *testing.T) {
	canvas := NewCanvas(types.Size{Width: 100, Height: 100})
	purple := color.RGBA{128, 0, 128, 255}

	// Test horizontal line
	canvas.Line(10, 50, 90, 50, purple, 2)

	// Check if line was drawn correctly
	for x := 10; x < 90; x++ {
		if canvas.Img.At(x, 49) != purple && canvas.Img.At(x, 50) != purple && canvas.Img.At(x, 51) != purple {
			t.Errorf("Expected purple pixel at y=50 near x=%d", x)
		}
	}

	// Test vertical line
	orange := color.RGBA{255, 165, 0, 255}
	canvas.Line(50, 10, 50, 90, orange, 2)

	// Check if line was drawn correctly
	for y := 10; y < 90; y++ {
		if canvas.Img.At(49, y) != orange && canvas.Img.At(50, y) != orange && canvas.Img.At(51, y) != orange {
			t.Errorf("Expected orange pixel at x=50 near y=%d", y)
		}
	}
}

func TestPolygon(t *testing.T) {
	canvas := NewCanvas(types.Size{Width: 100, Height: 100})
	blue := color.RGBA{0, 0, 255, 255}

	// Test triangle
	points := [][2]int{
		{50, 10}, // top
		{30, 50}, // bottom left
		{70, 50}, // bottom right
	}

	canvas.Polygon(points, blue, true)

	// Check if polygon was drawn correctly
	// We'll check a few key points that should be inside the triangle
	insidePoints := []struct{ x, y int }{
		{50, 30}, // center
		{40, 40}, // left side
		{60, 40}, // right side
	}

	for _, p := range insidePoints {
		if canvas.Img.At(p.x, p.y) != blue {
			t.Errorf("Expected blue pixel at (%d,%d)", p.x, p.y)
		}
	}
}
