package main

import (
	"encoding/json"
	"fmt"
	"os"
)


func main() {
	data, err := dataLoader("products.json")
	if err != nil {
		panic(err)
    }
	fmt.Println(data)

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
for

	return fmt.Sprint(p), nil


}