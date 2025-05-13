package widget

import (
	g "github.com/AllenDang/giu"
)

type barraMenu struct{}

func NewBarraMenu() *barraMenu {
	return &barraMenu{}
}

func (b *barraMenu) Build() {

	g.MenuBar().Layout(
		g.Menu("Configuración").Layout(
			g.MenuItem("Configurar ubicaciones"),
		),
	).Build()
}
