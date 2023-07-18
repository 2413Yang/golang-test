package main

import "fmt"
import "os"

func main() {
	var str string
	str = "Hello world"
	ch := str[1]
	fmt.Printf("%s %d\n",str,len(str))
	fmt.Printf("%c %d\n",ch,len(str))
	fmt.Println(os.Args[0:])

	s1 := []int{1,2,3,4}
	fixSlice(s1)
	fmt.Println(s1,s1[3])

	s2 := make([]int,0,10)
	s2 = append(s2,s1...)
	fmt.Println(s2[3])
	i := 1
	switch i {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fallthrough
	case 5:
		fmt.Println(5)
	case 6:
		fmt.Println(6)
	} 
}

func fixSlice(s []int){
	s[0] = 0
}