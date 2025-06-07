package render_objects

import (
	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

type Column struct {
	Alignment      types.MainAxisAlignment
	Sizing         types.MainAxisSize
	Children       []RenderObject
	cachedSize     *types.Size
	lastParentSize types.Size
}

func (c *Column) Paint(canvas *cv.Canvas) {
	// Calculate total height and get individual child sizes
	totalHeight := 0
	childSizes := make([]types.Size, len(c.Children))
	for i, child := range c.Children {
		size := child.Size(canvas.Size)
		childSizes[i] = size
		totalHeight += size.Height
	}

	// Calculate spacing and offsets based on alignment
	var yOffsets []int
	availableSpace := canvas.Size.Height - totalHeight

	switch c.Alignment {
	case types.MainAxisAlignmentStart:
		// Children start from the top (default behavior)
		yOffsets = make([]int, len(c.Children))
		offset := 0
		for i := range c.Children {
			yOffsets[i] = offset
			offset += childSizes[i].Height
		}

	case types.MainAxisAlignmentEnd:
		// Children start from the bottom
		yOffsets = make([]int, len(c.Children))
		offset := availableSpace
		for i := range c.Children {
			yOffsets[i] = offset
			offset += childSizes[i].Height
		}

	case types.MainAxisAlignmentCenter:
		// Children are centered
		yOffsets = make([]int, len(c.Children))
		offset := availableSpace / 2
		for i := range c.Children {
			yOffsets[i] = offset
			offset += childSizes[i].Height
		}

	case types.MainAxisAlignmentSpaceBetween:
		// Space is distributed between children
		if len(c.Children) <= 1 {
			yOffsets = make([]int, len(c.Children))
			if len(c.Children) == 1 {
				yOffsets[0] = availableSpace / 2
			}
		} else {
			spaceBetween := availableSpace / (len(c.Children) - 1)
			yOffsets = make([]int, len(c.Children))
			offset := 0
			for i := range c.Children {
				yOffsets[i] = offset
				offset += childSizes[i].Height + spaceBetween
			}
		}

	case types.MainAxisAlignmentSpaceAround:
		// Space is distributed around children
		spaceAround := availableSpace / len(c.Children)
		yOffsets = make([]int, len(c.Children))
		offset := spaceAround / 2
		for i := range c.Children {
			yOffsets[i] = offset
			offset += childSizes[i].Height + spaceAround
		}

	case types.MainAxisAlignmentSpaceEvenly:
		// Space is distributed evenly
		spaceEvenly := availableSpace / (len(c.Children) + 1)
		yOffsets = make([]int, len(c.Children))
		offset := spaceEvenly
		for i := range c.Children {
			yOffsets[i] = offset
			offset += childSizes[i].Height + spaceEvenly
		}
	}

	// Draw children at calculated positions
	for i, child := range c.Children {
		childCanvas := canvas.SubCanvas(0, yOffsets[i], childSizes[i], nil)
		child.Paint(childCanvas)
	}
}

func (c *Column) Size(parentSize types.Size) types.Size {
	// Check if we can use cached size
	if c.cachedSize != nil && c.lastParentSize == parentSize {
		return *c.cachedSize
	}

	totalHeight := 0
	maxWidth := 0

	// Calculate sizes in a single pass without storing all sizes
	for _, child := range c.Children {
		size := child.Size(parentSize)
		totalHeight += size.Height
		if size.Width > maxWidth {
			maxWidth = size.Width
		}
	}

	// If MainAxisSizeMax is set, use the parent height
	// Otherwise, use the minimum height needed for children
	height := totalHeight
	if c.Sizing == types.MainAxisSizeMax {
		height = parentSize.Height
	}

	// Cache the result
	size := types.Size{
		Width:  maxWidth,
		Height: height,
	}
	c.cachedSize = &size
	c.lastParentSize = parentSize

	return size
}
