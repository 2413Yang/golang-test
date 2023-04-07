package main 

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func main() {
	addr := "www.tianmao.com:80"
	tcpAddr,err := net.ResolveTCPAddr("tcp",addr)
	if err != nil {
		fmt.Println("Error:",err.Error())
		return
	}
	t1 := time.Now().Unix()
	myConn,err := net.DialTCP("tcp",nil,tcpAddr)
	if err != nil {
		fmt.Println("Error:",err.Error())
		return
	}
	t2 := time.Now().Unix()
	_,err = myConn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
	if err != nil {
		fmt.Println("Error:",err.Error())
		return
	}
	t3 := time.Now().Unix()
	result,err := ioutil.ReadAll(myConn)
	t5 := time.Now().Unix()
	if err != nil {
		fmt.Println("Error:",err.Error())
		return
	}
	t4 := time.Now().Unix()
	fmt.Println(string(result))
	fmt.Println("Dial:",t1,"Write:",t2,"Read:",t3,"end",t4," t5: ",t5)
	os.Exit(0)
}