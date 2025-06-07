package render_objects

import (
	"image/color"
	"testing"

	"github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

// NewColor helper function
func NewColor(r, g, b uint8) color.RGBA {
	return color.RGBA{r, g, b, 255}
}

func createComplexRenderTree() RenderObject {
	// Create colored boxes
	headerBox := &ColoredBox{
		Color:  NewColor(200, 200, 200),
		Width:  100,
		Height: 40,
	}

	mainBox := &ColoredBox{
		Color:  NewColor(240, 240, 240),
		Width:  400,
		Height: 300,
	}

	box1 := &ColoredBox{
		Color:  NewColor(150, 150, 150),
		Width:  50,
		Height: 50,
	}

	box2 := &ColoredBox{
		Color:  NewColor(150, 150, 150),
		Width:  50,
		Height: 50,
	}

	box3 := &ColoredBox{
		Color:  NewColor(150, 150, 150),
		Width:  50,
		Height: 50,
	}

	footerBox := &ColoredBox{
		Color:  NewColor(180, 180, 180),
		Width:  200,
		Height: 30,
	}

	// Create text elements
	headerText := NewText("Header Text", color.Black, 24, "Arial")
	leftText := NewText("Left Aligned", color.Black, 16, "Arial")
	centerText := NewText("Center Aligned", color.Black, 16, "Arial")
	rightText := NewText("Right Aligned", color.Black, 16, "Arial")
	footerText := NewText("Footer", color.Black, 14, "Arial")

	// Create aligned text elements
	alignedText1 := &Align{
		Child: leftText,
		Align: AlignLeftCenter,
	}
	alignedText2 := &Align{
		Child: centerText,
		Align: AlignCenter,
	}
	alignedText3 := &Align{
		Child: rightText,
		Align: AlignRightCenter,
	}

	// Create box row
	boxRow := &Row{
		Alignment: types.MainAxisAlignmentStart,
		Sizing:    types.MainAxisSizeMax,
		Children:  []RenderObject{box1, box2, box3},
	}

	// Create nested column
	nestedColumn := &Column{
		Alignment: types.MainAxisAlignmentStart,
		Sizing:    types.MainAxisSizeMax,
		Children:  []RenderObject{alignedText1, alignedText2, alignedText3, boxRow},
	}

	// Create main content stack
	mainContent := &Stack{
		Children: []RenderObject{
			&Border{
				Child: mainBox,
				Width: 2,
				Color: NewColor(100, 100, 100),
			},
			nestedColumn,
		},
	}

	// Create header row
	header := &Row{
		Alignment: types.MainAxisAlignmentStart,
		Sizing:    types.MainAxisSizeMax,
		Children:  []RenderObject{headerText, headerBox},
	}

	// Create footer row with border
	footer := &Border{
		Child: &Row{
			Alignment: types.MainAxisAlignmentStart,
			Sizing:    types.MainAxisSizeMax,
			Children:  []RenderObject{footerText, footerBox},
		},
		Width: 1,
		Color: NewColor(150, 150, 150),
	}

	// Create root column
	root := &Column{
		Alignment: types.MainAxisAlignmentStart,
		Sizing:    types.MainAxisSizeMax,
		Children:  []RenderObject{header, mainContent, footer},
	}

	return root
}

func BenchmarkComplexRenderTree(b *testing.B) {
	renderTree := createComplexRenderTree()
	canvas := canvas.NewCanvas(types.Size{Width: 800, Height: 600}, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderTree.Paint(canvas)
	}
}

func BenchmarkComplexRenderTreeSize(b *testing.B) {
	renderTree := createComplexRenderTree()
	parentSize := types.Size{Width: 800, Height: 600}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		renderTree.Size(parentSize)
	}
}
