package render_objects

import (
	cv "github.com/hvuhsg/render/canvas"
	types "github.com/hvuhsg/render/types"
)

type Stack struct {
	Children   []RenderObject
	cachedSize *types.Size
}

func (s *Stack) Paint(canvas *cv.Canvas) {
	for _, child := range s.Children {
		childSize := child.Size(canvas.Size)
		childCanvas := canvas.SubCanvas(0, 0, childSize, nil)
		child.Paint(childCanvas)
	}
}

func (s *Stack) Size(parentSize types.Size) types.Size {
	if s.cachedSize != nil {
		return *s.cachedSize
	}

	maxHeight := 0
	maxWidth := 0

	// Calculate sizes in a single pass
	for _, child := range s.Children {
		size := child.Size(parentSize)
		if size.Height > maxHeight {
			maxHeight = size.Height
		}
		if size.Width > maxWidth {
			maxWidth = size.Width
		}
	}

	size := types.Size{Width: maxWidth, Height: maxHeight}
	s.cachedSize = &size

	return size
}
