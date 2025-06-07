package main

import (
	"image/png"
	"os"

	"github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/render_objects"
	"github.com/hvuhsg/render/types"
)

func main() {
	// Create a new canvas
	canvas_ := canvas.NewCanvas(types.Size{Width: 800, Height: 600}, false)

	// Create a text element
	text := render_objects.NewText("Hello, World!", canvas.Purple, 36, "default")

	// Center the text
	align := &render_objects.Align{
		Child: text,
		Align: render_objects.AlignCenter,
	}

	// Render and save
	align.Paint(canvas_)

	// Save the result
	file, _ := os.Create("result.png")
	defer file.Close()
	png.Encode(file, canvas_.Img)
}
