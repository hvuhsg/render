package render_objects

import (
	cv "github.com/hvuhsg/render/canvas"
	types "github.com/hvuhsg/render/types"
)

type Align struct {
	Child RenderObject
	Align AlignType
}

type AlignType string

const (
	AlignTopLeft      AlignType = "topLeft"
	AlignTopCenter    AlignType = "topCenter"
	AlignTopRight     AlignType = "topRight"
	AlignLeftCenter   AlignType = "leftCenter"
	AlignRightCenter  AlignType = "rightCenter"
	AlignBottomLeft   AlignType = "bottomLeft"
	AlignBottomCenter AlignType = "bottomCenter"
	AlignBottomRight  AlignType = "bottomRight"
	AlignCenter       AlignType = "center"
)

func (a *Align) Paint(canvas *cv.Canvas) {
	childSize := a.Child.Size(canvas.Size)

	var x, y int
	switch a.Align {
	case AlignTopLeft:
		x = 0
		y = 0
	case AlignTopCenter:
		x = (canvas.Size.Width - childSize.Width) / 2
		y = 0
	case AlignTopRight:
		x = canvas.Size.Width - childSize.Width
		y = 0
	case AlignLeftCenter:
		x = 0
		y = (canvas.Size.Height - childSize.Height) / 2
	case AlignRightCenter:
		x = canvas.Size.Width - childSize.Width
		y = (canvas.Size.Height - childSize.Height) / 2
	case AlignBottomLeft:
		x = 0
		y = canvas.Size.Height - childSize.Height
	case AlignBottomCenter:
		x = (canvas.Size.Width - childSize.Width) / 2
		y = canvas.Size.Height - childSize.Height
	case AlignBottomRight:
		x = canvas.Size.Width - childSize.Width
		y = canvas.Size.Height - childSize.Height
	case AlignCenter:
		x = (canvas.Size.Width - childSize.Width) / 2
		y = (canvas.Size.Height - childSize.Height) / 2
	}

	childCanvas := canvas.SubCanvas(x, y, childSize, nil)
	a.Child.Paint(childCanvas)
}

func (a *Align) Size(parentSize types.Size) types.Size {
	return a.Child.Size(parentSize)
}
