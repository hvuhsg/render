package main

import (
	"image/png"
	"os"

	cv "github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/render_objects"
	"github.com/hvuhsg/render/types"
)

func saveToFile(canvas *cv.Canvas, filename string) error {
	img := canvas.Img

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	png.Encode(file, img)

	return nil
}

func main() {
	canvas := cv.NewCanvas(types.Size{Width: 800, Height: 600}, false)
	renderTree := createRenderTree()
	renderTree.Paint(canvas)
	saveToFile(canvas, "out.png")
}

func createRenderTree() render_objects.RenderObject {
	// Create a column of text elements with some spacing
	polygon := &render_objects.Painter{
		Painter: func(canvas *cv.Canvas) {
			canvas.Polygon([][2]int{
				{100, 100},
				{235, 189},
				{129, 599},
			}, cv.Red, false)
		},
		Width:  800,
		Height: 600,
	}

	// Center everything in the canvas
	align := &render_objects.Align{
		Child: polygon,
		Align: render_objects.AlignCenter,
	}

	return align
}
