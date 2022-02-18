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

// //agtregando metodos a las estructuras
// recibiendo la estructura como un puntero (*User)
// *User es un argumento que se pasa como una copia
func (this *User) set_nombre(n string) {

	this.nombre = n

}

func main() {

	uriel := new(User)
	uriel.nombre ="mario"
	uriel.set_nombre("abcdefg")
	fmt.Println(uriel.nombre)

}