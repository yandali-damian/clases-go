package main
import (
	"fmt"
)

func main() {
	/*
	copy: lo que hace es copiar el minimo de elementos
	*/
	slice := []int{1,2,3,4}
	copia := make([]int,len(slice),cap(slice)*2)//el numero 3 es lo que define la cantidad de elementos a copiar

	copy(copia,slice)

	fmt.Println(slice)
	fmt.Println(copia)
}