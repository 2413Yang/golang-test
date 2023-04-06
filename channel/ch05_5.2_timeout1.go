package main
import(
	"fmt"
	"time"
)
func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func(){
		time.Sleep(2*time.Second)
		ch1 <- "ch1 is ready!"
	}()
	go func(){
		time.Sleep(1*time.Second)
		ch2 <- "ch2 is ready!"
	}()
	select{
	case <- ch1:
		fmt.Println(<-ch1)
	case <- ch2:
		fmt.Println(<-ch2)
	case t:= <- time.After(3*time.Second):
		fmt.Println("ch1 timeout!",t)
	}
}