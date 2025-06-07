package render_objects

import (
	cv "github.com/hvuhsg/render/canvas"
	types "github.com/hvuhsg/render/types"
)

// Padding wraps a render object and adds padding around it
type Padding struct {
	Child  RenderObject
	Top    int
	Right  int
	Bottom int
	Left   int
}

// NewPadding creates a new Padding render object with equal padding on all sides
func NewPadding(child RenderObject, padding int) *Padding {
	return &Padding{
		Child:  child,
		Top:    padding,
		Right:  padding,
		Bottom: padding,
		Left:   padding,
	}
}

// NewPaddingWithSides creates a new Padding render object with different padding values for each side
func NewPaddingWithSides(child RenderObject, top, right, bottom, left int) *Padding {
	return &Padding{
		Child:  child,
		Top:    top,
		Right:  right,
		Bottom: bottom,
		Left:   left,
	}
}

func (p *Padding) Paint(canvas *cv.Canvas) {
	// Create a subcanvas for the child with adjusted size
	childSize := types.Size{
		Width:  canvas.Size.Width - p.Left - p.Right,
		Height: canvas.Size.Height - p.Top - p.Bottom,
	}
	childCanvas := canvas.SubCanvas(p.Left, p.Top, childSize, nil)

	// Paint the child on the subcanvas
	p.Child.Paint(childCanvas)
}

func (p *Padding) Size(parentSize types.Size) types.Size {
	// Calculate the available size for the child
	childSize := types.Size{
		Width:  parentSize.Width - p.Left - p.Right,
		Height: parentSize.Height - p.Top - p.Bottom,
	}

	// Get the child's size
	childActualSize := p.Child.Size(childSize)

	// Return the total size including padding
	return types.Size{
		Width:  childActualSize.Width + p.Left + p.Right,
		Height: childActualSize.Height + p.Top + p.Bottom,
	}
}
