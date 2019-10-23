package main

import (
	"encoding/json"
	"fmt"
	
    "io/ioutil"
    "os"
    //"strconv"
)
type Users struct {
    Users []User `json:"users"`
}
type User struct {
    Name   string `json:"name"`
    Type   string `json:"type"`
    Age    int    `json:"Age"`
    Social Social `json:"social"`
}
type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}
func main(){
     jsonFile,err := os.Open("user.json")
     if err != nil{
         fmt.Println(err)
     }
	 defer jsonFile.Close()
	 jsonData,err := ioutil.ReadAll(jsonFile)
	 if err != nil{
         fmt.Println(err)
     }
	 var myData map[string]interface{}
	 err = json.Unmarshal(jsonData,&myData)
     if err != nil{
         fmt.Println(err)
     }	
		fmt.Println(myData)
	
	}  

