package main

import (
	"prueba-gui/internal/widget"

	g "github.com/AllenDang/giu"
)

func main() {

	wnd := g.NewMasterWindow("Cargar plano", 1000, 480, g.MasterWindowFlagsNotResizable)

	wnd.Run(esquema)
}

var (
	sashPos1 float32 = 400
)

func esquema() {

	g.SingleWindowWithMenuBar().Layout(

		widget.NuevoBarraMenu(),

		widget.NuevoSelectorCarpeta(),

		g.Dummy(5, 5),

		widget.NuevoVisorPDFSeleccionado(),

		g.SplitLayout(g.DirectionVertical, &sashPos1,
			g.Layout{},
			g.Layout{

				widget.NuevoListadoArchivo(),
			},
		).SplitRefType(g.SplitRefRight),
	)

}
