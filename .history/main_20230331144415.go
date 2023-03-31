package main

import "os"


func main() {

}


func dataLoader(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
    }
	defer f.Close()

	reader := 

	var p Pro


}