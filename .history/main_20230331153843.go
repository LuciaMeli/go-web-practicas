package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)


func main() {
	data, err := dataLoader("products.json")
	if err != nil {
		panic(err)
    }
	fmt.Println(data)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, map[string]any{
			"message": "pong",
        })
    })
		
	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, data)
    })

	router.POST("/products/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, data)
    })

    router.POST("/products/search", func(c *gin.Context) {
		decoder := json.NewDecoder(c.Request.Body)

		type searchResult struct {
			Price float64 `json:"price"`
		}


		var query searchResult
		if err := decoder.Decode(&query); err != nil {
			c.JSON(http.StatusBadRequest, map[string]any){
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

func dataLoader(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
    }
	defer f.Close()

	

	var p []Product
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&p)
	if err != nil {
		return "", err
    }
for _, product := range p {
	return fmt.Printf("Id: %f, Name: %s, Quantity: %f, Code_value: %f, Is_published %t, Expiration %s, Price %f ", 
	product.Id, product.Name, product.Quantity, product.Code_Value, product.Is_published, 
	product.Expiration, product.Price)
	}, nil
}