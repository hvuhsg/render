package main

import (
	"image/png"
	"os"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/render_objects"
	"github.com/hvuhsg/render/types"
)

func main() {
	// Create a new canvas
	canvas := cv.NewCanvas(types.Size{Width: 800, Height: 600})

	// Create some colored boxes
	redBox := &render_objects.ColoredBox{Width: 100, Height: 100, Color: cv.Red}
	blueBox := &render_objects.ColoredBox{Width: 100, Height: 100, Color: cv.Blue}
	greenBox := &render_objects.ColoredBox{Width: 100, Height: 100, Color: cv.Green}

	// Create a row of boxes
	row := &render_objects.Row{
		Children: []render_objects.RenderObject{
			redBox,
			blueBox,
			greenBox,
		},
		Alignment: types.MainAxisAlignmentSpaceBetween,
		Sizing:    types.MainAxisSizeMax,
	}

	// Create a column with some text
	textColumn := &render_objects.Column{
		Children: []render_objects.RenderObject{
			render_objects.NewText("Layout Example", cv.Black, 24, "default"),
			row,
		},
	}

	// Center everything
	align := &render_objects.Align{
		Child: textColumn,
		Align: render_objects.AlignCenter,
	}

	// Render the layout
	align.Paint(canvas)

	// Save the result
	file, _ := os.Create("layout_composition.png")
	defer file.Close()
	png.Encode(file, canvas.Img)
}
