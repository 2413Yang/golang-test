package main

import (
	"github.com/2413Yang/tcpServer/tree/master/golang/tcpChat/chatserver/server"
)

func main() {
	var s server.Server
	s = server.NewServer()
	s.Listen(":3333")
	s.Start()
}