package main

import "os"


func main() {

}


type Product struc

func dataLoader(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
    }
	defer f.Close()

	reader := 

	var p Product


}