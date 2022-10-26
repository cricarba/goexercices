package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/cricarba/goexercices/products"
)

const basePath = "/api"

func main() {
	products.SetupRoutes(basePath)
	fs := http.FileServer(http.Dir("./static/test.html")) // ccreamos el servidor de archivos

	// cuando lleguen petidcion /static/ rediriga a el server de archivos
	http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", home)// primer parametro url, SEGUNDO la funcion que se procesa
	http.HandleFunc("/info/", getInfoRequest)
	err := http.ListenAndServe(":8082", nil) // segundo parametro es una funcion que se llama cada vez que llegue una funcion puede ser un middleware
	if err != nil {
		log.Fatal(err)
	}
}
// Escribo en al respuesta de la peticion
func home(w http.ResponseWriter, r *http.Request){
   w.Write([]byte("<html>Home</html>"))
}

func getInfoRequest(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Host: ",r.Host)
	fmt.Fprintln(w, "Uri: ",r.RequestURI)
	fmt.Fprintln(w, "Method: ",r.Method)
	fmt.Fprintln(w, "RemoteAddr: ",r.RemoteAddr)
}