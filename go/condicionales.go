package main

import (
	"fmt"
)
// sobre:
// if 
// else if
// else

// func main() {
// 	if true{
// 		fmt.Println("Hola Mundo")
// 	}
// }

/*
	operadores logicos 
	== igual a
	!= diferente de 
	<
	>
	>=
	<=
	&& and
	|| or solo requiere q uno sea verdadero

*/

func main() {
	// x:=10
	// y:=10
	var numero1 int
	var numero2 int
	fmt.Println("ingrese primer numero:")
	fmt.Scanf("%d\n", &numero1)

	fmt.Println("ingrese segundo numero:")
	fmt.Scanf("%d\n", &numero2)

	if numero1>numero2{
		fmt.Printf("%d es mayor que %d \n",numero1,numero2)
	}else if numero1<numero2{
		fmt.Printf("%d es Menor que %d \n",numero1,numero2)
	}else{
		fmt.Println("son iguales")
	}
}