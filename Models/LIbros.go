//Autor: Carlos Armando Iturrios Alcaraz ITIC 10-1
package Models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"unidad-3/BaseDeDatos"
)

//Clases para trabajar con la base de datos mediante el ORM

func InsertarNuevoLibro(libro *Libro) (err error) {
	if err = BaseDeDatos.DB.Create(libro).Error; err != nil {
		return err
	}
	return nil
}

func ActualizarLibro(libro *Libro, id string) (err error) {
	fmt.Println(libro)
	BaseDeDatos.DB.Save(libro)
	return nil
}
func ListarLibro(libro *[]Libro) (err error) {
	if err = BaseDeDatos.DB.Find(libro).Error; err != nil {
		return err
	}
	return nil
}

func ObtenerLibroPorID(libro *Libro, id string) (err error) {
	if err := BaseDeDatos.DB.Where("id = ?", id).First(libro).Error; err != nil {
		return err
	}
	return nil
}

func BorrarLibro(libro *Libro, id string) (err error) {
	BaseDeDatos.DB.Where("id = ?", id).Delete(libro)
	return nil
}
