package canvas

import (
	"image/color"
	"math"
)

func (c *Canvas) Line(x1, y1, x2, y2 int, color color.RGBA, width int) {
	// No bounds checking here; set will handle it

	// Handle single point case
	if x1 == x2 && y1 == y2 {
		c.Circle(x1, y1, width, color, true)
		return
	}

	// Use optimized line drawing based on width
	if width <= 1 {
		c.drawLine(x1, y1, x2, y2, color)
	} else {
		c.drawThickLine(x1, y1, x2, y2, color, width)
	}
}

// drawLine draws a line using Bresenham's algorithm
func (c *Canvas) drawLine(x1, y1, x2, y2 int, color color.RGBA) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx := 1
	if x1 > x2 {
		sx = -1
	}
	sy := 1
	if y1 > y2 {
		sy = -1
	}
	err := dx - dy

	for {
		c.set(x1, y1, color)
		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

// drawThickLine draws a line with width by drawing a filled polygon between parallel lines
func (c *Canvas) drawThickLine(x1, y1, x2, y2 int, color color.RGBA, width int) {
	if width <= 1 {
		c.drawLine(x1, y1, x2, y2, color)
		return
	}

	// Calculate the perpendicular vector
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	length := math.Sqrt(dx*dx + dy*dy)
	if length == 0 {
		c.Circle(x1, y1, width/2, color, true)
		return
	}

	// Normalize and scale by half width
	halfWidth := float64(width) / 2
	px := -dy / length * halfWidth
	py := dx / length * halfWidth

	// Calculate the four corners of the thick line
	corners := [][2]int{
		{x1 + int(math.Round(px)), y1 + int(math.Round(py))},
		{x1 - int(math.Round(px)), y1 - int(math.Round(py))},
		{x2 - int(math.Round(px)), y2 - int(math.Round(py))},
		{x2 + int(math.Round(px)), y2 + int(math.Round(py))},
	}

	// Draw the filled polygon
	c.Polygon(corners, color, true)

	// Add extra coverage for the end caps
	c.Circle(x1, y1, width/2, color, true)
	c.Circle(x2, y2, width/2, color, true)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
