package canvas

import (
	"image/color"
	"testing"

	"github.com/hvuhsg/render/types"
)

func TestNewCanvas(t *testing.T) {
	size := types.Size{Width: 100, Height: 100}
	canvas := NewCanvas(size)

	if canvas.Size != size {
		t.Errorf("Expected canvas size %v, got %v", size, canvas.Size)
	}

	if canvas.Img.Bounds().Dx() != size.Width || canvas.Img.Bounds().Dy() != size.Height {
		t.Errorf("Expected image bounds %v, got %v", size, canvas.Img.Bounds())
	}
}

func TestSubCanvas(t *testing.T) {
	parent := NewCanvas(types.Size{Width: 100, Height: 100})
	sub := parent.SubCanvas(10, 10, types.Size{Width: 50, Height: 50})

	if sub.Size.Width != 50 || sub.Size.Height != 50 {
		t.Errorf("Expected sub canvas size 50x50, got %v", sub.Size)
	}
}

func TestSubCanvasPixelOperations(t *testing.T) {
	parent := NewCanvas(types.Size{Width: 100, Height: 100})
	sub := parent.SubCanvas(10, 10, types.Size{Width: 50, Height: 50})
	red := color.RGBA{255, 0, 0, 255}

	// Test setting a pixel in the subcanvas
	sub.set(0, 0, red)
	if parent.Img.At(10, 10) != red {
		t.Error("Expected red pixel at (10,10) in parent canvas")
	}

	// Test getting a pixel from the subcanvas
	if sub.get(0, 0) != red {
		t.Error("Expected red pixel at (0,0) in subcanvas")
	}
}

func TestSubCanvasBounds(t *testing.T) {
	parent := NewCanvas(types.Size{Width: 100, Height: 100})
	sub := parent.SubCanvas(10, 10, types.Size{Width: 50, Height: 50})
	red := color.RGBA{255, 0, 0, 255}

	// Test valid bounds
	sub.set(49, 49, red)
	if parent.Img.At(59, 59) != red {
		t.Error("Expected red pixel at (59,59) in parent canvas")
	}

	// Test out of bounds
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for out of bounds pixel in subcanvas")
		}
	}()
	sub.set(50, 50, red)
}

func TestSubCanvasNested(t *testing.T) {
	parent := NewCanvas(types.Size{Width: 100, Height: 100})
	sub1 := parent.SubCanvas(10, 10, types.Size{Width: 50, Height: 50})
	sub2 := sub1.SubCanvas(5, 5, types.Size{Width: 20, Height: 20})
	red := color.RGBA{255, 0, 0, 255}

	// Test setting a pixel in the nested subcanvas
	sub2.set(0, 0, red)
	if parent.Img.At(15, 15) != red {
		t.Error("Expected red pixel at (15,15) in parent canvas")
	}

	// Test getting a pixel from the nested subcanvas
	if sub2.get(0, 0) != red {
		t.Error("Expected red pixel at (0,0) in nested subcanvas")
	}
}

func TestSetPixel(t *testing.T) {
	canvas := NewCanvas(types.Size{Width: 100, Height: 100})
	red := color.RGBA{255, 0, 0, 255}

	// Test valid pixel
	canvas.set(50, 50, red)
	if canvas.Img.At(50, 50) != red {
		t.Error("Expected red pixel at (50,50)")
	}

	// Test out of bounds
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for out of bounds pixel")
		}
	}()
	canvas.set(200, 200, red)
}

func TestDrawCanvas(t *testing.T) {
	parent := NewCanvas(types.Size{Width: 100, Height: 100})
	child := NewCanvas(types.Size{Width: 50, Height: 50})
	red := color.RGBA{255, 0, 0, 255}

	// Fill child canvas with red
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			child.set(x, y, red)
		}
	}

	// Draw child onto parent
	parent.DrawCanvas(child, 25, 25)

	// Check if child was drawn correctly
	for x := 25; x < 75; x++ {
		for y := 25; y < 75; y++ {
			if parent.Img.At(x, y) != red {
				t.Errorf("Expected red pixel at (%d,%d)", x, y)
			}
		}
	}
}

func TestDrawCanvasOutOfBounds(t *testing.T) {
	parent := NewCanvas(types.Size{Width: 100, Height: 100})
	child := NewCanvas(types.Size{Width: 50, Height: 50})

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for out of bounds canvas drawing")
		}
	}()
	parent.DrawCanvas(child, 60, 60)
}
