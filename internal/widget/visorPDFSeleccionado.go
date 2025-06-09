package widget

import (
	"prueba-gui/internal/seleccion"

	g "github.com/AllenDang/giu"
)

type visorPDFSeleccionado struct{}

func NuevoVisorPDFSeleccionado() *visorPDFSeleccionado {
	return &visorPDFSeleccionado{}
}

func (v *visorPDFSeleccionado) Build() {

	g.Style().
		SetFontSize(13).To(
		g.Label("Plano a cargar en el compartido:"),
	).Build()

	g.InputText(&seleccion.PDF.Nombre).Flags(g.InputTextFlagsReadOnly).Size(g.Auto).Build()
}
