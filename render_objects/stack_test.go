package render_objects

import (
	"image/color"
	"testing"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

func TestStack(t *testing.T) {
	canvas := cv.NewCanvas(types.Size{Width: 100, Height: 100}, false)
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	// Create two colored boxes
	box1 := &ColoredBox{Width: 100, Height: 100, Color: red}
	box2 := &ColoredBox{Width: 50, Height: 50, Color: blue}

	// Create a stack with the boxes
	stack := &Stack{
		Children: []RenderObject{box1, box2},
	}

	// Test Size method
	size := stack.Size(canvas.Size)
	if size.Width != 100 || size.Height != 100 {
		t.Errorf("Expected stack size 100x100, got %v", size)
	}

	// Test Paint method
	stack.Paint(canvas)

	// Verify that both boxes were drawn
	hasRedPixels := false
	hasBluePixels := false
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			c := canvas.Img.At(x, y)
			if c == red {
				hasRedPixels = true
			}
			if c == blue {
				hasBluePixels = true
			}
		}
	}
	if !hasRedPixels {
		t.Error("Expected red pixels to be drawn for the bottom box")
	}
	if !hasBluePixels {
		t.Error("Expected blue pixels to be drawn for the top box")
	}
}
