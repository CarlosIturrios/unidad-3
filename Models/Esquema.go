//Autor: Carlos Armando Iturrios Alcaraz ITIC 10-1

package Models

import (
	"github.com/jinzhu/gorm"
)

//Estructura del modelo del libro
type Libro struct {
	gorm.Model
	Nombre            string `json:"Nombre"`
	Descripcion       string `json:"Descripcion"`
	Autor             string `json:"Autor"`
	Editorial         string `json:"Editorial"`
	Fecha_publicacion string `json:"Fecha_publicacion"`
}

func (b *Libro) TableName() string {
	return "libro"
}
