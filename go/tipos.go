package main

import (
	"fmt"
	"strconv"
)

func main() {

	// entero

	// edad:=23
	/*Itoa: convierte un entero a string*/
	// edad_str :=strconv.Itoa(edad)
	// fmt.Println("Miedad es:"+edad_str)

	//string
	edad:="23"
	// aceptar multiples valores
	edad_int,_ :=strconv.Atoi(edad)/*Atoi: convierte un  string a entero*/
	fmt.Println(edad_int+10-5)

}