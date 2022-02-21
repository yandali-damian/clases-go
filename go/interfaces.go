package main

import(
	"fmt"
)

/*
estructura
es un tipo de datos que se puede pasar entre funciones
*/

//tipo de interfas
//una interface esta compuesta por estructura con  metodos que no estan implementados 

type User interface{
	Permisos() int //nombre y tipo de dato
	Nombre() string
}

type Admin struct {
	nombre string
}

func (this Admin) Permisos() int { 
	return 5
}

func (this Admin) Nombre() string {  
	return this.nombre
}

/*nueva estructura*/

type Editor struct {
	nombre string
}

func (this Editor) Permisos() int { 
	return 3
}

func (this Editor) Nombre() string {  
	return this.nombre
}

//funcion para indicar si es admin que deje pasar y sino no deja pasar
func auth(user User) string{
	persmisos:=user.Permisos()
	if persmisos > 4 {
		return user.Nombre() + "Tiene Permisos de Administrador"
	}else if persmisos == 3 {
		return user.Nombre() + "Tiene Permisos de editor"
	}
	return ""
}

func main() {
// admin := Admin{"Yandali "}
// editor := Editor{"Michi "}

//arreglos las lineas 58 y59 es la misma q la linea 62
usuarios := []User{Admin{"Yandali "}, Editor{"Michi "}}

//iterar usuarios
for _,usuario := range usuarios {
	fmt.Println(auth(usuario))
}

// fmt.Println(auth(admin))
// fmt.Println(auth(editor))
}