package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)


func main() {
	data, err := dataLoader("products.json")
	if err != nil {
		panic(err)
    }
	fmt.Println(data)

	router := gin.Default()

	router.GET("/ping", PingPong)
	router.GET("/products", GetProducts)
	router.GET("/products/:id", GetProductById)
	router.G
	
	
	func PingPong(c *gin.Context){
        c.string(http.StatusOK, "pong")
    }

	router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, map[string]any{
			"message": "pong",
        })
    })
		
	router.GET("/products", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, data)
    })

	router.POST("/products/:id", func(c *gin.Context) {
		id, error := strconv.Atoi(c.Param("id"))
		if error!= nil {
            c.JSON(http.StatusBadRequest, map[string]any{
				"error": "bad request",
            })
        }
		for _, a := range data{
			if a.Id == id{
                    c.JSON(http.StatusOK, a)
					return
                }
				c.IndentedJSON(http.StatusNotFound, nil)
		}
    })

    router.POST("/products/search", func(c *gin.Context) {
		decoder := json.NewDecoder(c.Request.Body)

		type searchQuery struct {
			Price float64 `json:"price"`
		}

		var query searchQuery
		if err := decoder.Decode(&query); err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{
				"error": "bad request error",			
			})
		}

		var matchProduct []Product
		for _, product := range data {
            if product.Name == query.Price {
                matchProduct = append(matchProduct, product)
            }
        }	
		c.JSON(http.StatusOK, map[string]any{
			"products": matchProduct,
            })
		})

	router.Run(":8080")	
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

func dataLoader(filename string) ([]Product, error) {
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

	/* var output []string
	for _, product := range p {
	output = append(output, fmt.Sprintf("Id: %f, Name: %s, Quantity: %f, Code_value: %f, Is_published %t, Expiration %s, Price %f ", 
	product.Id, product.Name, product.Quantity, product.Code_Value, product.Is_published, 
	product.Expiration, product.Price))
	} */
	return Products, nil
	
}