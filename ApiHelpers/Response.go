package ApiHelpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int
	Meta   interface{}
	Data   interface{}
}

func RespondJSON(w *gin.Context, status int, payload interface{}) {
	fmt.Println("Estado ", status)
	var res ResponseData

	res.Status = status
	res.Data = payload

	w.JSON(200, res)
}
