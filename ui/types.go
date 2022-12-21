package ui

import (
	"fyne.io/fyne/v2"
	"github.com/erayaydin/pixl/apptype"
	"github.com/erayaydin/pixl/pxcanvas"
	"github.com/erayaydin/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
