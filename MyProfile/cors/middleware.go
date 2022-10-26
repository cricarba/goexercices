package cors

import (
	"net/http"
	"fmt"
)

// Middleware : intercept request and add headers
func Middleware(handler http.Handler) http.Handler {
	fmt.Println("Middelware de CORS")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*") //si existe una con el mismo nombre agrega dos
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") //si existe borra la anterior y escribe una nueva
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		handler.ServeHTTP(w, r)
	})
}
