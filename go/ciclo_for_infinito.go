package main

import (
	"fmt"
)
// solo cuenta con un ciclo FOR
func main() {
	// ciclo infnitos
	i:=0

	for{
		/*No va a imprimir el valor de 2*/
		if i==2{
			i++
			continue
		}

		fmt.Println(i)
		i++
		if i>10 {
			break
		}
	}
}