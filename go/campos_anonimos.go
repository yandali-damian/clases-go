package main

import (
	"fmt"
)
/*
go no es oop
definir herencia
Nos ayuda a ingresar a los metodos y atributos los campos 
*/
type Human struct {
	name string
} 

func (this Human)hablar() string{
	return "bla bla bla"
}

type Dummy struct {
	name string
}

type Tutor struct {
	Human
	Dummy
}

func (this Tutor)hablar() string{
	return this.Human.hablar() + " Bienvenidxs *_*"
}

func main() {
	tutor:=Tutor{Human{"Yandali"}, Dummy{"Irwin"}}

	// fmt.Println(tutor.Dummy.name)
	fmt.Println(tutor.hablar())
}