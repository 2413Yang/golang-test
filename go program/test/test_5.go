package main

import "fmt"

type Bird struct {

}

func (b *Bird) Fly() {
	fmt.Println("Bird fly")
}

type IFly interface {
	Fly()
}

func main() {
	var fly IFly = new(Bird)
	fly.Fly()
}