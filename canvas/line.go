package canvas

import (
	"image/color"
	"math"
)

func (c *Canvas) Line(x1, y1, x2, y2 int, color color.RGBA, width int) {
	c.assertPointInBounds(x1, y1)
	c.assertPointInBounds(x2, y2)

	// Handle single point case
	if x1 == x2 && y1 == y2 {
		c.Circle(x1, y1, width, color, true)
		return
	}

	// Calculate the perpendicular vector for width
	dx := float64(x2 - x1)
	dy := float64(y2 - y1)
	length := math.Sqrt(dx*dx + dy*dy)

	// Normalize and scale by half width
	halfWidth := float64(width) / 2
	px := -dy / length * halfWidth
	py := dx / length * halfWidth

	// Calculate the four corners of the line rectangle
	x1f, y1f := float64(x1), float64(y1)
	x2f, y2f := float64(x2), float64(y2)

	// Calculate the four corners of the line rectangle
	corners := [][2]float64{
		{x1f + px, y1f + py},
		{x1f - px, y1f - py},
		{x2f - px, y2f - py},
		{x2f + px, y2f + py},
	}

	// Convert corners to integers
	points := make([][2]int, 4)
	for i, corner := range corners {
		points[i] = [2]int{
			int(math.Round(corner[0])),
			int(math.Round(corner[1])),
		}
	}

	// Draw the filled polygon
	c.Polygon(points, color, true)
}
