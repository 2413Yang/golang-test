package main

import "fmt"

func main() {
	test()
	panic("main panic了!")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err 信息是: %+v\n",err)
		}
	}()
}