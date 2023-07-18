package main

import (
	"fmt"
	"unsafe"
)
type Test struct{
	Name string
}
func new() Test {
	return Test{Name: "name"}
}
func main() {
	fmt.Println(unsafe.Pointer(&Test{Name:	"test"}))

	s := [...]int{1,2,3}
	fmt.Println(&s[0])
	fmt.Println(&s[1])
	fmt.Println(&s[2])

	fmt.Println([]int{1,2,3}[1:])

	b := new()
	fmt.Println(unsafe.Pointer(&b))
}