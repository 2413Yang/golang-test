package main

import(
	"fmt"
	"sync"
)

const(
	noGoroutine = 2
	onTask = 10
)
var wg sync.WaitGroup

func main(){
	tasks := make(chan int,onTask)
	for i := 1; i <= noGoroutine; i++{
		wg.Add(1)
		go taskProcess(tasks,i)
	}
	for taskNO := 1;taskNO <= onTask;taskNO++{
		tasks<-taskNO
	}
	fmt.Println("close before")
	close(tasks)
	fmt.Println("close after")
	wg.Wait()
	fmt.Println("endl")
}
func taskProcess(tasks chan int,workerNo int){
	defer wg.Done()
	
	for t := range tasks{
		fmt.Printf("Worker %d is processing Task no:%d, chanel len : %d \n",workerNo,t,len(tasks))
	}
	fmt.Printf("worker %d got off work \n",workerNo)
}