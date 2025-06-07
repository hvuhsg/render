package render_objects

import (
	"image/color"

	cv "github.com/hvuhsg/render/canvas"
	types "github.com/hvuhsg/render/types"
)

type ColoredBox struct {
	Width  int
	Height int
	Color  color.RGBA
}

func (c *ColoredBox) Paint(canvas *cv.Canvas) {
	canvas.Rectangle(0, 0, c.Width, c.Height, c.Color, true)
}

func (c *ColoredBox) Size(parentSize types.Size) types.Size {
	return types.Size{
		Width:  c.Width,
		Height: c.Height,
	}
}
