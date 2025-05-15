package widget

import (
	g "github.com/AllenDang/giu"
)

var (
	nombrePDFSeleccionado string
)

type visorPDFSeleccionado struct{}

func NewVisorPdfSeleccionado() *visorPDFSeleccionado {
	return &visorPDFSeleccionado{}
}

func (v *visorPDFSeleccionado) Build() {

	g.Style().
		SetFontSize(13).To(
		g.Label("Plano a cargar en el compartido:"),
	).Build()

	g.InputText(&nombrePDFSeleccionado).Flags(g.InputTextFlagsReadOnly).Size(g.Auto).Build()
}
