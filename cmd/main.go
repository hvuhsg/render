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

type CustomRenderObject struct{}

func (o *CustomRenderObject) Paint(canvas *cv.Canvas) {
	canvas.Circle(500, 500, 500, cv.Yellow, true)
	canvas.Line(0, 0, canvas.Size.Width, canvas.Size.Height, cv.Red, 40)
	canvas.Rectangle(100, 100, 100, 100, cv.Blue, true)
}

func main() {
	canvas := cv.NewCanvas(types.Size{Width: 800, Height: 600})
	renderTree := createRenderTree()
	renderTree.Paint(canvas)
	saveToFile(canvas, "out.png")
}

func createRenderTree() render_objects.RenderObject {
	// Create text objects with different properties
	titleText := render_objects.NewText("Hello, World!", cv.LightGreen, 36, "default")
	subtitleText := render_objects.NewText("Welcome to Render", cv.Blue, 24, "default")
	infoText := render_objects.NewText("A simple rendering library", cv.Gray, 18, "default")

	// Create a column of text elements with some spacing
	textColumn := &render_objects.Column{
		Children: []render_objects.RenderObject{
			titleText,
			subtitleText,
			infoText,
		},
	}

	// Create the existing shapes
	redBox := &render_objects.ColoredBox{Width: 200, Height: 100, Color: cv.Red}
	blueBox := &render_objects.ColoredBox{Width: 200, Height: 100, Color: cv.Blue}
	multiColoredBox := &render_objects.Column{
		Children: []render_objects.RenderObject{
			redBox,
			blueBox,
		},
	}
	circle := &render_objects.Painter{
		Painter: func(canvas *cv.Canvas) {
			canvas.Circle(100, 100, 100, cv.Yellow, true)
		},
		Width:  200,
		Height: 200,
	}

	// Create a row with text and shapes
	contentRow := &render_objects.Row{
		Children: []render_objects.RenderObject{
			multiColoredBox,
			circle,
			textColumn,
		},
		Alignment: types.MainAxisAlignmentSpaceBetween,
		Sizing:    types.MainAxisSizeMax,
	}

	// Center everything in the canvas
	align := &render_objects.Align{
		Child: contentRow,
		Align: render_objects.AlignCenter,
	}

	return align
}
