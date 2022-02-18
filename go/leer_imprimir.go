package main

import(
	"fmt"
)

func main() {
	// edad:=22
	// costo:=15.5
	// verbos q nos permiten interpretar
	// printf %d %v %e %t %f %b %s
	// fmt.Printf("Mi edad es %f \n",costo)
var nombre string
fmt.Println("coloca tu nombre: ")
fmt.Scanf("%v\n",&nombre)
fmt.Printf("mi nombre es %v \n", nombre)
}