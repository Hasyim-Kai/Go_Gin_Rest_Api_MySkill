package practice

import "fmt"

func RunArray() {
	array := [3]string{"Melon", "Papain", "Peach"}
	slice := []string{"Lion", "Tiger", "Elephant"}
	fmt.Println(array)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, "Dragon", "Leviathan")
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	mammals := make([]string, 3)
	copy(mammals, slice[0:3])
	fmt.Println(mammals)
	fmt.Println(len(mammals))
	fmt.Println(cap(mammals))
}
