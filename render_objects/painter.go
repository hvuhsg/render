package render_objects

import (
	cv "github.com/hvuhsg/render/canvas"
	types "github.com/hvuhsg/render/types"
)

type Painter struct {
	Painter func(canvas *cv.Canvas)
	Width   int
	Height  int
}

func (p *Painter) Paint(canvas *cv.Canvas) {
	p.Painter(canvas)
}

func (p *Painter) Size(parentSize types.Size) types.Size {
	return types.Size{Width: p.Width, Height: p.Height}
}
