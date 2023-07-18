package main

import "fmt"
import "math"

func test() string {
	return "测试字符串"
}
func test_1() string {
	return "test"
}
type intalias = int32
func main() {
	// bar := test()
	fmt.Println(test_1())

	var x int32 = 23.0
	// var y int = x
	var z rune = x
	var c intalias = x
	fmt.Println(z,c)

	f1 := 12.1
	f2 := 12.1

	fmt.Println(math.Abs(f1-f2)<0.0001)

	c1 := 12 + 1i
	fmt.Println(real(c1),imag(c1))
}