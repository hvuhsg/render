package render_objects

import (
	"image/color"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

// Border represents a border around a render object
type Border struct {
	Child RenderObject
	Width int
	Color color.RGBA
}

// Paint implements the RenderObject interface
func (b *Border) Paint(canvas *cv.Canvas) {
	// Get the canvas size
	width, height := canvas.Size.Width, canvas.Size.Height

	// First paint the child
	b.Child.Paint(canvas)

	// Draw border inside the content bounds
	// Top border
	canvas.Rectangle(0, 0, width, b.Width, b.Color, true)
	// Bottom border
	canvas.Rectangle(0, height-b.Width, width, b.Width, b.Color, true)
	// Left border
	canvas.Rectangle(0, b.Width, b.Width, height-2*b.Width, b.Color, true)
	// Right border
	canvas.Rectangle(width-b.Width, b.Width, b.Width, height-2*b.Width, b.Color, true)

}

// Size implements the RenderObject interface
func (b *Border) Size(parentSize types.Size) types.Size {
	return b.Child.Size(parentSize)
}
