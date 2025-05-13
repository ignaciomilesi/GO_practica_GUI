package main

import (
	"prueba-gui/widget"

	g "github.com/AllenDang/giu"
)

var (
	ubicacionCapeta string

	sashPos1 float32 = 200
)

func main() {
	wnd := g.NewMasterWindow("Cargar plano", 640, 480, g.MasterWindowFlagsNotResizable)

	wnd.Run(esquema)
}

func esquema() {

	g.SingleWindowWithMenuBar().Layout(

		widget.NewBarraMenu(),

		widget.NewSelectorCarpeta(),

		g.SplitLayout(g.DirectionVertical, &sashPos1,
			g.Layout{},
			g.Layout{

				widget.NewListadoArchvio(),
			},
		).SplitRefType(g.SplitRefRight),
	)

}
