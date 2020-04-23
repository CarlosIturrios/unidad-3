package Controllers

import (
	"github.com/gin-gonic/gin"
	"unidad-3/ApiHelpers"
	"unidad-3/Models"
)

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

func ActualizarLibro(c *gin.Context) {
	var libro Models.Libro
	id := c.Params.ByName("id")
	err := Models.ObtenerLibroPorID(&libro, id)

	if err != nil {
		ApiHelpers.RespondJSON(c, 404, "Libro no encontrado")
	} else{
		c.BindJSON(&libro)
		err = Models.ActualizarLibro(&libro, id)
		if err != nil {
			ApiHelpers.RespondJSON(c, 404, err)
		} else {
			ApiHelpers.RespondJSON(c, 200, libro)
		}
	}

}

func ListarLibro(c *gin.Context) {
	var libro []Models.Libro
	err := Models.ListarLibro(&libro)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, "No hay libros registrados")
	} else {
		ApiHelpers.RespondJSON(c, 200, libro)
	}
}

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
