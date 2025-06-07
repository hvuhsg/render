package canvas

import "image/color"

func (c *Canvas) Rectangle(x, y, w, h int, color color.RGBA, fill bool) {
	c.assertPointInBounds(x, y)
	c.assertPointInBounds(x+w, y)
	c.assertPointInBounds(x, y+h)
	c.assertPointInBounds(x+w, y+h)

	c.Polygon([][2]int{
		{x, y},
		{x + w, y},
		{x + w, y + h},
		{x, y + h},
	}, color, fill)
}
