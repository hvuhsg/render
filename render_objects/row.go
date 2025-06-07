package render_objects

import (
	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

type Row struct {
	Alignment      types.MainAxisAlignment
	Sizing         types.MainAxisSize
	Children       []RenderObject
	cachedSize     *types.Size
	lastParentSize types.Size
}

func (r *Row) Paint(canvas *cv.Canvas) {
	// Calculate total width and get individual child sizes
	totalWidth := 0
	childSizes := make([]types.Size, len(r.Children))
	for i, child := range r.Children {
		size := child.Size(canvas.Size)
		childSizes[i] = size
		totalWidth += size.Width
	}

	// Calculate spacing and offsets based on alignment
	var xOffsets []int
	availableSpace := canvas.Size.Width - totalWidth

	switch r.Alignment {
	case types.MainAxisAlignmentStart:
		// Children start from the left (default behavior)
		xOffsets = make([]int, len(r.Children))
		offset := 0
		for i := range r.Children {
			xOffsets[i] = offset
			offset += childSizes[i].Width
		}

	case types.MainAxisAlignmentEnd:
		// Children start from the right
		xOffsets = make([]int, len(r.Children))
		offset := availableSpace
		for i := range r.Children {
			xOffsets[i] = offset
			offset += childSizes[i].Width
		}

	case types.MainAxisAlignmentCenter:
		// Children are centered
		xOffsets = make([]int, len(r.Children))
		offset := availableSpace / 2
		for i := range r.Children {
			xOffsets[i] = offset
			offset += childSizes[i].Width
		}

	case types.MainAxisAlignmentSpaceBetween:
		// Space is distributed between children
		if len(r.Children) <= 1 {
			xOffsets = make([]int, len(r.Children))
			if len(r.Children) == 1 {
				xOffsets[0] = availableSpace / 2
			}
		} else {
			spaceBetween := availableSpace / (len(r.Children) - 1)
			xOffsets = make([]int, len(r.Children))
			offset := 0
			for i := range r.Children {
				xOffsets[i] = offset
				offset += childSizes[i].Width + spaceBetween
			}
		}

	case types.MainAxisAlignmentSpaceAround:
		// Space is distributed around children
		spaceAround := availableSpace / len(r.Children)
		xOffsets = make([]int, len(r.Children))
		offset := spaceAround / 2
		for i := range r.Children {
			xOffsets[i] = offset
			offset += childSizes[i].Width + spaceAround
		}

	case types.MainAxisAlignmentSpaceEvenly:
		// Space is distributed evenly
		spaceEvenly := availableSpace / (len(r.Children) + 1)
		xOffsets = make([]int, len(r.Children))
		offset := spaceEvenly
		for i := range r.Children {
			xOffsets[i] = offset
			offset += childSizes[i].Width + spaceEvenly
		}
	}

	// Draw children at calculated positions
	for i, child := range r.Children {
		childCanvas := canvas.SubCanvas(xOffsets[i], 0, childSizes[i], nil)
		child.Paint(childCanvas)
	}
}

func (r *Row) Size(parentSize types.Size) types.Size {
	// Check if we can use cached size
	if r.cachedSize != nil && r.lastParentSize == parentSize {
		return *r.cachedSize
	}

	totalWidth := 0
	maxHeight := 0

	// Calculate sizes in a single pass without storing all sizes
	for _, child := range r.Children {
		size := child.Size(parentSize)
		totalWidth += size.Width
		if size.Height > maxHeight {
			maxHeight = size.Height
		}
	}

	// If MainAxisSizeMax is set, use the parent width
	// Otherwise, use the minimum width needed for children
	width := totalWidth
	if r.Sizing == types.MainAxisSizeMax {
		width = parentSize.Width
	}

	// Cache the result
	size := types.Size{
		Width:  width,
		Height: maxHeight,
	}
	r.cachedSize = &size
	r.lastParentSize = parentSize

	return size
}
