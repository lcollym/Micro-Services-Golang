package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hola, Mundo!")
}

func main() {
	http.HandleFunc("/", handler) // Asocia la ruta "/" con la función handler
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil) // Inicia el servidor en el puerto 8080
}
