package widget

import (
	g "github.com/AllenDang/giu"
)

var (
	pdfSeleccionado int
)

type listadoArchvio struct{}

func NewListadoArchvio() *listadoArchvio {
	return &listadoArchvio{}
}

func (l *listadoArchvio) Build() {

	g.Style().
		SetFontSize(13).To(
		g.Label("Lista de archivos en la carpeta:"),
	).Build()

	g.TreeNode("Carpeta1").Layout(
		g.RadioButton("Prueba", pdfSeleccionado == 0).OnChange(func() { pdfSeleccionado = 0 }),
		g.RadioButton("Prueba", pdfSeleccionado == 1).OnChange(func() { pdfSeleccionado = 1 }),
	).Build()
	g.TreeNode("Carpeta2").Layout(
		g.RadioButton("Prueba", pdfSeleccionado == 2).OnChange(func() { pdfSeleccionado = 2 }),
		g.RadioButton("Prueba", pdfSeleccionado == 3).OnChange(func() { pdfSeleccionado = 3 }),
	).Build()
	g.RadioButton("Prueba", pdfSeleccionado == 4).OnChange(func() { pdfSeleccionado = 4 }).Build()
	g.RadioButton("Prueba", pdfSeleccionado == 5).OnChange(func() { pdfSeleccionado = 5 }).Build()

}
