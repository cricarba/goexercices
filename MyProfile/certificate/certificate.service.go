package certificate

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

	//get certificateId from url segments
	certificateId, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	certificate := getProduct(certificateId)

	if certificate == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		certificateJson, err := json.Marshal(certificate)

		fmt.Println(err)
		fmt.Println(certificateJson)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Header().Set("Content-Tye", "application/json")
		w.Write(certificateJson)
	case http.MethodPut:
		var updateCertificate certificates
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return

		}
		err = json.Unmarshal(bodyBytes, &updateCertificate)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return

		}
		if updateCertificate.Id != certificateId {
			w.WriteHeader(http.StatusBadRequest)
			return

		}

		addOrUpdateProduct(updateCertificate)

		w.WriteHeader(http.StatusOK)
		return
	case http.MethodDelete:
		removeProduct(certificateId)
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
		certificateList := getProductList()
		certiJson, err := json.Marshal(certificateList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Tye", "application/json")
		w.Write(certiJson)

	case http.MethodPost:
		var newCertificate certificates
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return

		}
		err = json.Unmarshal(bodyBytes, &newCertificate)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return

		}
		if newCertificate.Id != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return

		}
		_, err = addOrUpdateProduct(newCertificate)
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
