package canvas

import (
	"errors"
	"image"
	"image/color"

	"github.com/hvuhsg/render/types"
)

var ErrOutOfBounds = errors.New("object is trying to be painted out of bounds")

type Canvas struct {
	Img    *image.RGBA
	Size   types.Size
	offset image.Point
}

func NewCanvas(size types.Size) *Canvas {
	img := image.NewRGBA(image.Rect(0, 0, size.Width, size.Height))
	return &Canvas{
		Img:    img,
		Size:   size,
		offset: image.Point{X: 0, Y: 0},
	}
}

func (c *Canvas) SubCanvas(x, y int, size types.Size) *Canvas {
	return &Canvas{
		Img:    c.Img,
		Size:   size,
		offset: image.Point{X: c.offset.X + x, Y: c.offset.Y + y},
	}
}

func (c *Canvas) set(x, y int, color color.RGBA) {
	c.assertPointInBounds(x, y)
	c.Img.Set(c.offset.X+x, c.offset.Y+y, color)
}

func (c *Canvas) get(x, y int) color.Color {
	c.assertPointInBounds(x, y)
	return c.Img.At(c.offset.X+x, c.offset.Y+y)
}

func (c *Canvas) assertPointInBounds(x, y int) {
	inBounds := x >= 0 && x < c.Size.Width && y >= 0 && y < c.Size.Height
	if !inBounds {
		panic(ErrOutOfBounds)
	}
}

// Should only be used for drawing an already rendered canvas onto this canvas
// When rendering child render objects, use SubCanvas instead
func (c *Canvas) DrawCanvas(other *Canvas, x, y int) {
	// Check if the other canvas would be drawn out of bounds
	c.assertPointInBounds(x, y)
	c.assertPointInBounds(x+other.Size.Width-1, y+other.Size.Height-1)

	// Draw each pixel from the other canvas onto this canvas
	for i := range other.Size.Width {
		for j := range other.Size.Height {
			c.set(x+i, y+j, other.get(i, j).(color.RGBA))
		}
	}
}
