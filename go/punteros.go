package main

import (
	"fmt"
)

func main() {
	/*
	1.- Es una direccion de menoria
	2.- con los punteros en lugar de obtener el valor ,
	 tenemos la direccion en la que esta el valor
	3.- x,y =>as123d (la memoria) = 5
	4.- x =>as123d =(nuevo valor) 6, por lo tanto y toma el nuevo valor de
	5.- y=6
	los tipos de datos se declaran:
	*t=>*int,*string.*struct
	valor cero =nil(nulo)
	 */
	var x,y *int
	entero:=5
	//& retorna el valor de la memoria
	x=&entero
	y=&entero

	// fmt.Println(x)
	// fmt.Println(y)

	// * con el asterisco asignamos el valor de la memoria
	*x=6
	fmt.Println(*x)
	fmt.Println(*y)
}