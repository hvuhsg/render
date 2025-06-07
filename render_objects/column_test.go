package render_objects

import (
	"image/color"
	"testing"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

func TestColumn(t *testing.T) {
	canvas := cv.NewCanvas(types.Size{Width: 100, Height: 300})
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	// Create two colored boxes
	box1 := &ColoredBox{Width: 50, Height: 30, Color: red}
	box2 := &ColoredBox{Width: 50, Height: 30, Color: blue}

	// Create a column with the boxes
	column := &Column{
		Children:  []RenderObject{box1, box2},
		Alignment: types.MainAxisAlignmentStart,
		Sizing:    types.MainAxisSizeMax,
	}

	// Test Size method
	size := column.Size(canvas.Size)
	if size.Width != 50 || size.Height != 300 {
		t.Errorf("Expected column size 50x300, got %v", size)
	}

	// Test Paint method
	column.Paint(canvas)

	// Verify that the boxes were drawn correctly
	// First box should be at the top
	for x := 0; x < 50; x++ {
		for y := 0; y < 30; y++ {
			if canvas.Img.At(x, y) != red {
				t.Errorf("Expected red pixel at (%d,%d)", x, y)
			}
		}
	}

	// Second box should be below the first
	for x := 0; x < 50; x++ {
		for y := 30; y < 60; y++ {
			if canvas.Img.At(x, y) != blue {
				t.Errorf("Expected blue pixel at (%d,%d)", x, y)
			}
		}
	}
}
