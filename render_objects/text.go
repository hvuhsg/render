package render_objects

import (
	"image/color"

	"github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

type Text struct {
	text     string
	color    color.Color
	fontSize float64
	fontName string
	size     types.Size
}

func NewText(text string, color color.Color, fontSize float64, fontName string) *Text {
	return &Text{
		text:     text,
		color:    color,
		fontSize: fontSize,
		fontName: fontName,
		size:     types.Size{Width: 0, Height: 0}, // Will be calculated in Paint
	}
}

func (t *Text) Paint(c *canvas.Canvas) {
	// Create a text painter with the text properties
	painter := canvas.NewTextPainter()
	painter.TextColor = t.color
	painter.FontSize = t.fontSize
	// Note: font name is ignored for now as we only support the default font

	// Draw the text
	c.DrawText(t.text, 0, 0, painter)
}

func (t *Text) Size(parentSize types.Size) types.Size {
	if t.size.Width == 0 || t.size.Height == 0 {
		// Create a temporary canvas for measurement
		tempCanvas := canvas.NewCanvas(types.Size{Width: 1000, Height: 1000})
		painter := canvas.NewTextPainter()
		painter.TextColor = t.color
		painter.FontSize = t.fontSize
		t.size = tempCanvas.MeasureText(t.text, painter)
	}
	return t.size
}
