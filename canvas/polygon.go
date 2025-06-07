package canvas

import "image/color"

// Polygon draws a filled polygon defined by the given points
func (c *Canvas) Polygon(points [][2]int, color color.RGBA, filled bool) {
	if len(points) < 3 {
		return
	}

	// Find the bounding box
	minX, minY := points[0][0], points[0][1]
	maxX, maxY := points[0][0], points[0][1]
	for _, p := range points {
		if p[0] < minX {
			minX = p[0]
		}
		if p[0] > maxX {
			maxX = p[0]
		}
		if p[1] < minY {
			minY = p[1]
		}
		if p[1] > maxY {
			maxY = p[1]
		}
	}

	// For each point in the bounding box, check if it's inside the polygon
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if c.isPointInPolygon(x, y, points) {
				c.set(x, y, color)
			}
		}
	}
}

// isPointInPolygon uses the ray casting algorithm to determine if a point is inside a polygon
func (c *Canvas) isPointInPolygon(x, y int, points [][2]int) bool {
	inside := false
	j := len(points) - 1
	for i := range len(points) {
		if (points[i][1] > y) != (points[j][1] > y) &&
			x < (points[j][0]-points[i][0])*(y-points[i][1])/(points[j][1]-points[i][1])+points[i][0] {
			inside = !inside
		}
		j = i
	}
	return inside
}
