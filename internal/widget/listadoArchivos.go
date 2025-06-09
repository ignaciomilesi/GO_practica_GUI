package widget

import (
	"image/color"

	"prueba-gui/internal/estructuraCarpeta"
	"prueba-gui/internal/seleccion"

	g "github.com/AllenDang/giu"
)

type listadoArchivo struct{}

var layoutListadoArchivo g.Layout

func NuevoListadoArchivo() *listadoArchivo {
	return &listadoArchivo{}
}

func ActualizarLayoutListadoArchivo() {
	layoutListadoArchivo = NuevoListadoArchivo().generarTreeNode(seleccion.Carpeta)
}

func (l *listadoArchivo) Build() {

	g.Style().
		SetFontSize(13).To(
		g.Label("Lista de archivos en la carpeta:"),
	).Build()

	layoutListadoArchivo.Build()

}

func (l *listadoArchivo) generarTreeNode(carpetaSeleccionada estructuraCarpeta.Carpeta) (layoutTemp g.Layout) {

	// recorremos las subcarpetas
	for _, subcarpeta := range carpetaSeleccionada.SubCarpetas {

		//genero el sub-listado (iteración), y lo agrego como un treeNode

		sublistado := l.generarTreeNode(subcarpeta)

		layoutTemp = append(layoutTemp, g.TreeNode(subcarpeta.Nombre).Flags(g.TreeNodeFlagsDefaultOpen).Layout(sublistado...))

	}

	// recorremos los archivos
	for _, archivo := range carpetaSeleccionada.Archivos {

		// descarto los archivos internos de sistema
		if archivo.Tipo == "db" {
			continue
		}

		if archivo.Tipo == "pdf" {

			if seleccion.PDF.Nombre == "" {

				seleccion.PDF = archivo

			}

			// al ser del tipo pdf agrego un  radioButon
			layoutTemp = append(layoutTemp, g.RadioButton(archivo.Nombre, seleccion.PDF.Path == archivo.Path).OnChange(func() {

				seleccion.PDF = archivo

				// busco actualizar el layout por ello quito el marcar
				layoutListadoArchivo = l.generarTreeNode(seleccion.Carpeta)
			}))

			//si no es del tipo pdf agrego un  laber
		} else {
			layoutTemp = append(layoutTemp,
				g.Style().SetColor(g.StyleColorText, color.RGBA{100, 100, 100, 255}).To(
					g.Label(" - "+archivo.Nombre)))
		}
	}

	return

}
