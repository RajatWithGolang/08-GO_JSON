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
	address := Address{"Jaipur","Rajasthan","India"}
	person := Person{Name : "Rajat",Age : 29, Address : address}
	sb,_ := json.MarshalIndent(person,""," ") // for formatted output
	fmt.Println(string(sb))
}