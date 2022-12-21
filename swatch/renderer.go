package swatch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

type Renderer struct {
	square  canvas.Rectangle
	objects []fyne.CanvasObject
	parent  *Swatch
}

func (renderer *Renderer) MinSize() fyne.Size {
	return renderer.square.MinSize()
}

func (renderer *Renderer) Layout(size fyne.Size) {
	renderer.objects[0].Resize(size)
}

func (renderer *Renderer) Refresh() {
	renderer.Layout(fyne.NewSize(20, 20))
	renderer.square.FillColor = renderer.parent.Color
	if renderer.parent.Selected {
		renderer.square.StrokeWidth = 3
		renderer.square.StrokeColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	} else {
		renderer.square.StrokeWidth = 0
	}
	renderer.objects[0] = &renderer.square
	canvas.Refresh(renderer.parent)
}

func (renderer *Renderer) Objects() []fyne.CanvasObject {
	return renderer.objects
}

func (renderer *Renderer) Destroy() {}
