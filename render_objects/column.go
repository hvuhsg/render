package render_objects

import (
	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

type Column struct {
	Children []RenderObject
}

func (c *Column) Paint(canvas *cv.Canvas) {
	totalHeight := 0

	for _, child := range c.Children {
		// Draw the child onto a new canvas
		size := child.Size(canvas.Size)
		childCanvas := cv.NewCanvas(size)
		child.Paint(childCanvas)

		// Draw the child onto the parent canvas
		canvas.DrawCanvas(childCanvas, 0, totalHeight)

		totalHeight += size.Height
	}
}

func (c *Column) Size(parentSize types.Size) types.Size {
	totalHeight := 0
	for _, child := range c.Children {
		totalHeight += child.Size(parentSize).Height
	}

	maxWidth := 0
	for _, child := range c.Children {
		if child.Size(parentSize).Width > maxWidth {
			maxWidth = child.Size(parentSize).Width
		}
	}

	return types.Size{
		Width:  maxWidth,
		Height: totalHeight,
	}
}
