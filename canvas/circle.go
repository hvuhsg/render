package canvas

import (
	"image/color"
	"math"
)

func (c *Canvas) Circle(x, y, r int, color color.RGBA, fill bool) {
	c.assertPointInBounds(x, y)
	c.assertPointInBounds(x+r, y)
	c.assertPointInBounds(x, y+r)
	c.assertPointInBounds(x-r, y)
	c.assertPointInBounds(x, y-r)

	if fill {
		// Fill the circle by drawing horizontal lines
		for dy := -r; dy <= r; dy++ {
			// Calculate the width of the line at this y position
			dx := int(math.Sqrt(float64(r*r - dy*dy)))
			// Draw the horizontal line
			for i := -dx; i <= dx; i++ {
				c.set(x+i, y+dy, color)
			}
		}
	} else {
		// Draw just the outline
		for i := range 360 {
			angle := float64(i) * math.Pi / 180
			nx := x + int(float64(r)*math.Cos(angle))
			ny := y + int(float64(r)*math.Sin(angle))
			c.set(nx, ny, color)
		}
	}
}
