package main

import (
	"fmt"
	"sync"
	"time"
)
const varNumber  = 100000000
var byteSlicePool = sync.Pool{
	New : func() interface{} {
		fmt.Println("create");
		b := make([]byte,1024)
		return &b
	},
}

func main() {
	t1 := time.Now().Unix()

	for i:=0;i<varNumber;i++{
		bytes := make([]byte,1024)
		_ = bytes
	}
	t2 := time.Now().Unix()
	for i:=0;i<varNumber;i++{
		bytes := byteSlicePool.Get().(*[]byte)
		_ = bytes
		byteSlicePool.Put(bytes)
	}
	t3 := time.Now().Unix()
	fmt.Printf("不使用pool %d s \n:",t2-t1)
	fmt.Printf("使用pool: %d s \n",t3-t2)
}