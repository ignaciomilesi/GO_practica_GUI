package widget

import (
	"fmt"
	"image/color"
	"os"
	"strings"

	g "github.com/AllenDang/giu"
)

var (
	TreeNode g.Layout
)

type Archivo struct {
	Nombre string
	Tipo   string
	Path   string
}

type listadoArchvio struct{}

func NewListadoArchvio() *listadoArchvio {
	return &listadoArchvio{}
}

func (l *listadoArchvio) Build() {

	g.Style().
		SetFontSize(13).To(
		g.Label("Lista de archivos en la carpeta:"),
	).Build()

	TreeNode.Build()

}

func nuevoListarArchivos(path string) (lll g.Layout) {
	nombrePDFSeleccionado = ""
	TreeNode = generarTreeNode(ubicacionCapeta, true)

	return TreeNode
}

func generarTreeNode(path string, marcarAlBuscar bool) (lll g.Layout) {

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {

		if file.IsDir() {

			sublistado := generarTreeNode(path+"\\"+file.Name(), marcarAlBuscar)

			lll = append(lll, g.TreeNode(file.Name()).Flags(g.TreeNodeFlagsDefaultOpen).Layout(sublistado...))

		} else {

			nuevoArchivo := Archivo{
				Nombre: file.Name(),
				Tipo:   strings.ToLower(file.Name()[strings.LastIndex(file.Name(), ".")+1:]),
				Path:   path + "\\" + file.Name(),
			}

			if nuevoArchivo.Tipo == "db" {
				continue
			}
			if nuevoArchivo.Tipo == "pdf" {

				if marcarAlBuscar && nombrePDFSeleccionado == "" {
					nombrePDFSeleccionado = nuevoArchivo.Path
				}

				lll = append(lll, g.RadioButton(nuevoArchivo.Nombre, nombrePDFSeleccionado == nuevoArchivo.Path).OnChange(func() {

					nombrePDFSeleccionado = nuevoArchivo.Path

					TreeNode = generarTreeNode(ubicacionCapeta, false)
				}))

			} else {
				lll = append(lll,
					g.Style().SetColor(g.StyleColorText, color.RGBA{100, 100, 100, 255}).To(
						g.Label(" - "+nuevoArchivo.Nombre)))
			}
		}
	}

	return

}
