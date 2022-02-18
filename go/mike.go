package main

import (
	"fmt"
)

func main() {

	slice := make([]int,5,7) //se genera un areglo de enteros de 3 columnas
	//tanto len y cap son para medir la longitud
	slice = append(slice,2)//append agregar un elemnto al arreglo
	fmt.Println(slice)
	// fmt.Println(cap(slince))
}