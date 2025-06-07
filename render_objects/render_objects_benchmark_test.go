package render_objects

import (
	"image/color"
	"testing"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

func BenchmarkColoredBox(b *testing.B) {
	canvas := cv.NewCanvas(types.Size{Width: 1000, Height: 1000}, false)
	red := color.RGBA{255, 0, 0, 255}
	box := &ColoredBox{Width: 100, Height: 100, Color: red}

	b.Run("Paint", func(b *testing.B) {
		for b.Loop() {
			box.Paint(canvas)
		}
	})

	b.Run("Size", func(b *testing.B) {
		for b.Loop() {
			box.Size(canvas.Size)
		}
	})
}

func BenchmarkText(b *testing.B) {
	canvas := cv.NewCanvas(types.Size{Width: 1000, Height: 1000}, false)
	black := color.Black
	text := NewText("Hello, World!", black, 24, "default")

	b.Run("Paint", func(b *testing.B) {
		for b.Loop() {
			text.Paint(canvas)
		}
	})

	b.Run("Size", func(b *testing.B) {
		for b.Loop() {
			text.Size(canvas.Size)
		}
	})
}

func BenchmarkRow(b *testing.B) {
	canvas := cv.NewCanvas(types.Size{Width: 1000, Height: 1000}, false)
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	// Create boxes for the row
	boxes := []RenderObject{
		&ColoredBox{Width: 100, Height: 100, Color: red},
		&ColoredBox{Width: 100, Height: 100, Color: blue},
		&ColoredBox{Width: 100, Height: 100, Color: red},
	}

	row := &Row{
		Children:  boxes,
		Alignment: types.MainAxisAlignmentSpaceBetween,
		Sizing:    types.MainAxisSizeMax,
	}

	b.Run("Paint", func(b *testing.B) {
		for b.Loop() {
			row.Paint(canvas)
		}
	})

	b.Run("Size", func(b *testing.B) {
		for b.Loop() {
			row.Size(canvas.Size)
		}
	})
}

func BenchmarkColumn(b *testing.B) {
	canvas := cv.NewCanvas(types.Size{Width: 1000, Height: 1000}, false)
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	// Create boxes for the column
	boxes := []RenderObject{
		&ColoredBox{Width: 100, Height: 100, Color: red},
		&ColoredBox{Width: 100, Height: 100, Color: blue},
		&ColoredBox{Width: 100, Height: 100, Color: red},
	}

	column := &Column{
		Children:  boxes,
		Alignment: types.MainAxisAlignmentSpaceBetween,
		Sizing:    types.MainAxisSizeMax,
	}

	b.Run("Paint", func(b *testing.B) {
		for b.Loop() {
			column.Paint(canvas)
		}
	})

	b.Run("Size", func(b *testing.B) {
		for b.Loop() {
			column.Size(canvas.Size)
		}
	})
}

func BenchmarkStack(b *testing.B) {
	canvas := cv.NewCanvas(types.Size{Width: 1000, Height: 1000}, false)
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	green := color.RGBA{0, 255, 0, 255}

	// Create boxes for the stack
	boxes := []RenderObject{
		&ColoredBox{Width: 200, Height: 200, Color: red},
		&ColoredBox{Width: 150, Height: 150, Color: blue},
		&ColoredBox{Width: 100, Height: 100, Color: green},
	}

	stack := &Stack{
		Children: boxes,
	}

	b.Run("Paint", func(b *testing.B) {
		for b.Loop() {
			stack.Paint(canvas)
		}
	})

	b.Run("Size", func(b *testing.B) {
		for b.Loop() {
			stack.Size(canvas.Size)
		}
	})
}

func BenchmarkAlign(b *testing.B) {
	canvas := cv.NewCanvas(types.Size{Width: 1000, Height: 1000}, false)
	red := color.RGBA{255, 0, 0, 255}
	box := &ColoredBox{Width: 100, Height: 100, Color: red}

	align := &Align{
		Child: box,
		Align: AlignCenter,
	}

	b.Run("Paint", func(b *testing.B) {
		for b.Loop() {
			align.Paint(canvas)
		}
	})

	b.Run("Size", func(b *testing.B) {
		for b.Loop() {
			align.Size(canvas.Size)
		}
	})
}

func BenchmarkPainter(b *testing.B) {
	canvas := cv.NewCanvas(types.Size{Width: 1000, Height: 1000}, false)
	red := color.RGBA{255, 0, 0, 255}

	painter := &Painter{
		Painter: func(canvas *cv.Canvas) {
			canvas.Circle(500, 500, 100, red, true)
			canvas.Rectangle(200, 200, 100, 100, red, false)
		},
		Width:  1000,
		Height: 1000,
	}

	b.Run("Paint", func(b *testing.B) {
		for b.Loop() {
			painter.Paint(canvas)
		}
	})

	b.Run("Size", func(b *testing.B) {
		for b.Loop() {
			painter.Size(canvas.Size)
		}
	})
}
