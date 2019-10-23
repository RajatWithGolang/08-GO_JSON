package main

import (
	"encoding/json"
	//"time"
	//"log"
	"fmt"
)
func main(){
	bolB,_ := json.Marshal(true)
	fmt.Println(string(bolB))
	fmt.Println("===================================")
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fmt.Println("===================================")
    fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	fmt.Println("===================================")
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	fmt.Println("===================================")
	slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	fmt.Println("===================================")
	mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	
}