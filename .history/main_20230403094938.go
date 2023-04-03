package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func main() {
	//read file and make a products slice from it
	file, err := ioutil.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &Products)
	if err != nil {
		log.Fatal(err)
	}


	router := gin.Default()

	router.GET("/ping", PingPong)
	router.GET("/products", GetAllProducts)
	router.GET("/products/:id", ProductById)
	router.GET("/products/search", ProductSearch)

	router.Run(":8080")
}


	func PingPong(c *gin.Context){
		c.String(http.StatusOK, "pong")
	}
		
	func GetAllProducts(c *gin.Context) {
        c.IndentedJSON(http.StatusOK, Products)
    }

	func ProductById(c *gin.Context) {
		var Id int

	if value, err := strconv.Atoi(c.Param("id")); err == nil {
		Id = value
	}

	product, ok := ProductExist(Id)
	if ok {
		c.IndentedJSON(http.StatusOK, product)
	} else {
		c.IndentedJSON(http.StatusNotFound, nil)
	}
}

	func ProductExist(index int) (Product, bool) {

		var product Product
	
		for i := range Products {
			if Products[i].Id == index {
				return Products[i], true
			}
		}
		return product , false
	}

	func ProductSearch(c *gin.Context) {
		price, err := strconv.ParseFloat(c.Query("priceGt"), 64)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, FindProducts(price))

}
func FindProducts(price float64) []Product{

	var products []Product

	for i := range Products {
		if Products[i].Price > price {
			products = append(products, Products[i])
		}
	}

	return products
}

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	Code_Value string `json:"code_value"`
	Is_published bool `json:"is_published"`
	Expiration string `json:"expiration"`
	Price float64 `json:"price"`
}

var Products = []Product{}