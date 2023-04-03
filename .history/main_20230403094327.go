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
	/* data, err := dataLoader("products.json")
	if err != nil {
		panic(err)
    }
	fmt.Println(data) */
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
		/* id, error := strconv.Atoi(c.Param("id"))
		if error!= nil {
            c.JSON(http.StatusBadRequest, map[string]any{
				"error": "bad request",
            })
        }
		for i, p := range Products {
			if Products[i].Id == id{
                    c.IndentedJSON(http.StatusOK, p)
					
                }
				c.IndentedJSON(http.StatusNotFound, nil)
		} */
    }

	func ProductSearch(c *gin.Context) {
		price, err := strconv.ParseFloat(c.Query("priceGt"), 64)
		if err!= nil {
            c.JSON(http.StatusBadRequest, map[string]any{
                "error": "bad request",
            })
        }
		var products []Product
		for i := range Products {
			if Products[i].Price > price {
				products = append(products, Products[i])
                    c.IndentedJSON(http.StatusOK, products)
                    return
                }
                c.IndentedJSON(http.StatusNotFound, nil)
		}
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

/* func dataLoader(filename string) ([]Product, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
    }
	defer f.Close()


	decoder := json.NewDecoder(f)
	err = decoder.Decode(&Products)
	if err != nil {
		return nil, err
    }

	return Products, nil
} */