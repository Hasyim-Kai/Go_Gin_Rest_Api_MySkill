package practice

import (
	"fmt"
	"math"
)

type persegiBlueprint interface {
	luas() float64
	keliling() float64
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func RunInterface() {
	defer Catch() // kode recover() harus defer
	var object1 persegiBlueprint = persegi{10.0}
	fmt.Println("===================================")
	fmt.Println("luas :", object1.luas())
	fmt.Println("keliling :", object1.keliling())
	fmt.Println("===================================")
}
