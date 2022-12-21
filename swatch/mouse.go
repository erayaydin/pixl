package swatch

import "fyne.io/fyne/v2/driver/desktop"

func (swatch *Swatch) MouseDown(_ *desktop.MouseEvent) {
	swatch.clickHandler(swatch)
	swatch.Selected = true
	swatch.Refresh()
}

func (swatch *Swatch) MouseUp(_ *desktop.MouseEvent) {}
