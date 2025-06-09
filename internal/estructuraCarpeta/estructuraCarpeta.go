package estructuraCarpeta

// generar la estructura de archivos de la carpeta seleccionada

import (
	"fmt"
	"os"
	"strings"
)

type Carpeta struct {
	Nombre      string
	Path        string
	SubCarpetas []Carpeta
	Archivos    []Archivo
}

type Archivo struct {
	Nombre string
	Tipo   string
	Path   string
}

func NuevaCarpeta(pathCarpeta string) Carpeta {

	// cargamos la Carpeta
	files, err := os.ReadDir(pathCarpeta)

	if err != nil {
		fmt.Println(err)

		return Carpeta{}
	}

	nuevaCarpeta := Carpeta{
		Nombre: strings.ToLower(pathCarpeta[strings.LastIndex(pathCarpeta, "\\")+1:]),
		Path:   pathCarpeta,
	}

	for _, file := range files {

		// si encuentro una Carpeta, genero el sub-listado (iteración), y lo agrego como un treeNode
		if file.IsDir() {

			nuevaCarpeta.SubCarpetas = append(nuevaCarpeta.SubCarpetas, NuevaCarpeta(pathCarpeta+"\\"+file.Name()))

			// si encuentro un Archivo
		} else {

			nuevoArchivo := NewArchivo(file, pathCarpeta)

			// no me interesa los Archivos internos del sistema
			if nuevoArchivo.Tipo == "db" {
				continue
			}

			nuevaCarpeta.Archivos = append(nuevaCarpeta.Archivos, nuevoArchivo)
		}
	}

	return nuevaCarpeta
}

func NewArchivo(file os.DirEntry, pathCarpeta string) Archivo {

	return Archivo{
		Nombre: file.Name(),

		Tipo: strings.ToLower(file.Name()[strings.LastIndex(file.Name(), ".")+1:]),

		Path: pathCarpeta + "\\" + file.Name(),
	}
}
