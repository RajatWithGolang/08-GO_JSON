package main

import (
	"encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
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
     var usersData Users
    json.Unmarshal(jsonData,&usersData)
    fmt.Println(len(usersData.Users))
    for i:=0; i< len(usersData.Users); i++{
   fmt.Println("Name of the user is ", usersData.Users[i].Name)
    fmt.Println("Name of the user is ", usersData.Users[i].Type)
     fmt.Println("Name of the user is ", strconv.Itoa(usersData.Users[i].Age))
      fmt.Println("Name of the user is ", usersData.Users[i].Social)
    }

}
