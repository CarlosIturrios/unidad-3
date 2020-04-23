package Routers

import (
	"github.com/gin-gonic/gin"
	"unidad-3/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("libro", Controllers.ListarLibro)
		v1.POST("libro", Controllers.InsertarNuevoLibro)
		v1.GET("libro/:id", Controllers.ObtenerLibroPorID)
		v1.PUT("libro/:id", Controllers.ActualizarLibro)
		v1.DELETE("libro/:id", Controllers.BorrarLibro)
	}

	return r
}
