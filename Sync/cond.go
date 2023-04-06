package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ready = false
	singerNum = 3
)

func Sing(singerID int,c *sync.Cond) {
	fmt.Printf("Singer (%d) is ready\n",singerID)
	c.L.Lock()
	for !ready {
		fmt.Printf("Singer (%d) is waiting\n",singerID)
		c.Wait()
	}
	fmt.Printf("Singer (%d) sing a song\n",singerID)
	ready = false //broadcast使用保证每次只有一个goroutine处于运行状态
	c.L.Unlock()
}

func main(){
	var wg sync.WaitGroup
	cond := sync.NewCond(&sync.Mutex{})
	for i:=0;i<singerNum;i++{
		go Sing(i,cond)
		wg.Add(1)
	}
	time.Sleep(3*time.Second)
	go func(){
		for i:=0;i<singerNum;i++{
			ready = true
			// cond.Signal() //单个通知
			cond.Broadcast()	//广播
			time.Sleep(3*time.Second)
			wg.Done()
		}
	}()
	wg.Wait()
}