package render_objects

import (
	"image/color"
	"testing"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

func TestRow(t *testing.T) {
	canvas := cv.NewCanvas(types.Size{Width: 300, Height: 100})
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	// Create two colored boxes
	box1 := &ColoredBox{Width: 50, Height: 30, Color: red}
	box2 := &ColoredBox{Width: 50, Height: 30, Color: blue}

	// Create a row with the boxes
	row := &Row{
		Children:  []RenderObject{box1, box2},
		Alignment: types.MainAxisAlignmentStart,
		Sizing:    types.MainAxisSizeMax,
	}

	// Test Size method
	size := row.Size(canvas.Size)
	if size.Width != 300 || size.Height != 30 {
		t.Errorf("Expected row size 300x30, got %v", size)
	}

	// Test Paint method
	row.Paint(canvas)

	// Verify that the boxes were drawn correctly
	// First box should be at the start
	for x := 0; x < 50; x++ {
		for y := 0; y < 30; y++ {
			if canvas.Img.At(x, y) != red {
				t.Errorf("Expected red pixel at (%d,%d)", x, y)
			}
		}
	}

	// Second box should be next to the first
	for x := 50; x < 100; x++ {
		for y := 0; y < 30; y++ {
			if canvas.Img.At(x, y) != blue {
				t.Errorf("Expected blue pixel at (%d,%d)", x, y)
			}
		}
	}
}
