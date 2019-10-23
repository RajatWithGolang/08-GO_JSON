package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int   `json:"age"`
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
	var p1 Person
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&p1)
	fmt.Printf("Name is = %s",p1.Name)
	
}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))

}
