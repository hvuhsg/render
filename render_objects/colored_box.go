package render_objects

import (
	"image/color"

	"github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

type ColoredBox struct {
	Color  color.RGBA
	Width  int
	Height int
}

func (cb *ColoredBox) Paint(c *canvas.Canvas) {
	c.Rectangle(0, 0, cb.Width, cb.Height, cb.Color, true)
}

func (cb *ColoredBox) Size(parentSize types.Size) types.Size {
	return types.Size{
		Width:  cb.Width,
		Height: cb.Height,
	}
}
