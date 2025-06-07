package canvas

import (
	"image"
	"image/color"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/hvuhsg/render/types"
	"golang.org/x/image/font/gofont/goregular"
)

type TextPainter struct {
	Font      *truetype.Font
	FontSize  float64
	TextColor color.Color
}

func NewTextPainter() *TextPainter {
	font, _ := truetype.Parse(goregular.TTF)
	return &TextPainter{
		Font:      font,
		FontSize:  12,
		TextColor: color.Black,
	}
}

func (c *Canvas) DrawText(text string, x, y int, painter *TextPainter) {
	if painter == nil {
		painter = NewTextPainter()
	}

	// Create a new freetype context
	ctx := freetype.NewContext()
	ctx.SetDPI(72)
	ctx.SetFont(painter.Font)
	ctx.SetFontSize(painter.FontSize)
	ctx.SetClip(image.Rect(0, 0, c.Size.Width, c.Size.Height))
	ctx.SetDst(c.Img)
	ctx.SetSrc(image.NewUniform(painter.TextColor))

	// Calculate the baseline position
	// The y position is the baseline, so we need to adjust for the font height
	baseline := y + int(painter.FontSize)
	pt := freetype.Pt(x, baseline)
	ctx.DrawString(text, pt)
}

func (c *Canvas) MeasureText(text string, painter *TextPainter) types.Size {
	if painter == nil {
		painter = NewTextPainter()
	}

	// Create a new freetype context for measurement
	ctx := freetype.NewContext()
	ctx.SetDPI(72)
	ctx.SetFont(painter.Font)
	ctx.SetFontSize(painter.FontSize)

	// Get the bounds of the text
	bounds, _ := ctx.DrawString(text, freetype.Pt(0, 0))

	// Add some padding for better text rendering
	return types.Size{
		Width:  int(bounds.X.Round()) + 4, // Add padding
		Height: int(painter.FontSize) + 4, // Use font size plus padding for height
	}
}
