package practice

import (
	"fmt"
	"strconv"
)

func catch() {
	if r := recover(); r == nil {
		fmt.Println("app berjalan mulus")
	} else {
		fmt.Println("Error bang :", r)
	}
}

func RunDeferPanicError() {
	var input string
	var number int
	var err error
	defer catch() // kode recover() harus defer
	// defer fmt.Println("jalan pake defer")
	// fmt.Println("jalan")

	fmt.Print("Type a number : ")
	fmt.Scanln(&input)
	if number, err = strconv.Atoi(input); err == nil {
		fmt.Println(number, "is Number")
	} else {
		panic(err.Error())
		fmt.Println(number, "isnt Number")
	}
}
