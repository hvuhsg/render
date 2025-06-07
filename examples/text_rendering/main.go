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
	canvas := cv.NewCanvas(types.Size{Width: 800, Height: 600}, false)

	// Create text objects with different properties
	title := render_objects.NewText("Hello, Render!", cv.Red, 48, "default")
	subtitle := render_objects.NewText("A Simple Text Example", cv.Blue, 24, "default")
	body := render_objects.NewText("This demonstrates different text styles and colors", cv.Gray, 18, "default")

	// Arrange text in a column
	textColumn := &render_objects.Column{
		Children: []render_objects.RenderObject{
			title,
			subtitle,
			body,
		},
	}

	// Center the text column
	align := &render_objects.Align{
		Child: textColumn,
		Align: render_objects.AlignCenter,
	}

	// Render the text
	align.Paint(canvas)

	// Save the result
	file, _ := os.Create("text_rendering.png")
	defer file.Close()
	png.Encode(file, canvas.Img)
}
