package products

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

// used to hold our product list in memory
var productMap = struct {
	sync.RWMutex
	m map[int]product
}{m: make(map[int]product)}


//este metodo se llama cuando se inicializa el package
func init() {
	fmt.Println("loading products...")
	prodMap, err := loadProductMap()
	productMap.m = prodMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d products loaded...\n", len(productMap.m))
}

func loadProductMap() (map[int]product, error) {
	fileName := "certificates.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	productList := make([]product, 0)
	err = json.Unmarshal([]byte(file), &productList)
	if err != nil {
		log.Fatal(err)
	}
	prodMap := make(map[int]product)
	for i := 0; i < len(productList); i++ {
		prodMap[productList[i].Id] = productList[i]
	}
	return prodMap, nil
}

func getProduct(productID int) *product{
	productMap.RLock()
	defer productMap.RUnlock()
	if product, ok := productMap.m[productID]; ok {
		return &product
	}
	return nil
}

func removeProduct(productID int) {
	productMap.Lock()
	defer productMap.Unlock()
	delete(productMap.m, productID)
}

func getProductList() []product {
	productMap.RLock()
	products := make([]product, 0, len(productMap.m))
	for _, value := range productMap.m {
		products = append(products, value)
	}
	productMap.RUnlock()
	return products
}

func getProductIds() []int {
	productMap.RLock()
	productIds := []int{}
	for key := range productMap.m {
		productIds = append(productIds, key)
	}
	productMap.RUnlock()
	sort.Ints(productIds)
	return productIds
}

func getNextProductID() int {
	productIds := getProductIds()
	return productIds[len(productIds)-1] + 1
}

func addOrUpdateProduct(item product) (int, error) {
	// if the product id is set, update, otherwise add
	addOrUpdateID := -1
	if item.Id > 0 {
		oldProduct := getProduct(item.Id)
		// if it exists, replace it, otherwise return error
		if oldProduct == nil {
			return 0, fmt.Errorf("product id [%d] doesn't exist", item.Id)
		}
		addOrUpdateID = item.Id
	} else {
		addOrUpdateID = getNextProductID()
		item.Id = addOrUpdateID
	}
	productMap.Lock()
	productMap.m[addOrUpdateID] = item
	productMap.Unlock()
	return addOrUpdateID, nil
}
