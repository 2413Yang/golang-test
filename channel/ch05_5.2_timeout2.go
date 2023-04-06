package main
import(
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	rand.Seed(time.Now().UnixNano())
	no := rand.Intn(6)
	no *= 1000
	du := time.Duration(int32(no))*time.Millisecond
	fmt.Println("timeout duration is:",du)
	wg.Done()
	if isTimeout(&wg,du){
		fmt.Println("Time out!")
	}else{
		fmt.Println("Not time out")
	}
}
func isTimeout(wg *sync.WaitGroup,du time.Duration)bool{
	ch1 := make(chan int)
	go func(){
		time.Sleep(1*time.Second)
		ch1<-1
		defer close(ch1)
		wg.Wait()
	}()
	select{
	case num := <-ch1:
		fmt.Println(num)
		return false
	case <-time.After(du):
		return true
	}
}