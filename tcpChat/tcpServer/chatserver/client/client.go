package client

import (
	"github.com/2413Yang/tcpServer/tree/master/golang/tcpChat/chatserver/protocol"
)

type Client interface {
	/*Make a request to the server
		address : the server address*/
	Dial(address string) error 
	Start()
	Close()
	Send(command interface{}) error
	SetName(name string) error
	SendMess(message string) error
	InComing() chan protocol.MessCmd
}