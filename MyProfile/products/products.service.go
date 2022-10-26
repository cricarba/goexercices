package products

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"github.com/cricarba/goexercices/cors"
)

// SetupRoutes : define path en route
const productsPath = "product"


//Define Routes set de basepath /api/
func SetupRoutes(apiBasePath string) {
	fmt.Println("Disparo estas funciones para llenar la lista de productos")
	//Disparo estas funciones para llenar la lista de productos
	productListHandler := http.HandlerFunc(handleProducts)
	productItemHandler := http.HandlerFunc(handleProduct)


	// creo el ruteo para las peticiones a /api/products/ y ejecuto el middleware
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, productsPath), cors.Middleware(productListHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, productsPath), cors.Middleware(productItemHandler))
}


//Routing for DELETE, GET/{0}, UPDATE
func handleProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT DELETE")
	
	urlPathSegments := strings.Split(r.URL.Path, "product/")

	//get productId from url segments y convierte a entero Atoi
	productId, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
    //prodId, ok := r.URL.Query()["id"] tomar los parametros del query string
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	productItem := getProduct(productId)

	if productItem == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		//convertir a json un struc
		productJson, err := json.Marshal(productItem)

		//fmt.Println(err)
		fmt.Println(productJson)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Header().Set("Content-Tye", "application/json")
		w.Write(productJson)
	case http.MethodPut:
		var updateProduct product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return

		}
		err = json.Unmarshal(bodyBytes, &updateProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return

		}
		if updateProduct.Id != productId {
			w.WriteHeader(http.StatusBadRequest)
			return

		}

		addOrUpdateProduct(updateProduct)

		w.WriteHeader(http.StatusOK)
		return
	case http.MethodDelete:
		removeProduct(productId)
		return
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// ROUTING POST A GET 
func handleProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET O POST")
	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET")
		productList := getProductList()
		productJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Tye", "application/json")
		w.Write(productJson)

	case http.MethodPost:
		fmt.Println("POST")
		var newproduct product
		
		bodyBytes, err := ioutil.ReadAll(r.Body)
		
		if err != nil {
			
			w.WriteHeader(http.StatusBadRequest)
			
			return

		}
		
		err = json.Unmarshal(bodyBytes, &newproduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("json")
			return

		}

		fmt.Println(newproduct)
		if newproduct.Id != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return

		}
		
		_, err = addOrUpdateProduct(newproduct)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)

	case http.MethodOptions:
		return
	}
}
