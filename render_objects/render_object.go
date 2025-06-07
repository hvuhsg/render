package render_objects

import (
	"github.com/hvuhsg/render/canvas"
	"github.com/hvuhsg/render/types"
)

type RenderObject interface {
	Paint(canvas *canvas.Canvas)
	Size(parentSize types.Size) types.Size
}
