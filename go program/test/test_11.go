package main

import "fmt"
import "strings"

func main() {
	var str1 string = "ababcdeba"
	fmt.Println(str1)
	fmt.Println(strings.TrimSpace(str1))
	fmt.Println(strings.Trim(str1, "ab"))
	fmt.Println(strings.TrimLeft(str1, "ab"))
	fmt.Println(strings.TrimRight(str1, "ab"))
	fmt.Println(strings.Fields(str1))
	fmt.Println(strings.Replace(str1, "a", "c", strings.Count(str1, "a")))
	fmt.Println(strings.Count(str1, "a"))
	fmt.Println(strings.Repeat(str1, 2))
	str2 := strings.Split(str1, "a")
	fmt.Println(str2)
	fmt.Println(str2[1])
	fmt.Println(strings.Join(str2, "a"))
}
