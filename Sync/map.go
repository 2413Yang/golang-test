package main

import (
	"fmt"
	"sync"
)

var m sync.Map

func main(){
	m.Store(1,"one")
	m.Store(2,"two")

	v,ok := m.LoadOrStore(3,"three")
	fmt.Println(v,ok)

	v,ok = m.LoadOrStore(1,"thisone")
	fmt.Println(v,ok)

	v,ok = m.Load(1)
	if ok{
		fmt.Println("key is existed, and value is :",v)
	}else{
		fmt.Println("key is not existed")
	}

	f := func(k,v interface{}) bool{
		fmt.Println(k,v)
		return true
	}
	m.Range(f)
	m.Delete(2)
	fmt.Println(m.Load(2))
}