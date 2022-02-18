package main

//importando el paquete para el servidor web
import (
	"net/http"
)

//funcion principal
func main() {
// println("hello world")

//creando rutas
http.HandleFunc("/",homeHandler)
http.HandleFunc("/contact",contactHandler)

// inicia el servidor
// ListenAndServe: este metodo escucha y sirve un contenido
http.ListenAndServe(":3000", nil)
}

// manejador de la ruta inicial
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola Mundo Cruel"))

}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pagina de Contacto"))

}