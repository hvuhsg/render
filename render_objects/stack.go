package render_objects

import (
	cv "github.com/hvuhsg/render/canvas"
	types "github.com/hvuhsg/render/types"
)

type Stack struct {
	Children []RenderObject
}

func (s *Stack) Paint(canvas *cv.Canvas) {
	stackCanvas := cv.NewCanvas(s.Size(canvas.Size))
	for _, child := range s.Children {
		child.Paint(stackCanvas)
	}

	canvas.DrawCanvas(stackCanvas, 0, 0)
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
