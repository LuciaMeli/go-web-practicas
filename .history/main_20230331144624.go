package main

import "os"


func main() {

}


type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	Code_Value string `json:"code_value"`
	Is_published bool `json:"is_published"`
	Expiration 
}

func dataLoader(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
    }
	defer f.Close()

	reader := 

	var p Product


}