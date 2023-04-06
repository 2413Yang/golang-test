package main
import(
	"fmt"
	"time"
)
func main(){
	chan1 := make(chan int)
	chan2 := make(chan int)
	go func(){
		time.Sleep(1*time.Second)
		chan1<-1
	}()
	go func(){
		time.Sleep(1*time.Second)
		chan2<-2
	}()
	select {
	case <-chan1:
		fmt.Println("ch1")
	case <-chan2:
		fmt.Println("ch2")
	// default:
	}
	if len(chan1) > 0{
		fmt.Println("ch1:",<-chan1)
	}
	if len(chan2) > 0{
		fmt.Println("ch2:",<-chan2)
	}
	fmt.Println("endl")

}