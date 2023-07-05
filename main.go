package main

import (
	"fmt"

	"github.com/Hasyim-Kai/Go_Gin_Rest_Api/practice"
)

type Person struct {
	name string
	age  int
}

func (p Person) yelyel() {
	fmt.Println("Nama Saya adalah " + p.name)
}

func main() {
	// sugiono := Person{"Sugiono", 69}
	// sugiono.yelyel()
	// practice.RunArray()
	practice.RunDeferPanicError()
}
