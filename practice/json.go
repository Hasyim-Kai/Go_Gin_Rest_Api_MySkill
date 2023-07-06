package practice

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func RunJson() {
	defer Catch() // kode recover() harus defer
	var jsonString = `[
		{"Name": "john wick", "Age": 27},
		{"Name": "ethan hunt", "Age": 32}
	]`
	var data []User
	var err = json.Unmarshal([]byte(jsonString), &data) // decode to struct
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user 1 :", data[0].FullName)
	fmt.Println("user 2 :", data[1].FullName)
}
