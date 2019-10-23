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
	p1 := Person{"Rajat", 29}
	jsonData, err := json.Marshal(p1) 
	if err != nil {
		panic("ooops")
	}
	fmt.Fprintf(w, string(jsonData))
}

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatalln(http.ListenAndServe(":8080", nil))

}
