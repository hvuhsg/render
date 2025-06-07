package canvas

import "image/color"

func (c *Canvas) Rectangle(x, y, w, h int, color color.RGBA, fill bool) {
	// No bounds checking here; set will handle it

	if fill {
		// For filled rectangles, we can use a simpler approach
		for i := x; i < x+w; i++ {
			for j := y; j < y+h; j++ {
				c.set(i, j, color)
			}
		}
	} else {
		// For unfilled rectangles, we only draw the outline
		// Draw top and bottom lines
		for i := x; i < x+w; i++ {
			c.set(i, y, color)
			c.set(i, y+h-1, color)
		}
		// Draw left and right lines
		for j := y + 1; j < y+h-1; j++ {
			c.set(x, j, color)
			c.set(x+w-1, j, color)
		}
	}
}
