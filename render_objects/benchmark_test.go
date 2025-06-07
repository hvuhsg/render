package render_objects

import (
	"image/color"
	"testing"

	"github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

// BenchmarkRenderObjects runs benchmarks for various render objects
func BenchmarkRenderObjects(b *testing.B) {
	// Create a test canvas
	testCanvas := canvas.NewCanvas(types.Size{Width: 800, Height: 600}, false)

	// Define test colors
	green := color.RGBA{R: 144, G: 238, B: 144, A: 255}
	blue := color.RGBA{R: 173, G: 216, B: 230, A: 255}

	// Benchmark Text rendering
	b.Run("Text", func(b *testing.B) {
		text := NewText("Benchmark Text", green, 36, "default")
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			text.Paint(testCanvas)
		}
	})

	// Benchmark Row layout
	b.Run("Row", func(b *testing.B) {
		row := &Row{
			Children: []RenderObject{
				NewText("Item 1", green, 24, "default"),
				NewText("Item 2", blue, 24, "default"),
			},
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			row.Paint(testCanvas)
		}
	})

	// Benchmark Column layout
	b.Run("Column", func(b *testing.B) {
		col := &Column{
			Children: []RenderObject{
				NewText("Item 1", green, 24, "default"),
				NewText("Item 2", blue, 24, "default"),
			},
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			col.Paint(testCanvas)
		}
	})

	// Benchmark Align
	b.Run("Align", func(b *testing.B) {
		align := &Align{
			Child: NewText("Centered Text", green, 36, "default"),
			Align: AlignCenter,
		}
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			align.Paint(testCanvas)
		}
	})
}

// BenchmarkMemoryAllocations measures memory allocations for render objects
func BenchmarkMemoryAllocations(b *testing.B) {
	testCanvas := canvas.NewCanvas(types.Size{Width: 800, Height: 600}, false)

	// Define test colors
	green := color.RGBA{R: 144, G: 238, B: 144, A: 255}
	blue := color.RGBA{R: 173, G: 216, B: 230, A: 255}
	red := color.RGBA{R: 255, G: 182, B: 193, A: 255}

	b.Run("TextAllocations", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			text := NewText("Test Text", green, 36, "default")
			text.Paint(testCanvas)
		}
	})

	b.Run("ComplexLayoutAllocations", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			col := &Column{
				Children: []RenderObject{
					NewText("Item 1", green, 24, "default"),
					NewText("Item 2", blue, 24, "default"),
				},
			}
			row := &Row{
				Children: []RenderObject{
					col,
					NewText("Side Text", red, 24, "default"),
				},
			}
			row.Paint(testCanvas)
		}
	})
}
