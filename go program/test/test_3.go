package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil{
			fmt.Printf("err信息是:%+v\n",err)
		}
	}()
	panic("我是panic!")
}