//Autor: Carlos Armando Iturrios Alcaraz ITIC 10-1

package Controllers

import (
	"github.com/gin-gonic/gin"
	"unidad-3/ApiHelpers"
	"unidad-3/Models"
)

//Clase para realizar un nuevo libro en la base de datos
func InsertarNuevoLibro(c *gin.Context) {
	var libro Models.Libro
	c.BindJSON(&libro)
	err := Models.InsertarNuevoLibro(&libro)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, err)
	} else {
		ApiHelpers.RespondJSON(c, 200, libro)
	}
}

//Clase para actualizar un libro guardado en la base de datos
func ActualizarLibro(c *gin.Context) {
	var libro Models.Libro
	id := c.Params.ByName("id")
	err := Models.ObtenerLibroPorID(&libro, id)

	//Validacion para regresar un mensaje de no encontrado en caso de que el ID no exista en la base de datos
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, "Libro no encontrado")
	} else {
		c.BindJSON(&libro)
		err = Models.ActualizarLibro(&libro, id)
		if err != nil {
			ApiHelpers.RespondJSON(c, 404, err)
		} else {
			ApiHelpers.RespondJSON(c, 200, libro)
		}
	}

}

//Clase para obtener todos los libros de la base de datos
func ListarLibro(c *gin.Context) {
	var libro []Models.Libro
	err := Models.ListarLibro(&libro)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, "No hay libros registrados")
	} else {
		ApiHelpers.RespondJSON(c, 200, libro)
	}
}

//Clase para obtener un libro por su ID
func ObtenerLibroPorID(c *gin.Context) {
	id := c.Params.ByName("id")
	var libro Models.Libro
	err := Models.ObtenerLibroPorID(&libro, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, "Libro no encontrado")
	} else {
		ApiHelpers.RespondJSON(c, 200, libro)
	}
}

//Clase para borrar el registro de un libro
func BorrarLibro(c *gin.Context) {
	var libro Models.Libro
	id := c.Params.ByName("id")
	err := Models.BorrarLibro(&libro, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, "Libro no encontrado")
	} else {
		ApiHelpers.RespondJSON(c, 200, libro)
	}
}
