package main

import (
	"fmt"
)

/*
	structs 
*/

type User struct {
	edad int
	nombre,apellido string
	
}

/*pasando una funcion a la estructura*/
//this es un identificador de la estructura

func (this User) nombre_completo() string {

	return this.nombre + " " + this.apellido

}

func main() {

	uriel := new(User)
	uriel.nombre ="mario"
	uriel.apellido ="abc"
	fmt.Println(uriel.nombre_completo())

}