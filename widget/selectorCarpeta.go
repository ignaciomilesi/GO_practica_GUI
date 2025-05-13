package widget

import (
	"fmt"

	g "github.com/AllenDang/giu"
	"github.com/sqweek/dialog"
)

type selectorCarpeta struct{}

var (
	ubicacionCapeta string
)

func NewSelectorCarpeta() *selectorCarpeta {
	return &selectorCarpeta{}
}

func (s *selectorCarpeta) Build() {
	g.Style().
		SetFontSize(13).To(
		g.Label("Carpeta con sólidos y PDF:"),
	).Build()

	g.Row(
		g.InputText(&ubicacionCapeta).Flags(g.InputTextFlagsReadOnly).Size(530),
		g.Button("Seleccionar").OnClick(s.buscarCarpeta),
	).Build()

	g.Style().
		SetFontSize(10).To(
		g.Label("(Seleccionar o arrastrar carpeta)"),
	).Build()

}

func (s *selectorCarpeta) buscarCarpeta() {
	carpeta, err := dialog.Directory().Browse()

	if err != nil {
		fmt.Print(err)
		return
	}

	ubicacionCapeta = carpeta

}
