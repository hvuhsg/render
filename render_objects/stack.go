package render_objects

import (
	cv "github.com/hvuhsg/render/canvas"
	types "github.com/hvuhsg/render/types"
)

type Stack struct {
	Children []RenderObject
}

func (s *Stack) Paint(canvas *cv.Canvas) {
	for _, child := range s.Children {
		childSize := child.Size(canvas.Size)
		childCanvas := canvas.SubCanvas(0, 0, childSize)
		child.Paint(childCanvas)
	}
}

func (s *Stack) Size(parentSize types.Size) types.Size {
	maxHeight := 0
	for _, child := range s.Children {
		if child.Size(parentSize).Height > maxHeight {
			maxHeight = child.Size(parentSize).Height
		}
	}

	maxWidth := 0
	for _, child := range s.Children {
		if child.Size(parentSize).Width > maxWidth {
			maxWidth = child.Size(parentSize).Width
		}
	}

	return types.Size{Width: maxWidth, Height: maxHeight}
}
