package canvas

import (
	"image/color"
	"sort"
)

// Polygon draws a filled polygon defined by the given points
func (c *Canvas) Polygon(points [][2]int, color color.RGBA, filled bool) {
	if len(points) < 3 {
		return
	}

	if !filled {
		// Draw the outline by connecting points with lines
		for i := 0; i < len(points); i++ {
			j := (i + 1) % len(points)
			c.drawLine(points[i][0], points[i][1], points[j][0], points[j][1], color)
		}
		return
	}

	// Find the bounding box
	minY, maxY := points[0][1], points[0][1]
	for _, p := range points {
		if p[1] < minY {
			minY = p[1]
		}
		if p[1] > maxY {
			maxY = p[1]
		}
	}

	// For each scan line
	for y := minY; y <= maxY; y++ {
		// Find intersections with polygon edges
		var intersections []int
		j := len(points) - 1
		for i := 0; i < len(points); i++ {
			if (points[i][1] > y) != (points[j][1] > y) {
				// Calculate x-coordinate of intersection
				x := points[i][0] + (y-points[i][1])*(points[j][0]-points[i][0])/(points[j][1]-points[i][1])
				intersections = append(intersections, x)
			}
			j = i
		}

		// Sort intersections
		sort.Ints(intersections)

		// Fill between pairs of intersections
		for i := 0; i < len(intersections); i += 2 {
			if i+1 < len(intersections) {
				startX := intersections[i]
				endX := intersections[i+1]
				for x := startX; x <= endX; x++ {
					c.set(x, y, color)
				}
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
