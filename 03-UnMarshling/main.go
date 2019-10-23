package main 
import (
	"fmt"
	"encoding/json"
)

type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Address Address `json:"address"`

}
type Address struct{
   City string `json:"city"`
   State string `json:"state"`
   Country string `json:"country"`
}

func main(){
jsonData:= []byte(`{
 "name": "Rajat",
 "age": 29,
 "address": {
  "city": "Jaipur",
  "state": "Rajasthan",
  "country": "India"
 }
}
`)
var person Person
err := json.Unmarshal(jsonData,&person)
if err != nil{
  fmt.Println("Found error in unmarshling",err)
}
fmt.Println(person.Name,person.Age,person.Address)
}