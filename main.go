//Autor: Carlos Armando Iturrios Alcaraz ITIC 10-1
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"unidad-3/BaseDeDatos"
	"unidad-3/Models"
	"unidad-3/Routers"
)

var err error

func main() {
	//Cadena conexion a la base de datos en docker
	BaseDeDatos.DB, err = gorm.Open("mysql", "user001:j23FJwHFvHJUyzkZ@(127.0.0.1:42333)/biblioteca?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("Estado: ", err)
	}
	defer BaseDeDatos.DB.Close()
	//Migracion del modelo libro a la base de datos
	BaseDeDatos.DB.AutoMigrate(&Models.Libro{})
	// mandar llamar el metodo para poder correr el router
	r := Routers.SetupRouter()
	// Correr el router y activar las uris para ser escuchadas por el server
	r.Run()
}
