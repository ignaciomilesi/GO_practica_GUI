package widget

import (
	"fmt"
	"prueba-gui/internal/estructuraCarpeta"
	"prueba-gui/internal/seleccion"

	g "github.com/AllenDang/giu"
	"github.com/sqweek/dialog"
)

type selectorCarpeta struct{}

func NuevoSelectorCarpeta() *selectorCarpeta {
	return &selectorCarpeta{}
}

func (s *selectorCarpeta) Build() {
	g.Row(
		g.Style().
			SetFontSize(13).To(
			g.Label("Carpeta con sólidos y PDF:"),
		),
		g.Style().
			SetFontSize(10).To(
			g.Label(" (Seleccionar o arrastrar carpeta)"),
		),
	).Build()

	g.Row(
		g.InputText(&seleccion.Carpeta.Path).Flags(g.InputTextFlagsReadOnly).Size(890),
		g.Button("Seleccionar").OnClick(s.buscarCarpeta),
	).Build()
}

func (s *selectorCarpeta) buscarCarpeta() {

	carpeta, err := dialog.Directory().Browse()

	if err != nil {
		fmt.Print(err)
		return
	}

	seleccion.Carpeta = estructuraCarpeta.NuevaCarpeta(carpeta)

	ActualizarLayoutListadoArchivo()

}
