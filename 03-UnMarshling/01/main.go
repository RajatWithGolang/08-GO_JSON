package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"strconv"
)

type Person struct {
	Name string `json:"name"`
	}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var p1 Person
	// read data from request Body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	// Decode data to object
	err = json.Unmarshal(data, &p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Name is = %s", p1.Name)
}
func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
