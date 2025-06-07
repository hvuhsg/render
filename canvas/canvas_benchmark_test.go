package canvas

import (
	"image/color"
	"testing"

	"github.com/hvuhsg/render/types"
)

func BenchmarkRectangle(b *testing.B) {
	canvas := NewCanvas(types.Size{Width: 1000, Height: 1000})
	red := color.RGBA{255, 0, 0, 255}

	b.Run("Filled", func(b *testing.B) {
		for b.Loop() {
			canvas.Rectangle(0, 0, 100, 100, red, true)
		}
	})

	b.Run("Unfilled", func(b *testing.B) {
		for b.Loop() {
			canvas.Rectangle(0, 0, 100, 100, red, false)
		}
	})
}

func BenchmarkCircle(b *testing.B) {
	canvas := NewCanvas(types.Size{Width: 1000, Height: 1000})
	blue := color.RGBA{0, 0, 255, 255}

	b.Run("Filled", func(b *testing.B) {
		for b.Loop() {
			canvas.Circle(500, 500, 50, blue, true)
		}
	})

	b.Run("Unfilled", func(b *testing.B) {
		for b.Loop() {
			canvas.Circle(500, 500, 50, blue, false)
		}
	})
}

func BenchmarkLine(b *testing.B) {
	canvas := NewCanvas(types.Size{Width: 1000, Height: 1000})
	green := color.RGBA{0, 255, 0, 255}

	b.Run("Horizontal", func(b *testing.B) {
		for b.Loop() {
			canvas.Line(100, 500, 900, 500, green, 2)
		}
	})

	b.Run("Vertical", func(b *testing.B) {
		for b.Loop() {
			canvas.Line(500, 100, 500, 900, green, 2)
		}
	})

	b.Run("Diagonal", func(b *testing.B) {
		for b.Loop() {
			canvas.Line(100, 100, 900, 900, green, 2)
		}
	})
}

func BenchmarkPolygon(b *testing.B) {
	canvas := NewCanvas(types.Size{Width: 1000, Height: 1000})
	purple := color.RGBA{128, 0, 128, 255}

	// Create a hexagon
	points := [][2]int{
		{500, 100}, // top
		{700, 200}, // top right
		{700, 400}, // bottom right
		{500, 500}, // bottom
		{300, 400}, // bottom left
		{300, 200}, // top left
	}

	b.Run("Filled", func(b *testing.B) {
		for b.Loop() {
			canvas.Polygon(points, purple, true)
		}
	})
}

func BenchmarkText(b *testing.B) {
	canvas := NewCanvas(types.Size{Width: 1000, Height: 1000})
	black := color.Black
	painter := NewTextPainter()
	painter.TextColor = black
	painter.FontSize = 24

	b.Run("Short", func(b *testing.B) {
		for b.Loop() {
			canvas.DrawText("Hello", 100, 100, painter)
		}
	})

	b.Run("Long", func(b *testing.B) {
		for b.Loop() {
			canvas.DrawText("This is a longer text that needs more processing", 100, 100, painter)
		}
	})
}

func BenchmarkSubCanvas(b *testing.B) {
	parent := NewCanvas(types.Size{Width: 1000, Height: 1000})
	red := color.RGBA{255, 0, 0, 255}

	b.Run("Create", func(b *testing.B) {
		for b.Loop() {
			sub := parent.SubCanvas(100, 100, types.Size{Width: 200, Height: 200})
			sub.Rectangle(0, 0, 100, 100, red, true)
		}
	})

	b.Run("Nested", func(b *testing.B) {
		for b.Loop() {
			sub1 := parent.SubCanvas(100, 100, types.Size{Width: 200, Height: 200})
			sub2 := sub1.SubCanvas(50, 50, types.Size{Width: 100, Height: 100})
			sub2.Rectangle(0, 0, 50, 50, red, true)
		}
	})
}
