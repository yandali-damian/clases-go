package main

import (
	"fmt"
	"time"
	"strings"
)

//concurrentes== goroutines

// 

func main() {
	//sincrono
	// mi_nombre_lentoo("Yandali")
	// fmt.Println("que aburrido")
   // convertiremos el sincrono a concurrentes
	go mi_nombre_lentoo("Yandali")
	fmt.Println("que aburrido")

	//dividir el problema en diferentes ejecuciones es lo q se llama concurrente
	var wait string
	fmt.Scanln(&wait)
}

//funcion para dividir una palabra en letras
//split funcion para dividir una palabra o una cadena en varias cadenas
func mi_nombre_lentoo(nombre string){
 letras := strings.Split(nombre, "")

 //iteramos la cadena de la
 for _,letra :=range(letras) {

	//mi nombre se imprimira mas lento
	time.Sleep(1000 * time.Millisecond)

	fmt.Println(letra)
 }
}
