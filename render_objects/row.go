package render_objects

import (
	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

type RowAlignment int
type RowSizing int

const (
	RowAlignmentStart RowAlignment = iota
	RowAlignmentCenter
	RowAlignmentEnd
	RowAlignmentSpaceBetween
	RowAlignmentSpaceAround
	RowAlignmentSpaceEvenly
)

const (
	RowSizingMin RowSizing = iota // Row takes minimum width needed for children
	RowSizingMax                  // Row takes maximum width (canvas width)
)

type Row struct {
	Alignment RowAlignment
	Sizing    RowSizing
	Children  []RenderObject
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
	case RowAlignmentStart:
		// Children start from the left (default behavior)
		xOffsets = make([]int, len(r.Children))
		offset := 0
		for i := range r.Children {
			xOffsets[i] = offset
			offset += childSizes[i].Width
		}

	case RowAlignmentEnd:
		// Children start from the right
		xOffsets = make([]int, len(r.Children))
		offset := availableSpace
		for i := range r.Children {
			xOffsets[i] = offset
			offset += childSizes[i].Width
		}

	case RowAlignmentCenter:
		// Children are centered
		xOffsets = make([]int, len(r.Children))
		offset := availableSpace / 2
		for i := range r.Children {
			xOffsets[i] = offset
			offset += childSizes[i].Width
		}

	case RowAlignmentSpaceBetween:
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

	case RowAlignmentSpaceAround:
		// Space is distributed around children
		spaceAround := availableSpace / len(r.Children)
		xOffsets = make([]int, len(r.Children))
		offset := spaceAround / 2
		for i := range r.Children {
			xOffsets[i] = offset
			offset += childSizes[i].Width + spaceAround
		}

	case RowAlignmentSpaceEvenly:
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
		childCanvas := cv.NewCanvas(childSizes[i])
		child.Paint(childCanvas)
		canvas.DrawCanvas(childCanvas, xOffsets[i], 0)
	}
}

func (r *Row) Size(parentSize types.Size) types.Size {
	totalWidth := 0
	for _, child := range r.Children {
		totalWidth += child.Size(parentSize).Width
	}

	maxHeight := 0
	for _, child := range r.Children {
		if child.Size(parentSize).Height > maxHeight {
			maxHeight = child.Size(parentSize).Height
		}
	}

	// If RowSizingMax is set, use the parent width
	// Otherwise, use the minimum width needed for children
	width := totalWidth
	if r.Sizing == RowSizingMax {
		width = parentSize.Width
	}

	return types.Size{
		Width:  width,
		Height: maxHeight,
	}
}
