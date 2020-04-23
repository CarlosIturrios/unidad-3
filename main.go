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

	BaseDeDatos.DB, err = gorm.Open("mysql", "user001:j23FJwHFvHJUyzkZ@(127.0.0.1:42333)/biblioteca?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("Estado: ", err)
	}
	defer BaseDeDatos.DB.Close()
	BaseDeDatos.DB.AutoMigrate(&Models.Libro{})

	r := Routers.SetupRouter()
	// running
	r.Run()
}
