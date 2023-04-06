package main

import(
	"fmt"
	"math/rand"
)
var done bool = false
var Mess = make(map[int]bool)
func main(){
	A := make(chan int)
	B := make(chan int)
	go sendRandom(50,10,B)
	go receive(A,B)
	sum(A)
}

func genRandom(max,min int)int{
	return rand.Intn(max-min)+min
}

func sendRandom(max,min int, out chan<- int){
	for{
		if done {
			close(out)
			return
		}
		out <- genRandom(max,min)
	}
}
func receive(out chan<- int,in <-chan int){
	for r := range in{
		fmt.Println(" ",r)
		_,ok := Mess[r]
		if ok{
			fmt.Println("duplicate num is:",r)
			done = true
		}else{
			Mess[r] = true
			out <- r
		}
	}
	close(out)
}
func sum(in <- chan int){
	var sum int
	for i := range in{
		sum += i
	}
	fmt.Println("The sum is :",sum)
}