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

func main() {
	/*asignandole valor a la estructura*/
	// uriel:=User{20,"yandali","damian"}
	// fmt.Println(uriel)

	/*dandole nuevo valor con punteros*/
	uriel := new(User)
	uriel.nombre ="mario"
	uriel.apellido ="abc"
	fmt.Println(uriel.nombre, uriel.apellido)

}