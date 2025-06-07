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
	box := &render_objects.ColoredBox{
		Color:  cv.Red,
		Width:  100,
		Height: 100,
	}

	padding := &render_objects.Padding{
		Child:  box,
		Top:    100,
		Left:   100,
		Right:  100,
		Bottom: 100,
	}

	column := &render_objects.Column{
		Children: []render_objects.RenderObject{padding, box},
		Sizing:   types.MainAxisSizeMin,
	}

	// Center everything in the canvas
	align := &render_objects.Align{
		Child: column,
		Align: render_objects.AlignTopLeft,
	}

	return align
}
