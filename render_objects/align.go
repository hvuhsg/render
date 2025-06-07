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
	childCanvas := cv.NewCanvas(childSize)
	a.Child.Paint(childCanvas)

	switch a.Align {
	case AlignTopLeft:
		canvas.DrawCanvas(childCanvas, 0, 0)
	case AlignTopCenter:
		canvas.DrawCanvas(childCanvas, (canvas.Size.Width-childSize.Width)/2, 0)
	case AlignTopRight:
		canvas.DrawCanvas(childCanvas, canvas.Size.Width-childSize.Width, 0)
	case AlignLeftCenter:
		canvas.DrawCanvas(childCanvas, 0, (canvas.Size.Height-childSize.Height)/2)
	case AlignRightCenter:
		canvas.DrawCanvas(childCanvas, canvas.Size.Width-childSize.Width, (canvas.Size.Height-childSize.Height)/2)
	case AlignBottomLeft:
		canvas.DrawCanvas(childCanvas, 0, canvas.Size.Height-childSize.Height)
	case AlignBottomCenter:
		canvas.DrawCanvas(childCanvas, (canvas.Size.Width-childSize.Width)/2, canvas.Size.Height-childSize.Height)
	case AlignBottomRight:
		canvas.DrawCanvas(childCanvas, canvas.Size.Width-childSize.Width, canvas.Size.Height-childSize.Height)
	case AlignCenter:
		canvas.DrawCanvas(childCanvas, (canvas.Size.Width-childSize.Width)/2, (canvas.Size.Height-childSize.Height)/2)
	}
}

func (a *Align) Size(parentSize types.Size) types.Size {
	return a.Child.Size(parentSize)
}
