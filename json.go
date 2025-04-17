package main

import (
	"fmt"
	"encoding/json"
)

/*	json.Unmarsal = decode data json dalam bentuk []byte, argument ke 2 diisi dengan variabel pointer yang akan menampung hasil dari decode
 */

type user struct {
	FullName string `json:"Name"`
	Age int
}

func main() {
	var jsonString = `{"Name" : "John Wick", "Age" : 27}`

	var data user

	err := json.Unmarshal([]byte(jsonString), &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Name:", data.FullName)
	fmt.Println("Age:", data.Age)
}